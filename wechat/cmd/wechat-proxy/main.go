package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/it512/box"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twiglab/crm/psdk/conf"
	"github.com/twiglab/crm/psdk/webx"
	"github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/cmd/wechat-proxy/config"
	"github.com/twiglab/crm/wechat/sns"

	"github.com/twiglab/crm/wechat/rpc/gql"
)

func main() {
	cfg := &config.App{}

	c := conf.WithContext(context.Background())
	_ = c.ReadInConfig()
	_ = c.Unmarshal(cfg)

	fmt.Println(cfg)

	prog := config.CreateMiniProfram(cfg.Wechat)

	auth := sns.NewAuth(prog.GetAuth())

	ctx := box.WithContext(c)
	ctx.Put(sns.AuthKey, auth)

	conn := config.CreateMQConn(cfg.MQ)
	xmq := config.CreateBcXMQ(conn)

	bcc := bc.BcExchange{BC: xmq, ApiV3Key: cfg.Wechat.APIKey}

	// ------------------
	mux := webx.Root()
	mux.Mount("/gqlrpc", gql.New(auth))
	mux.Mount("/notify", bc.WxBCNotify(bcc))

	svr := &http.Server{
		Addr:        cfg.Web.Addr,
		IdleTimeout: 10 * time.Second,
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
