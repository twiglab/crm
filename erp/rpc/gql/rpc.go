package gql

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/twiglab/crm/erp/orm/ent"
	"github.com/twiglab/crm/erp/rpc/gql/graph"
	"github.com/twiglab/crm/psdk/gqlx"
)

func New(client *ent.Client) chi.Router {
	srv := gqlx.NewGQLHandler(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Client: client}}))
	r := chi.NewRouter()
	r.Handle("/", playground.ApolloSandboxHandler("pg", "/gqlrpc/query"))
	r.Handle("/query", srv)
	return r
}
