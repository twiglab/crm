package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/bonus/balance"
	"github.com/twiglab/crm/bonus/cmd/bonus/config"
	"github.com/twiglab/crm/bonus/gql"
	"github.com/twiglab/crm/bonus/mq"
	"github.com/twiglab/crm/psdk/conf"
)

func main() {
	cfg := config.App{}
	configCtx := conf.WithContext(context.Background())
	_ = configCtx.ReadInConfig()
	_ = configCtx.Unmarshal(&cfg)

	fmt.Println(cfg)

	client := config.EntClient(cfg.DB)
	conn := config.MQConn(cfg.MQ)

	paymentQ := config.PaymentQueue(conn)
	refundQ := config.RefundQueue(conn)

	mbCli := config.MbCli(cfg.MemberCli)
	bcCli := config.BcCli(cfg.BcCli)

	b := &balance.Balance{Client: client}

	ctx, cancelFn := context.WithCancel(configCtx)
	defer cancelFn()

	if err := paymentQ.Recieve(ctx, &mq.BcPaymentHandle{Client: client, MemberCli: mbCli, BcClient: bcCli}); err != nil {
		log.Fatal(err)
	}

	if err := refundQ.Recieve(ctx, &mq.BcRefundHandle{Client: client, MemberCli: mbCli}); err != nil {
		log.Fatal(err)
	}

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gql", gql.GQLRouter(gql.GQLConf{Balance: b}))

	svr := &http.Server{
		Addr:        cfg.Web.Addr,
		IdleTimeout: 10 * time.Second,
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return configCtx },
	}

	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
