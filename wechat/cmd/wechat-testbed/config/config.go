package config

import (
	"context"
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
	conn, err := mq.Dial(c.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

type App struct {
	ID         string     `yaml:"id" mapstructure:"id"`
	MQ         MQ         `yaml:"mq" mapstructure:"mq"`
	BcExchange BcExchange `yaml:"bc-exchange" mapstructure:"bc-exchange"`
	Web        Web        `yaml:"web" mapstructure:"web"`
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
