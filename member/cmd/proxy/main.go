package main

import (
	"context"
	"fmt"
	"log"

	"github.com/twiglab/crm/member/mq"
	"github.com/twiglab/crm/member/orm"
	"github.com/twiglab/crm/wechat/pkg/bc"
)

const DATABASE_URL = "user=crm password=cRm9ijn)OKM host=pipi.dev port=5432 database=crm sslmode=disable"

func main() {
	db, err := orm.FromURL(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.Ping())

	client := orm.OpenClient(db)

	conn, err := mq.Dial("amqp://erp:erpPass123@pipi.dev:5672/")

	if err != nil {
		log.Fatal(err)
	}

	recvExitChan := make(chan struct{})

	q := &mq.RabbitMQ{
		Conn:      conn,
		QueueName: "q.member.auth",
		Exchange:  bc.MQ_BC_EXCHANGE_NAME,
		BindKey:   bc.MQ_WX_TOC_BC_AUTH,
		ExitChan:  recvExitChan,
	}

	if err := q.Recieve(context.Background(), &mq.MemberAuthReciverHandle{Client: client}); err != nil {
		log.Fatal(err)
	}

	select {}
}
