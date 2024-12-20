package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/member/orm"
	"github.com/twiglab/crm/member/rpc/gql"
)

const DATABASE_URL = "user=crm password=crm0okm_PL< host=it9i.com port=15432 database=crm sslmode=disable"

func main() {
	db, err := orm.FromURL(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.Ping())

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/rpc", gql.New(context.Background()))
	log.Fatal(http.ListenAndServe(":10009", mux))
}
