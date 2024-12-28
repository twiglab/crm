package gql

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/twiglab/crm/psdk/gqlx"
	"github.com/twiglab/crm/wxlogin"
	"github.com/twiglab/crm/wxlogin/gql/graph"
)

func GQLRouter(acli *wxlogin.AuthClient) chi.Router {
	svr := gqlx.NewGQLHandler(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{AuthCli: acli},
		},
	))

	r := chi.NewRouter()
	r.Handle("/", playground.ApolloSandboxHandler("pg", "/gql/query"))
	r.Handle("/query", srv)

	return r
}
