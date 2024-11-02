package gql

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/twiglab/crm/wechat/rpc/gql/graph"
)

func New(ctx context.Context) chi.Router {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r := chi.NewRouter()
	r.Handle("/", playground.ApolloSandboxHandler("pg", "/rpc/gql"))
	r.Handle("/gql", srv)

	return r
}
