package config

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/web"
)

type WebServerConfig struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

func (c WebServerConfig) Create(ctx context.Context) *http.Server {
	return web.NewHttpServer(ctx, c.Addr, nil)
}

type BcExchangeConfig struct {
	ExchangeName string
}

func (c BcExchangeConfig) Create(conn *amqp091.Connection) *mq.MQ {
	q := mq.New(slog.Default(), 5*time.Second)
	if err := q.BuildWith(conn, c.ExchangeName); err != nil {
		log.Fatal(err)
	}
	return q
}

type MQConfig struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

func (c MQConfig) Create() *amqp091.Connection {
	conn, err := mq.Dial(c.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

type WechatConfig struct {
	AppId     string `yaml:"app_id" mapstructure:"app_id"`
	AppSecret string `yaml:"app_secret" mapstructure:"app_secret"`

	MchId      string `yaml:"mcd_id" mapstructure:"mcd_id"`
	SerialNo   string `yaml:"serial_no" mapstructure:"serial_no"`
	APIKey     string `yaml:"api_key" mapstructure:"api_key"` // 微信商户平台—>账户设置—>API安全—>设置APIv3密钥
	PrivateKey string `yaml:"private_key" mapstructure:"private_key"`
}

type AppConfig struct {
	APPID    string       `yaml:"appid" mapstructure:"appid"`
	MQConfig MQConfig     `yaml:"mq" mapstructure:"mq"`
	Wechat   WechatConfig `yaml:"wechat" mapstructure:"wechat"`
	BcExchangeConfig
}

func InitConfig(config *AppConfig) {
	viper.SetConfigName("notify-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)

	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal(err)
	}
}
