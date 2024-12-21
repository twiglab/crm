package mq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type MQWapper struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	memberQ amqp.Queue
	pointQ  amqp.Queue
	config  *MQConfig
}

type MQConfig struct {
	URL             string
	ExchangeName    string
	MemberQueueName string
	MemberBindKey   string
	PointQueueName  string
	PointBindKey    string
	Timeout         time.Duration
}

func NewMQWapper(config *MQConfig) (*MQWapper, error) {
	// 连接到AMQP服务器
	conn, err := amqp.Dial(config.URL)
	if err != nil {
		return nil, fmt.Errorf("无法连接到服务器: %v", err)
	}
	// 创建一个通道
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("无法打开通道: %v", err)
	}
	// 交换机
	if err = ch.ExchangeDeclare(
		config.ExchangeName, // name
		amqp.ExchangeTopic,  // type
		true,                // durable
		true,                // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	); err != nil {
		return nil, fmt.Errorf("ExchangeDeclare Fail, err: %v", err)
	}
	// 会员队列
	memQueue, err := ch.QueueDeclare(config.MemberQueueName, true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("QueueDeclare Member Fail, err: %v", err)
	}
	// 积分队列
	pointQueue, err := ch.QueueDeclare(config.PointQueueName, true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("QueueDeclare Member Fail, err: %v", err)
	}
	// 会员绑定Exchange
	err = ch.QueueBind(memQueue.Name, config.MemberBindKey, config.ExchangeName, false, nil)
	if err != nil {
		return nil, fmt.Errorf("QueueBind Member Fail, err: %v", err)
	}
	// 积分绑定Exchange
	err = ch.QueueBind(pointQueue.Name, config.PointBindKey, config.ExchangeName, false, nil)
	if err != nil {
		return nil, fmt.Errorf("QueueBind Point Fail, err: %v", err)
	}
	return &MQWapper{conn: conn, channel: ch, memberQ: memQueue, pointQ: pointQueue, config: config}, nil
}

func (ins *MQWapper) SendMemberMessage(ctx context.Context, msg amqp.Publishing) error {
	return ins.sendMessage(ctx, ins.config.MemberBindKey, msg)
}

func (ins *MQWapper) SendPointMessage(ctx context.Context, msg amqp.Publishing) error {
	return ins.sendMessage(ctx, ins.config.PointBindKey, msg)
}

func (ins *MQWapper) sendMessage(ctx context.Context, routingKey string, msg amqp.Publishing) error {
	octx, cancel := context.WithTimeout(ctx, ins.config.Timeout)
	defer cancel()

	err := ins.channel.PublishWithContext(octx,
		ins.config.ExchangeName, // exchange
		routingKey,              // routing key
		false,                   // mandatory
		false,                   // immediate
		msg,
	)
	return err
}
