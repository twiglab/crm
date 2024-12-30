package bc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/twiglab/crm/wechat/mq"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
	"github.com/twiglab/crm/wechat/web"
)

type BcExchange struct {
	BC       *mq.XMQ
	ApiV3Key string
}

func parse(s string) time.Time {
	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Now()
	}
	return x
}

func BusiCircleAuth(exchange BcExchange) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := wechat.V3ParseNotify(r)
		if err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
		ca := BusinessCircleAuthorSource{}
		if err := req.DecryptCipherTextToStruct(exchange.ApiV3Key, &ca); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}

		msg := &msg.BusinessCircleAuthor{
			MchID:    ca.MchID,
			OpenID:   ca.OpenID,
			Code:     ca.Code,
			AuthType: ca.AuthType,

			BusinessCircleMsg: msg.BusinessCircleMsg{
				MsgID:      req.Id,
				Summary:    req.Summary,
				MsgType:    req.EventType,
				CreateTime: parse(req.CreateTime),
			},
		}

		body, _ := mq.FromMsgp(msg)

		ctx := r.Context()

		if err := exchange.BC.Send(ctx, bc.MQ_WX_TOC_BC_AUTH, body); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
	}
}

func BusiCirclePayment(exchange BcExchange) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := wechat.V3ParseNotify(r)
		if err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
		ca := BusinessCirclePaymentSource{}
		if err := req.DecryptCipherTextToStruct(exchange.ApiV3Key, &ca); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}

		ctx := r.Context()

		timeEnd, _ := time.Parse(time.RFC3339, ca.TimeEnd)
		// 这里代码有错，请黄总修改正确
		msg := &msg.BusinessCirclePayment{
			AppID:                  ca.AppID,
			OpenID:                 ca.OpenID,
			Amount:                 ca.Amount,
			TimeEnd:                timeEnd,
			TransactionID:          ca.TransactionID,
			BusinessCircleShopBase: msg.BusinessCircleShopBase(ca.BusinessCircleShopBase),
			BusinessCircleMsg: msg.BusinessCircleMsg{
				MsgID:      req.Id,
				Summary:    req.Summary,
				MsgType:    req.EventType,
				CreateTime: parse(req.CreateTime),
			},
		}

		body, _ := mq.FromMsgp(msg)

		if err := exchange.BC.Send(ctx, bc.MQ_WX_TOC_BC_PAYMENT, body); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
	}
}

func BusiCircleRefund(exchange BcExchange) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := wechat.V3ParseNotify(r)
		if err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
		ca := BusinessCircleRefundSource{}
		if err := req.DecryptCipherTextToStruct(exchange.ApiV3Key, &ca); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}

		msg := &msg.BusinessCircleRefund{
			AppID:                  ca.AppID,
			OpenID:                 ca.OpenID,
			RefundTime:             parse(ca.RefundTime),
			PayAmount:              ca.PayAmount,
			RefundAmount:           ca.RefundAmount,
			TransactionID:          ca.TransactionID,
			RefundID:               ca.RefundID,
			BusinessCircleShopBase: msg.BusinessCircleShopBase(ca.BusinessCircleShopBase),
			BusinessCircleMsg: msg.BusinessCircleMsg{
				MsgID:      req.Id,
				Summary:    req.Summary,
				MsgType:    req.EventType,
				CreateTime: parse(req.CreateTime),
			},
		}

		ctx := r.Context()
		body, _ := mq.FromMsgp(msg)
		if err := exchange.BC.Send(ctx, bc.MQ_WX_TOC_BC_REFUND, body); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
	}
}

func WxBCNotify(exchange BcExchange) http.Handler {
	wx := chi.NewRouter()

	wx.Post("/auth", BusiCircleAuth(exchange))
	wx.Post("/payment", BusiCirclePayment(exchange))
	wx.Post("/refund", BusiCircleRefund(exchange))

	wx.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	r := chi.NewRouter()
	r.Mount("/wx", wx)

	return r
}
