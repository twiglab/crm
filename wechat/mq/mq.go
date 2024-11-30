package mq

import (
	"context"
	"encoding"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/twiglab/crm/wechat/pkg/bc"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type MQ struct {
	ch      *amqp.Channel
	conn    *amqp.Connection
	exname  string
	timeout time.Duration
}

func X(url string) *MQ {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		bc.MQ_BC_EXCHANGE_NAME, // name
		amqp.ExchangeTopic,     // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	return &MQ{ch: ch, conn: conn, exname: "wechat_topic"}
}

func (mq *MQ) Send(ctx context.Context, body []byte, key string) {
	octx, cancel := context.WithTimeout(ctx, mq.timeout)
	defer cancel()

	err := mq.ch.PublishWithContext(octx,
		mq.exname, // exchange
		key,       // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
}

func (mq *MQ) SendMarshaler(ctx context.Context, body encoding.BinaryMarshaler, key string) {
	bs, _ := body.MarshalBinary()
	mq.Send(ctx, bs, key)
}

func (mq *MQ) Close() error {
	mq.ch.Close()
	mq.conn.Close()
	return nil
}
