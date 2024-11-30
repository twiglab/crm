package main

import (
	"context"
	"log"
	"net/http"

	"github.com/twiglab/crm/wechat/bc"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/wechat/rpc/gql"
)

func main() {
	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/rpc", gql.New(context.Background()))
	mux.Mount("/wxnotify", bc.BCNotify())
	log.Fatal(http.ListenAndServe(":10009", mux))
}
