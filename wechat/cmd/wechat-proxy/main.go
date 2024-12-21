package main

import (
	"github.com/it512/box"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/twiglab/crm/wechat/config"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twiglab/crm/wechat/sns"
	"github.com/twiglab/crm/wechat/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/wechat/rpc/gql"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	cfg := config.GetConfig()

	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	miniCfg := &miniConfig.Config{
		AppID:     cfg.Wechat.AppId,
		AppSecret: cfg.Wechat.AppSecret,
		Cache:     memory,
	}
	prog := wc.GetMiniProgram(miniCfg)

	auth := sns.NewAuth(prog.GetAuth())

	ctx := box.Background()
	ctx.Put(sns.AuthKey, auth)

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gqlrpc", gql.New(auth))

	svr := web.NewHttpServer(ctx, cfg.App.Addr, mux)
	log.Fatal(web.RunServer(ctx, svr))
}
