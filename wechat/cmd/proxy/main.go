package main

import (
	"context"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/twiglab/crm/wechat/sns"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twiglab/crm/wechat/bc"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/wechat/rpc/gql"
)

const (
	WECHAT_APPID     = "WECHAT_APPID"
	WECHAT_APPSECRET = "WECHAT_APPSECRET"
)

func main() {
	cfg := miniConfig.Config{
		AppID:     os.Getenv(WECHAT_APPID),
		AppSecret: os.Getenv(WECHAT_APPSECRET),
	}
	sns.InitMini(&cfg, cache.NewMemory())

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/rpc", gql.New(context.Background()))
	mux.Mount("/wxnotify", bc.WxBCNotify())
	log.Fatal(http.ListenAndServe(":10009", mux))
}
