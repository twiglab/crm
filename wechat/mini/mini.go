package mini

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

var (
	wxMiniProg *miniprogram.MiniProgram
)

func InitMini(cfg *miniConfig.Config, cache cache.Cache) {
	wc := wechat.NewWechat()
	wc.SetCache(cache)
	wxMiniProg = wc.GetMiniProgram(cfg)
}

func GetAuth() *auth.Auth {
	return wxMiniProg.GetAuth()
}
