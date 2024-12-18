package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
)

func main() {
	fmt.Println("请黄总再此处写代码，提交并保证运行正确！多谢！")

	conn, err := mq.Dial("amqp://admin:rabbitmq123456.PWD.rabbitmq@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	q := mq.New(slog.Default(), 30*time.Second)
	if err = q.BuildWith(conn, bc.MQ_BC_EXCHANGE_NAME); err != nil {
		log.Fatal(err)
	}

	a := &msg.BusinessCircleAuthor{OpenID: "xxxxx"}
	msg, _ := mq.MsgpMsg(a)

	if err := q.Send(context.Background(), bc.MQ_WX_TOC_BC_AUTH, msg); err != nil {
		log.Fatal(err)
	}
}
