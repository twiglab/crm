package mq

import (
	"context"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tinylib/msgp/msgp"
)

const MSGPACK_MIME = "application/msgpack"

var mqInstance *MQ

func Instance() *MQ {
	return mqInstance
}

func InitMQ(uri string) error {
	mqInstance = New(slog.Default(), 30*time.Second)
	conn, err := Dial(uri)
	if err != nil {
		return err
	}
	if err = mqInstance.BuildWith(conn, bc.MQ_BC_EXCHANGE_NAME); err != nil {
		return err
	}
	return nil
}

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

	binit bool
}

func New(logger *slog.Logger, timeout time.Duration) *MQ {
	return &MQ{
		timeout: timeout,
		logger:  logger,
		binit:   false,
	}
}

func (q *MQ) BuildWith(conn *amqp.Connection, exanme string) error {
	if q.binit {
		return nil
	}

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
		true,               // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)

	q.binit = true
	return err
}

func (mq *MQ) Send(ctx context.Context, routingKey string, msg amqp.Publishing) error {
	octx, cancel := context.WithTimeout(ctx, mq.timeout)
	defer cancel()

	err := mq.ch.PublishWithContext(octx,
		mq.exchangeName, // exchange
		routingKey,      // routing key
		false,           // mandatory
		false,           // immediate
		msg,
	)
	return err
}

func (mq *MQ) Close() error {
	return mq.ch.Close()
}
