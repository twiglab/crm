package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/it512/box"
	"github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/cmd/wechat-bc-notify/config"
	"github.com/twiglab/crm/wechat/web"
)

func main() {
	cfg := &config.App{}
	config.InitConfig(cfg)

	conn := cfg.MQ.Create()
	exchange := cfg.BcExchange.Create(conn)

	bcc := bc.BcExchange{BC: exchange, ApiV3Key: cfg.Wechat.APIKey}

	ctx := box.Background()

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/notify", bc.WxBCNotify(bcc))

	svr := cfg.Web.Create(ctx)

	if err := web.RunServer2(ctx, svr, mux); err != nil {
		log.Fatal(err)
	}
}
