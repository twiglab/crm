package config

import (
	"context"
	"log"
	"net/http"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/spf13/viper"
	"github.com/twiglab/crm/wechat/web"
)

type Web struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

func (c Web) Create(ctx context.Context) *http.Server {
	return web.NewHttpServer(ctx, c.Addr, nil)
}

type Wechat struct {
	AppID     string `yaml:"app_id" mapstructure:"app_id"`
	AppSecret string `yaml:"app_secret" mapstructure:"app_secret"`

	MchId      string `yaml:"mcd_id" mapstructure:"mcd_id"`
	SerialNo   string `yaml:"serial_no" mapstructure:"serial_no"`
	APIKey     string `yaml:"api_key" mapstructure:"api_key"` // 微信商户平台—>账户设置—>API安全—>设置APIv3密钥
	PrivateKey string `yaml:"private_key" mapstructure:"private_key"`
}

func (c Wechat) CreateMiniProfram() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	miniCfg := &miniConfig.Config{
		AppID:     c.AppID,
		AppSecret: c.AppSecret,
		Cache:     memory,
	}
	return wc.GetMiniProgram(miniCfg)

}

type App struct {
	ID     string `yaml:"id" mapstructure:"id"`
	Wechat Wechat `yaml:"wechat" mapstructure:"wechat"`
	Web    Web    `yaml:"web" mapstructure:"web"`
}

func InitConfig(config any) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)

	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatal(err)
	}
}
