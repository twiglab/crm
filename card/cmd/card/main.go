package card

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/twiglab/crm/card/orm"
	"github.com/twiglab/crm/card/rpc/gql"
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
