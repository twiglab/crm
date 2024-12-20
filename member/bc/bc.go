package bc

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
)

type MemberAuthReciverHandle struct {
	Client ent.Client
}

func (h *MemberAuthReciverHandle) RecieveDelivery(ctx context.Context, delivery amqp.Delivery) {
	b := &msg.BusinessCircleAuthor{}
	if _, err := b.UnmarshalMsg(delivery.Body); err != nil {
		return
	}

	if err := h.Handle(b); err != nil {
		return
	}

	_ = delivery.Ack(false)
}

func (h *MemberAuthReciverHandle) Handle(msg *msg.BusinessCircleAuthor) error {
	return nil
}
