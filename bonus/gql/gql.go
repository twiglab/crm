package gql

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/twiglab/crm/bonus/balance"
	"github.com/twiglab/crm/bonus/gql/graph"
	"github.com/twiglab/crm/bonus/orm"
	"github.com/twiglab/crm/psdk/gqlx"
)

type GQLConf struct {
	Balance *balance.Balance
	Items   *orm.Items
}

func GQLRouter(conf GQLConf) chi.Router {
	srv := gqlx.NewGQLHandler(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{Balance: conf.Balance, Items: conf.Items},
		},
	))

	r := chi.NewRouter()
	r.Handle("/", playground.ApolloSandboxHandler("pg", "/gql/query"))
	r.Handle("/query", srv)

	return r
}
