package main

import (
	"log"

	"github.com/it512/box"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twiglab/crm/wechat/cmd/wechat-proxy/config"
	"github.com/twiglab/crm/wechat/sns"
	"github.com/twiglab/crm/wechat/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/wechat/rpc/gql"
)

func main() {
	cfg := &config.App{}
	config.InitConfig(cfg)

	prog := cfg.Wechat.CreateMiniProfram()

	auth := sns.NewAuth(prog.GetAuth())

	ctx := box.Background()
	ctx.Put(sns.AuthKey, auth)

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gqlrpc", gql.New(auth))

	svr := cfg.Web.Create(ctx)
	log.Fatal(web.RunServer2(ctx, svr, mux))
}
