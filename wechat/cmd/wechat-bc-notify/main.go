package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/it512/box"
	busiCircle "github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/config"
	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/web"
	"log"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	cfg := config.GetConfig()

	if err := mq.InitMQ(cfg.MQ.Addr); err != nil {
		panic(err)
	}

	//if err := busiCircle.InitWeChatClient(cfg.BC.MchId, cfg.BC.SerialNo, cfg.BC.APIKey, cfg.BC.PrivateKey); err != nil {
	//	panic(err)
	//}

	ctx := box.Background()

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/wxnotify", busiCircle.WxBCNotify())

	svr := web.NewHttpServer(ctx, cfg.App.Addr, mux)
	log.Fatal(web.RunServer(ctx, svr))
}
