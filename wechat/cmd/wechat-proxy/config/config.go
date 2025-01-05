package config

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/web"
)

type Web struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

func Create(ctx context.Context, c Web) *http.Server {
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

func CreateMiniProfram(c Wechat) *miniprogram.MiniProgram {
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
	MQ     MQ     `yaml:"mq" mapstructure:"mq"`
}

func CreateBcXMQ(conn *amqp.Connection) *mq.XMQ {
	q, err := mq.BcMQ(conn, slog.Default())
	if err != nil {
		log.Fatal(err)
	}
	return q
}

type MQ struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

func CreateMQConn(c MQ) *amqp.Connection {
	conn, err := mq.Dial(c.Addr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
