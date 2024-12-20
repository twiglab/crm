package main

import (
	"context"
	"log"

	"github.com/twiglab/crm/member/mq"
	"github.com/twiglab/crm/wechat/pkg/bc"
)

func main() {

	conn, err := mq.Dial("amqp://erp:erpPass123@pipi.dev:5672/")

	if err != nil {
		log.Fatal(err)
	}

	q := &mq.RabbitMQ{
		Conn:      conn,
		QueueName: "q.member.auth",
		Exchange:  bc.MQ_BC_EXCHANGE_NAME,
		BindKey:   bc.MQ_WX_TOC_BC_AUTH,
	}

	ch, err := q.Recieve(context.Background(), &mq.MemberAuthReciverHandle{})
	if err != nil {
		log.Fatal(err)
	}

	<-ch

}
