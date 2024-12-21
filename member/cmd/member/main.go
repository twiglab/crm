package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/member/orm"
	"github.com/twiglab/crm/member/rpc/gql"
)

const DATABASE_URL = "user=crm password=cRm9ijn)OKM host=pipi.dev port=5432 database=crm sslmode=disable"

func main() {
	db, err := orm.FromURL(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	client := orm.OpenClient(db)

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gqlrpc", gql.New(client))
	log.Fatal(http.ListenAndServe(":10008", mux))
}
