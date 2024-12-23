package gql

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/twiglab/crm/wechat/rpc/gql/graph"
	"github.com/twiglab/crm/wechat/sns"
)

func New(a *sns.Auth) chi.Router {
	es := graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{Auth: a},
		},
	)

	srv := NewSvr(es)

	r := chi.NewRouter()
	r.Handle("/", playground.ApolloSandboxHandler("rpc", "/gqlrpc/query"))
	r.Handle("/query", srv)

	return r
}
