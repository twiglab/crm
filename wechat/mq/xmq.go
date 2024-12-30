package mq

import (
	"context"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tinylib/msgp/msgp"
	"github.com/twiglab/crm/wechat/pkg/bc"
)

type Message struct {
	Body        []byte
	ContentType string
}

func FromMsgp(m msgp.Marshaler) (Message, error) {
	bs, err := m.MarshalMsg(nil)
	if err != nil {
		return Message{}, err
	}

	return Message{Body: bs, ContentType: MSGPACK_MIME}, nil
}

type XMQ struct {
	Timeout  time.Duration
	Logger   *slog.Logger
	Conn     *amqp.Connection
	Exchange string

	ch *amqp.Channel
}

func BcMQ(conn *amqp.Connection, Logger *slog.Logger) (*XMQ, error) {

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		bc.MQ_BC_EXCHANGE_NAME, // name
		amqp.ExchangeTopic,     // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		return nil, err
	}

	if err := ch.Confirm(true); err != nil {
		return nil, err
	}
	return &XMQ{
		Timeout:  10 * time.Second,
		Logger:   Logger,
		Conn:     conn,
		Exchange: bc.MQ_BC_EXCHANGE_NAME,
		ch:       ch,
	}, nil
}

func (mq *XMQ) Send(ctx context.Context, routingKey string, msg Message) error {
	octx, cancel := context.WithTimeout(ctx, mq.Timeout)
	defer cancel()

	err := mq.ch.PublishWithContext(octx,
		mq.Exchange, // exchangeName
		routingKey,  // routing key
		true,        // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:  msg.ContentType,
			Body:         msg.Body,
			DeliveryMode: 2,
		},
	)
	return err
}

func (mq *XMQ) Close() error {
	return mq.ch.Close()
}
