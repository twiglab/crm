package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/twiglab/crm/wechat/cmd/wechat-testbed/config"
	"github.com/twiglab/crm/wechat/cmd/wechat-testbed/gql/graph"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	cfg := config.App{}
	config.InitConfig(&cfg)

	mqc := cfg.MQ.Create()
	q := cfg.BcExchange.Create(mqc)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Q: q}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.ApolloSandboxHandler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":10007", nil))
}
