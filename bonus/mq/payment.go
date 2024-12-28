package mq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/twiglab/crm/bonus/bc"
	"github.com/twiglab/crm/bonus/mb"
	"github.com/twiglab/crm/bonus/orm/ent"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
)

type BcPaymentHandle struct {
	Client    *ent.Client
	MemberCli *mb.MemberCli
	BcClient  *bc.WxBcCli
}

func (h *BcPaymentHandle) RecieveDelivery(ctx context.Context, delivery amqp.Delivery) {
	b := &msg.BusinessCirclePayment{}
	if _, err := b.UnmarshalMsg(delivery.Body); err != nil {
		return
	}

	if err := h.Handle(ctx, b); err != nil {
		return
	}

	_ = delivery.Ack(false)
}

func (h *BcPaymentHandle) Handle(ctx context.Context, msg *msg.BusinessCirclePayment) error {
	member, err := h.MemberCli.GetMemberByWxOpenID(ctx, msg.ShopNumber)
	if err != nil {
		return nil
	}

	cr := h.Client.BonusItem.Create()

	cr.SetBcmbNotifyID(msg.MsgID)
	cr.SetBcmbNotifyTime(msg.CreateTime)
	cr.SetBcmbTransCode(msg.TransactionID)
	cr.SetBcmbTransPayCode(msg.TransactionID)
	cr.SetBcmbTransTime(msg.TimeEnd)
	cr.SetBcmbTransType(0)

	cr.SetMchID(msg.MchID)
	cr.SetShopCode(msg.ShopNumber)

	cr.SetMemberCode(member.Code)
	cr.SetWxOpenID(member.WxOpenID)

	cr.SetAmount(msg.Amount)
	cr.SetBonus(msg.Amount)
	cr.SetBonusRate(100)

	ms := time.Now().UnixMilli()
	cr.SetCreateTs(ms)

	err = cr.Exec(ctx)

	if err != nil {
		return err
	}

	_ = h.BcClient.PointsSync(ctx, bc.PointsSync{
		TransactionID:   msg.TransactionID,
		OpenID:          msg.OpenID,
		IncreasedPoints: msg.Amount,
	})

	return nil
}