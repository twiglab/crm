package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/it512/box"
	bcdef "github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/cmd/wechat-test/config"
	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
	"github.com/twiglab/crm/wechat/web"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.App{}
	config.InitConfig(&cfg)
	conn := cfg.MQ.Create()
	exchange := cfg.BcExchange.Create(conn)

	bcc := bcdef.BcExchange{BC: exchange, ApiV3Key: cfg.Wechat.APIKey}

	ctx := box.Background()

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/member_test", SendAuthMQTest(bcc))

	svr := cfg.Web.Create(ctx)

	if err := web.RunServer2(ctx, svr, mux); err != nil {
		log.Fatal(err)
	}
}

func SendAuthMQTest(exchange bcdef.BcExchange) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := &msg.BusinessCircleAuthor{
			MchID:    getQueryValue(r, "mch_id"),
			OpenID:   getQueryValue(r, "open_id"),
			Code:     getQueryValue(r, "code"),
			AuthType: getQueryValue(r, "auth_type"),

			BusinessCircleMsg: msg.BusinessCircleMsg{
				MsgID:      getQueryValue(r, "id"),
				Summary:    getQueryValue(r, "summmary"),
				MsgType:    getQueryValue(r, "msg_type"),
				CreateTime: time.Now(),
			},
		}

		body, _ := mq.MsgpMsg(msg)

		ctx := r.Context()

		if err := exchange.BC.Send(ctx, bc.MQ_WX_TOC_BC_AUTH, body); err != nil {
			_ = web.JsonTo(http.StatusInternalServerError, &wechat.V3NotifyRsp{Code: gopay.FAIL, Message: err.Error()}, w)
			return
		}
	}
}

func getQueryValue(r *http.Request, key string, defVal ...string) string {
	val := r.URL.Query().Get(key)
	if len(val) > 0 {
		return val
	}
	if len(defVal) > 0 {
		return defVal[0]
	}
	return key + "_test"
}
