package mq

import (
	"context"
	"log"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tinylib/msgp/msgp"
)

const MSGPACK_MIME = "application/msgpack"

func MsgpMsg(m msgp.Marshaler) (amqp.Publishing, error) {
	bs, err := m.MarshalMsg(nil)
	if err != nil {
		return amqp.Publishing{}, err
	}

	return amqp.Publishing{Body: bs, ContentType: MSGPACK_MIME}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type MQ struct {
	Conn         *amqp.Connection
	ExchangeName string
	Timeout      time.Duration

	Logger *slog.Logger

	ch *amqp.Channel
}

func New(logger *slog.Logger, timeout time.Duration) *MQ {
	return &MQ{
		Timeout: timeout,
		Logger:  logger,
	}
}

func (q *MQ) BuildWith(conn *amqp.Connection) {
	var err error

	q.Conn = conn

	q.ch, err = q.Conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = q.ch.ExchangeDeclare(
		q.ExchangeName,     // name
		amqp.ExchangeTopic, // type
		true,               // durable
		true,               // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare an exchange")

}

func (q *MQ) Build(url string) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	q.BuildWith(conn)
}

func (mq *MQ) Send(ctx context.Context, routingKey string, msg amqp.Publishing) error {
	octx, cancel := context.WithTimeout(ctx, mq.Timeout)
	defer cancel()

	err := mq.ch.PublishWithContext(octx,
		mq.ExchangeName, // exchange
		routingKey,      // routing key
		false,           // mandatory
		false,           // immediate
		msg,
	)
	return err
}

func (mq *MQ) Close() error {
	mq.ch.Close()
	mq.Conn.Close()
	return nil
}
