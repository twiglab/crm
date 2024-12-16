package member

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnErr(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RecieverHandler interface {
	RecieveDelivery(ctx context.Context, delivery amqp.Delivery)
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
func (r *RabbitMQ) Recieve(ctx context.Context, h RecieverHandler) chan struct{} {
	//1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "Failed to declare an exchange")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	failOnErr(err, "Failed to declare a queue")

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		r.Key,
		r.Exchange,
		false,
		nil)
	failOnErr(err, "Failed to declare a queue")

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "Failed to declare a queue")

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

	return forever

}
