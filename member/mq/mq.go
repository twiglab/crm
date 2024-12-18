package mq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
)

func Dial(url string) (*amqp.Connection, error) {
	return amqp.Dial(url)
}

type RecieverHandler interface {
	RecieveDelivery(ctx context.Context, delivery amqp.Delivery)
}

type MemberAuthReciverHandle struct {
}

func (h *MemberAuthReciverHandle) RecieveDelivery(ctx context.Context, delivery amqp.Delivery) {
	b := &msg.BusinessCircleAuthor{}
	if _, err := b.UnmarshalMsg(delivery.Body); err != nil {
		return
	}

	fmt.Println(b)

	_ = delivery.Ack(false)
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
}

// 话题模式接受消息
// 要注意key,规则
// 其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
// 匹配 kuteng.* 表示匹配 kuteng.hello, kuteng.hello.one需要用kuteng.#才能匹配到
func (r *RabbitMQ) Recieve(ctx context.Context, h RecieverHandler) (chan struct{}, error) {
	var err error
	if r.channel, err = r.Conn.Channel(); err != nil {
		return nil, err
	}
	//1.试探性创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		amqp.ExchangeTopic,
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil)

	if err != nil {
		return nil, err
	}

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"memberAuth",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	forever := make(chan struct{})

	go func() {
		for {
			select {
			case delivery := <-messges:
				h.RecieveDelivery(ctx, delivery)
			case <-forever:
				return
			}
		}
	}()

	return forever, nil

}
