package main

import (
	"log"
	"net/http"
	"os"

	"github.com/it512/box"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/sns"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/twiglab/crm/wechat/rpc/gql"
)

const (
	WECHAT_APPID     = "WECHAT_APPID"
	WECHAT_APPSECRET = "WECHAT_APPSECRET"
	WECHAT_MCHID     = "WECHAT_APPSECRET"

	WECHAT_SERIALNO   = "WECHAT_SERIALNO"
	WECHAT_APIKEY     = "WECHAT_APIKEY"
	WECHAT_PRIVATEKEY = "WECHAT_PRIVATEKEY"
)

func main() {
	appId := os.Getenv(WECHAT_APPID)
	appSecret := os.Getenv(WECHAT_APPSECRET)
	/*
		mchId := os.Getenv(WECHAT_MCHID)
		serialNo := os.Getenv(WECHAT_SERIALNO)
		apiKey := os.Getenv(WECHAT_APIKEY)
		privateKey := os.Getenv(WECHAT_PRIVATEKEY)
	*/

	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     appId,
		AppSecret: appSecret,
		Cache:     memory,
	}
	prog := wc.GetMiniProgram(cfg)

	auth := sns.NewAuth(prog.GetAuth())

	ctx := box.Background()
	ctx.Put(sns.AuthKey, auth)

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/rpc", gql.New(ctx))
	mux.Mount("/wxnotify", bc.WxBCNotify())
	log.Fatal(http.ListenAndServe(":10009", mux))
}
