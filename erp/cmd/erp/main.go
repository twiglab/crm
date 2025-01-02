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
	"github.com/twiglab/crm/erp/cmd/erp/config"
	"github.com/twiglab/crm/erp/rpc/gql"
	"github.com/twiglab/crm/psdk/conf"
)

func main() {
	cfg := config.App{}
	configCtx := conf.WithContext(context.Background())
	_ = configCtx.ReadInConfig()
	_ = configCtx.Unmarshal(&cfg)

	fmt.Println(cfg)

	client := config.EntClient(cfg.DB)

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gqlrpc", gql.New(client))

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
