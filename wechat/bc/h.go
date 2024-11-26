package bc

import (
	"net/http"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/twiglab/crm/wechat/web"
)

func X() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := wechat.V3ParseNotify(r)
		if err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: "失败"}, w)
			return
		}
		ca := BusinessCircleAuthorSource{}
		if err := req.DecryptCipherTextToStruct("", &ca); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: "失败"}, w)
			return
		}
	}
}
