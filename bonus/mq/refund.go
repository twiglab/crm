package mq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/twiglab/crm/bonus/mb"
	"github.com/twiglab/crm/bonus/orm/ent"
	"github.com/twiglab/crm/bonus/orm/ent/bonusitem"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
)

type BcRefundHandle struct {
	Client    *ent.Client
	MemberCli *mb.MemberCli
}

func (h *BcRefundHandle) RecieveDelivery(ctx context.Context, delivery amqp.Delivery) {
	b := &msg.BusinessCircleRefund{}
	if _, err := b.UnmarshalMsg(delivery.Body); err != nil {
		return
	}

	if err := h.Handle(ctx, b); err != nil {
		return
	}

	_ = delivery.Ack(false)
}

func (h *BcRefundHandle) Handle(ctx context.Context, msg *msg.BusinessCircleRefund) error {
	member, err := h.MemberCli.GetMemberByWxOpenID(ctx, msg.ShopNumber)
	if err != nil {
		return nil
	}

	q := h.Client.BonusItem.Query()
	q.Where(bonusitem.BcmbTransCodeEQ(msg.TransactionID))
	payItem, err := q.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}

	cr := h.Client.BonusItem.Create()

	cr.SetBcmbNotifyID(msg.MsgID)
	cr.SetBcmbNotifyTime(msg.CreateTime)
	cr.SetBcmbTransCode(msg.RefundID)
	cr.SetBcmbTransTime(msg.RefundTime)
	cr.SetBcmbTransPayCode(msg.TransactionID)
	cr.SetBcmbTransType(1)

	cr.SetMchID(msg.MchID)
	cr.SetShopCode(msg.ShopNumber)

	cr.SetMemberCode(member.Code)
	cr.SetWxOpenID(member.WxOpenID)

	cr.SetAmount(msg.RefundAmount)
	cr.SetBonus(-payItem.Bonus) //扣分
	cr.SetBonusRate(100)

	ms := time.Now().UnixMilli()
	cr.SetCreateTs(ms)

	return cr.Exec(ctx)
}
