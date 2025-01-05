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

	"github.com/twiglab/crm/member/cmd/member/config"
	"github.com/twiglab/crm/member/gql"
	"github.com/twiglab/crm/member/mq"
	rpcgql "github.com/twiglab/crm/member/rpc/gql"

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

	q := config.AuthQueue(conn)

	ctx, cancelFn := context.WithCancel(configCtx)
	defer cancelFn()

	if err := q.Recieve(ctx, &mq.MemberAuthHandle{Client: client}); err != nil {
		log.Fatal(err)
	}

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gql", gql.New(client))
	mux.Mount("/gqlrpc", rpcgql.New(client))

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
