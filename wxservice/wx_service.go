package wxservice

import (
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

var (
	wxMiniProg *miniprogram.MiniProgram
)

func InitService(AppID, AppSecret string) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     AppID,
		AppSecret: AppSecret,
		Cache:     memory,
	}
	wxMiniProg = wc.GetMiniProgram(cfg)
}

func WeChatLogin(jsCode string) (*LoginResp, error) {
	if wxMiniProg == nil {
		return nil, errors.New("wxMiniProg not init")
	}
	return login(jsCode)
}
