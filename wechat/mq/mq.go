package mq

import (
	"context"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tinylib/msgp/msgp"
)

const MSGPACK_MIME = "application/msgpack"

func Dial(url string) (*amqp.Connection, error) {
	return amqp.Dial(url)
}

func MsgpMsg(m msgp.Marshaler) (amqp.Publishing, error) {
	bs, err := m.MarshalMsg(nil)
	if err != nil {
		return amqp.Publishing{}, err
	}

	return amqp.Publishing{Body: bs, ContentType: MSGPACK_MIME}, nil
}

type MQ struct {
	timeout time.Duration
	logger  *slog.Logger

	conn         *amqp.Connection
	exchangeName string

	ch *amqp.Channel
}

func New(logger *slog.Logger, timeout time.Duration) *MQ {
	return &MQ{
		timeout: timeout,
		logger:  logger,
	}
}

func (q *MQ) BuildWith(conn *amqp.Connection, exanme string) error {
	var err error

	q.conn = conn
	q.exchangeName = exanme

	if q.ch, err = q.conn.Channel(); err != nil {
		return err
	}

	err = q.ch.ExchangeDeclare(
		q.exchangeName,     // name
		amqp.ExchangeTopic, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		return err
	}

	return q.ch.Confirm(true)
}

func (mq *MQ) Send(ctx context.Context, routingKey string, msg amqp.Publishing) error {
	octx, cancel := context.WithTimeout(ctx, mq.timeout)
	defer cancel()

	err := mq.ch.PublishWithContext(octx,
		mq.exchangeName, // exchange
		routingKey,      // routing key
		true,            // mandatory
		false,           // immediate
		msg,
	)
	return err
}

func (mq *MQ) Close() error {
	return mq.ch.Close()
}
