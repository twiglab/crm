package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/it512/box"
	busiCircle "github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/config"
	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"github.com/twiglab/crm/wechat/web"
	"log"
	"time"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	cfg := config.GetConfig()

	var mqConig = &mq.MQConfig{
		URL:             cfg.MQ.Addr,
		ExchangeName:    bc.MQ_BC_EXCHANGE_NAME,
		MemberQueueName: bc.MQ_BC_QUEUE_MEMBER_NAME,
		MemberBindKey:   bc.MQ_WX_TOC_BC_AUTH,
		PointQueueName:  bc.MQ_BC_QUEUE_POINT_NAME,
		PointBindKey:    bc.MQ_WX_TOC_BC_PAYMENT,
		Timeout:         time.Second * 3,
	}
	mqWapper, err := mq.NewMQWapper(mqConig)
	if err != nil {
		panic(err)
	}
	exchange := &busiCircle.BcExchange{
		BC:       mqWapper,
		ApiV3Key: cfg.BC.APIKey,
	}

	//if err := busiCircle.InitWeChatClient(cfg.BC.MchId, cfg.BC.SerialNo, cfg.BC.APIKey, cfg.BC.PrivateKey); err != nil {
	//	panic(err)
	//}

	ctx := box.Background()

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/wxnotify", busiCircle.WxBCNotify(exchange))

	svr := web.NewHttpServer(ctx, cfg.App.Addr, mux)
	log.Fatal(web.RunServer(ctx, svr))
}
