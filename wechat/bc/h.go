package bc

import (
	"github.com/twiglab/crm/wechat/mq"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
	"github.com/twiglab/crm/wechat/web"
)

func parse(s string) time.Time {
	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Now()
	}
	return x

}

func BusiCircleAuth(x Xx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := wechat.V3ParseNotify(r)
		if err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
		ca := BusinessCircleAuthorSource{}
		if err := req.DecryptCipherTextToStruct(x.WxApiKey, &ca); err != nil {
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

		ctx := r.Context()
		x.Q.SendMarshaler(ctx, msg, bc.MQ_WX_TOC_BC_AUTH)
	}
}

func BusiCirclePayment(x Xx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := wechat.V3ParseNotify(r)
		if err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
		ca := BusinessCirclePaymentSource{}
		if err := req.DecryptCipherTextToStruct("", &ca); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}

		timeEnd, _ := time.Parse(time.RFC3339, ca.TimeEnd)
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

		ctx := r.Context()
		x.Q.SendMarshaler(ctx, msg, bc.MQ_WX_TOC_BC_PAYMENT)
	}
}

func BusiCircleRefund(x Xx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := wechat.V3ParseNotify(r)
		if err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
		ca := BusinessCircleRefundSource{}
		if err := req.DecryptCipherTextToStruct("", &ca); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}

		refundTime, _ := time.Parse(time.RFC3339, ca.RefundTime)
		msg := &msg.BusinessCircleRefund{
			AppID:                  ca.AppID,
			OpenID:                 ca.OpenID,
			RefundTime:             refundTime,
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
		x.Q.SendMarshaler(ctx, msg, bc.MQ_WX_TOC_BC_REFUND)
	}
}

type Xx struct {
	Q        *mq.MQ
	WxApiKey string
}

func WxBCNotify() http.Handler {
	wx := chi.NewRouter()

	wx.Post("/auth", BusiCircleAuth(Xx{}))
	wx.Post("/payment", BusiCirclePayment(Xx{}))
	wx.Post("/refund", BusiCircleRefund(Xx{}))

	r := chi.NewRouter()
	r.Mount("/wx", wx)

	return r
}
