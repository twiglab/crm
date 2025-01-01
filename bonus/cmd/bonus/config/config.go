package config

import (
	"context"
	"log"

	"github.com/twiglab/crm/bonus/mq"
	"github.com/twiglab/crm/bonus/orm"
	"github.com/twiglab/crm/bonus/orm/ent"

	"github.com/twiglab/crm/wechat/pkg/bc"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Web struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}
type DB struct {
	DSN string `yaml:"dsn" mapstructure:"dsn"`
}

type GqlRpc struct {
	RpcUrl string `mapstructure:"rpc-url"`
}

type App struct {
	ID        string `yaml:"id" mapstructure:"id"`
	Web       Web    `yaml:"web" mapstructure:"web"`
	MQ        MQ     `yaml:"mq" mapstructure:"mq"`
	DB        DB     `yaml:"db" mapstructure:"db"`
	MemberCli GqlRpc `yaml:"member-cli" mapstructure:"member-cli"`
	BcCli     GqlRpc `yaml:"bc-cli" mapstructure:"bc-cli"`
}

func PaymentQueue(conn *amqp.Connection) *mq.RabbitMQ {
	q := &mq.RabbitMQ{
		Conn:      conn,
		QueueName: "q.member.auth",
		Exchange:  bc.MQ_BC_EXCHANGE_NAME,
		BindKey:   bc.MQ_WX_TOC_BC_PAYMENT,
	}
	return q
}

func RefundQueue(conn *amqp.Connection) *mq.RabbitMQ {
	q := &mq.RabbitMQ{
		Conn:      conn,
		QueueName: "q.member.auth",
		Exchange:  bc.MQ_BC_EXCHANGE_NAME,
		BindKey:   bc.MQ_WX_TOC_BC_REFUND,
	}
	return q
}

type MQ struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

func MQConn(c MQ) *amqp.Connection {
	conn, err := amqp.Dial(c.Addr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func EntClient(dbc DB) *ent.Client {
	db, err := orm.FromURL(context.Background(), dbc.DSN)
	if err != nil {
		log.Fatal(err)
	}

	return orm.OpenClient(db)
}

/*
func WxBcLie(rpc GqlRpc) *bc.WxBcCli {
	retur nil
}
*/
