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
	WECHAT_APPID      = "WECHAT_APPID"
	WECHAT_APPSECRET  = "WECHAT_APPSECRET"
	WECHAT_MCHID      = "WECHAT_APPSECRET"
	WECHAT_SERIALNO   = "WECHAT_SERIALNO"
	WECHAT_APIKEY     = "WECHAT_APIKEY"
	WECHAT_PRIVATEKEY = "WECHAT_PRIVATEKEY"
)

func main() {
	appId := os.Getenv(WECHAT_APPID)
	appSecret := os.Getenv(WECHAT_APPSECRET)
	mchId := os.Getenv(WECHAT_MCHID)
	serialNo := os.Getenv(WECHAT_SERIALNO)
	apiKey := os.Getenv(WECHAT_APIKEY)
	privateKey := os.Getenv(WECHAT_PRIVATEKEY)

	sns.InitWeChatClient(mchId, serialNo, apiKey, privateKey)

	cfg := miniConfig.Config{
		AppID:     appId,
		AppSecret: appSecret,
	}
	sns.InitMini(&cfg, cache.NewMemory())

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/rpc", gql.New(context.Background()))
	mux.Mount("/wxnotify", bc.WxBCNotify())
	log.Fatal(http.ListenAndServe(":10009", mux))
}
