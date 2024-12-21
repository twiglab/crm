package config

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"github.com/twiglab/crm/wechat/web"
)

type Web struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

func (c Web) Create(ctx context.Context) *http.Server {
	return web.NewHttpServer(ctx, c.Addr, nil)
}

type BcExchange struct {
	Name string `yaml:"name" mapstructure:"name"`
}

func (c BcExchange) Create(conn *amqp.Connection) *mq.MQ {
	q := mq.New(slog.Default(), 5*time.Second)
	if err := q.BuildWith(conn, bc.MQ_BC_EXCHANGE_NAME); err != nil {
		log.Fatal(err)
	}
	return q
}

type MQ struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

func (c MQ) Create() *amqp.Connection {
	fmt.Println(c.Addr)
	conn, err := mq.Dial(c.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

type Wechat struct {
	AppID     string `yaml:"app_id" mapstructure:"app_id"`
	AppSecret string `yaml:"app_secret" mapstructure:"app_secret"`

	MchId      string `yaml:"mcd_id" mapstructure:"mcd_id"`
	SerialNo   string `yaml:"serial_no" mapstructure:"serial_no"`
	APIKey     string `yaml:"api_key" mapstructure:"api_key"` // 微信商户平台—>账户设置—>API安全—>设置APIv3密钥
	PrivateKey string `yaml:"private_key" mapstructure:"private_key"`
}

type App struct {
	MQ         MQ         `yaml:"mq" mapstructure:"mq"`
	Wechat     Wechat     `yaml:"wechat" mapstructure:"wechat"`
	BcExchange BcExchange `yaml:"bc-exchange" mapstructure:"bc-exchange"`
	Web        Web        `yaml:"web" mapstructure:"web"`
}

func InitConfig(config *App) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("cmd/wechat-test/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)

	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatal(err)
	}
}
