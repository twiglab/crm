package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/twiglab/crm/card/orm/ent"
	"github.com/twiglab/crm/card/rpc/gql/graph"
)

func New(client *ent.Client) chi.Router {

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Client: client}}))
	r := chi.NewRouter()
	r.Handle("/", playground.ApolloSandboxHandler("pg", "/gqlrpc/query"))
	r.Handle("/query", srv)
	return r
}
