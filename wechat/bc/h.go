package bc

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/twiglab/crm/wechat/mq"
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

func BusiCirclePayment() http.HandlerFunc {
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
	}
}

func BusiCircleRefund() http.HandlerFunc {
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
	}
}

type Xx struct {
	Q        *mq.MQ
	WxApiKey string
}

func WxBCNotify() http.Handler {
	wx := chi.NewRouter()

	wx.Post("/auth", BusiCircleAuth())
	wx.Post("/payment", BusiCirclePayment())
	wx.Post("/refund", BusiCircleRefund())

	r := chi.NewRouter()
	r.Mount("/wx", wx)

	return r
}
