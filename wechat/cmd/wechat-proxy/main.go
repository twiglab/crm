package main

import (
	"context"
	"log"

	"github.com/it512/box"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twiglab/crm/psdk/conf"
	"github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/cmd/wechat-proxy/config"
	"github.com/twiglab/crm/wechat/sns"
	"github.com/twiglab/crm/wechat/testbed"
	"github.com/twiglab/crm/wechat/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/wechat/rpc/gql"
)

func main() {
	cfg := &config.App{}

	c := conf.WithContext(context.Background())
	c.ReadInConfig()
	c.Unmarshal(cfg)

	prog := config.CreateMiniProfram(cfg.Wechat)

	auth := sns.NewAuth(prog.GetAuth())

	ctx := box.WithContext(c)
	ctx.Put(sns.AuthKey, auth)

	conn := config.CreateMQConn(cfg.MQ)
	xmq := config.CreateBcXMQ(conn)

	bcc := bc.BcExchange{BC: xmq, ApiV3Key: cfg.Wechat.APIKey}

	// ------------------
	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gqlrpc", gql.New(auth))
	mux.Mount("/notify", bc.WxBCNotify(bcc))
	mux.Mount("/testbed", testbed.TestBed(xmq))

	svr := config.Create(ctx, cfg.Web)
	log.Fatal(web.RunServer2(ctx, svr, mux))
}
