package mq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MustDial(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

type RecieverHandler interface {
	RecieveDelivery(ctx context.Context, delivery amqp.Delivery)
}

type RabbitMQ struct {
	Conn     *amqp.Connection
	channel  *amqp.Channel
	Exchange string

	QueueName string

	BindKey string
}

// 话题模式接受消息
// 要注意key,规则
// 其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
// 匹配 kuteng.* 表示匹配 kuteng.hello, kuteng.hello.one需要用kuteng.#才能匹配到
func (r *RabbitMQ) Recieve(ctx context.Context, h RecieverHandler) error {
	var err error
	if r.channel, err = r.Conn.Channel(); err != nil {
		return err
	}
	//1.试探性创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
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
		return err
	}

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		r.BindKey,
		r.Exchange,
		false,
		nil)

	if err != nil {
		return err
	}

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	r.RecvMessage(ctx, h, messges)

	return nil
}

func (r *RabbitMQ) RecvMessage(ctx context.Context, h RecieverHandler, message <-chan amqp.Delivery) {
	go func() {
		for {
			select {
			case delivery := <-message:
				h.RecieveDelivery(ctx, delivery)
			case <-ctx.Done():
				return
			}
		}
	}()
}
