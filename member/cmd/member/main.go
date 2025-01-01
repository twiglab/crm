package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/member/cmd/member/config"
	"github.com/twiglab/crm/member/mq"
	"github.com/twiglab/crm/member/rpc/gql"
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

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	if err := q.Recieve(ctx, &mq.MemberAuthHandle{Client: client}); err != nil {
		log.Fatal(err)
	}

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gqlrpc", gql.New(client))
	log.Fatal(http.ListenAndServe(cfg.Web.Addr, mux))
}
