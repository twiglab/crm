package config

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/twiglab/crm/member/mq"
	"github.com/twiglab/crm/member/orm"
	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/wechat/pkg/bc"
)

type Web struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}
type DB struct {
	DSN string `yaml:"dsn" mapstructure:"dsn"`
}

type App struct {
	ID  string `yaml:"id" mapstructure:"id"`
	Web Web    `yaml:"web" mapstructure:"web"`
	MQ  MQ     `yaml:"mq" mapstructure:"mq"`
	DB  DB     `yaml:"db" mapstructure:"db"`
}

func AuthQueue(conn *amqp.Connection) *mq.RabbitMQ {
	q := &mq.RabbitMQ{
		Conn:      conn,
		QueueName: "q.member.auth",
		Exchange:  bc.MQ_BC_EXCHANGE_NAME,
		BindKey:   bc.MQ_WX_TOC_BC_AUTH,
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
