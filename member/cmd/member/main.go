package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/twiglab/crm/member/cmd/member/config"
	"github.com/twiglab/crm/member/gql"
	"github.com/twiglab/crm/member/mq"
	rpcgql "github.com/twiglab/crm/member/rpc/gql"

	"github.com/twiglab/crm/psdk/conf"
	"github.com/twiglab/crm/psdk/webx"
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

	mux := webx.Root()
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
