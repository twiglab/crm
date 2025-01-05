package mq

import (
	"context"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/member/orm/ent/member"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
)

type MemberAuthHandle struct {
	Client *ent.Client
}

func (h *MemberAuthHandle) RecieveDelivery(ctx context.Context, delivery amqp.Delivery) {
	b := &msg.BusinessCircleAuthor{}
	if _, err := b.UnmarshalMsg(delivery.Body); err != nil {
		return
	}

	if err := h.Handle(ctx, b); err != nil {
		return
	}

	_ = delivery.Ack(true)
}

func (h *MemberAuthHandle) Handle(ctx context.Context, msg *msg.BusinessCircleAuthor) error {
	q := h.Client.Member.Query()
	mb, err := q.Where(member.WxOpenIDEQ(msg.OpenID)).Only(ctx)
	if ent.IsNotFound(err) {
		return nil
	}

	if err != nil {
		return err
	}

	if strings.EqualFold(mb.BcmbRegMsgID, msg.MsgID) {
		return nil
	}

	u := mb.Update()
	u.SetBcmbCode(msg.Code)
	u.SetBcmbRegMsgID(msg.MsgID)
	u.SetBcmbRegTime(time.Now())
	u.SetLevel(10)
	u.SetSource(1)
	return u.Exec(ctx)
}
