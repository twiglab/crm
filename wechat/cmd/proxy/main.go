package main

import (
	"context"
	"github.com/twiglab/crm/wechat/bc"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/wechat/rpc/gql"
)

func main() {
	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/rpc", gql.New(context.Background()))
	mux.Mount("/notify", notifyRouter())
	log.Fatal(http.ListenAndServe(":10009", mux))
}

func notifyRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/authorization", bc.BusiCircleAuth())
	r.Post("/payment", bc.BusiCirclePayment())
	r.Post("/refund", bc.BusiCircleRefund())
	return r
}
