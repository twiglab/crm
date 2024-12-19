package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/it512/box"
	busiCircle "github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/web"
	"log"
	"os"
)

const (
	SERVER_ADDR       = "SERVER_ADDR"
	MQ_ADDR           = "MQ_ADDR"
	WECHAT_MCHID      = "WECHAT_APPSECRET"
	WECHAT_SERIALNO   = "WECHAT_SERIALNO"
	WECHAT_APIKEY     = "WECHAT_APIKEY"
	WECHAT_PRIVATEKEY = "WECHAT_PRIVATEKEY"
)

func main() {
	serverAddr := os.Getenv(SERVER_ADDR)
	mqAddr := os.Getenv(MQ_ADDR)
	mchId := os.Getenv(WECHAT_MCHID)
	serialNo := os.Getenv(WECHAT_SERIALNO)
	// 微信商户平台—>账户设置—>API安全—>设置APIv3密钥
	apiKey := os.Getenv(WECHAT_APIKEY)
	privateKey := os.Getenv(WECHAT_PRIVATEKEY)

	if err := busiCircle.InitWeChatClient(mchId, serialNo, apiKey, privateKey); err != nil {
		panic(err)
	}

	if err := mq.InitMQ(mqAddr); err != nil {
		panic(err)
	}

	ctx := box.Background()

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/wxnotify", busiCircle.WxBCNotify())

	svr := web.NewHttpServer(ctx, serverAddr, mux)
	log.Fatal(web.RunServer(ctx, svr))
}
