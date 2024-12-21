package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/it512/box"
	"github.com/twiglab/crm/wechat/bc"
	busiCircle "github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/cmd/wechat-bc-notify/config"
	"github.com/twiglab/crm/wechat/web"
)

func main() {
	cfg := config.AppConfig{}
	config.InitConfig(&cfg)

	conn := cfg.MQConfig.Create()
	exchange := cfg.BcExchangeConfig.Create(conn)

	bcc := bc.BcExchange{BC: exchange, ApiV3Key: cfg.Wechat.APIKey}

	ctx := box.Background()

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/notify", busiCircle.WxBCNotify(bcc))

	svr := web.NewHttpServer(ctx, cfg.App.Addr, mux)
	log.Fatal(web.RunServer(ctx, svr))
}
