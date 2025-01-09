package consume

import (
	"context"
	"fmt"
	"time"

	"github.com/twiglab/crm/card/orm/ent"
	"github.com/twiglab/crm/card/orm/ent/card"
	"github.com/twiglab/crm/card/orm/ent/chargerecord"
)

// GetCardDetailByCode 根据卡号获取卡详情
//
//	@param ctx
//	@param cardCode
//	@return *ent.Card
//	@return error
func (c *Consume) GetCardDetailByCode(ctx context.Context, cardCode string) (*ent.Card, error) {
	return c.Client.Card.Query().Where(card.CodeEQ(cardCode)).First(ctx)
}

// GetChangeRecordDetailByCode 根据code获取消费记录详情
//
//	@param ctx
//	@param client
//	@param code
//	@return *ent.ChargeRecord
//	@return error
func (c *Consume) GetChangeRecordDetailByCode(ctx context.Context, code string) (*ent.ChargeRecord, error) {
	return c.Client.ChargeRecord.Query().Where(chargerecord.CodeEQ(code)).First(ctx)
}

// AddConsumeRecord 数据库添加支付请求
//
//	@param ctx
//	@param cardCode
//	@param balance
//	@return *ent.ChargeRecord
//	@return error
func (c *Consume) AddConsumeRecord(ctx context.Context, cardCode string, balance int64) (*ent.ChargeRecord, error) {
	tx, err := c.Client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	cobj, err := tx.ChargeRecord.Create().SetCardCode(cardCode).Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	payCode := genCode(cobj.Code, balance)

	if _, err := tx.ChargeRecord.UpdateOne(cobj).SetPayCode(payCode).Save(ctx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	cobj.PayCode = payCode
	return cobj, nil
}

// GetRecordByPaycode 根据paycode 查找记录
//
//	@param ctx
//	@param paycode
//	@return *ent.ChargeRecord
//	@return error
func (c *Consume) GetRecordByPaycode(ctx context.Context, paycode string) (*ent.ChargeRecord, error) {
	return c.Client.ChargeRecord.Query().Where(chargerecord.PayCodeEQ(paycode)).First(ctx)
}

// GetNotConsumeRecords 查找所有未清算消费记录
//
//	@param ctx
//	@param cardCode
//	@return []*ent.ChargeRecord
//	@return error
func (c *Consume) GetNotConsumeRecords(ctx context.Context, cardCode string) ([]*ent.ChargeRecord, error) {
	return c.Client.ChargeRecord.Query().Where(chargerecord.CardCodeEQ(cardCode), chargerecord.StatusEQ(0)).All(ctx)
}

// ChangeCardExpend 添加卡的消费记录
//
//	@param ctx
//	@param code
//	@param consume
//	@return error
func (c *Consume) ChangeCardExpend(ctx context.Context, cardCode, code string, consume int64) error {
	// 当前时间戳
	now := time.Now().Unix()

	// 开启事务
	tx, err := c.Client.Tx(ctx)
	if err != nil {
		return err
	}

	// 更新信息
	_, err = tx.ChargeRecord.Update().
		Where(chargerecord.CodeEQ(code)).
		SetPayTs(now).
		SetDeduct(consume).
		SetStatus(1).
		Save(ctx)

	if err != nil {
		tx.Rollback()
	}

	_, err = tx.Card.Update().Where(card.CodeEQ(cardCode)).SetLastUseTs(now).Save(ctx)
	if err != nil {
		tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// CardBindMember 绑卡
//
//	@param ctx
//	@param cardCode
//	@param member
//	@return error
func (c *Consume) CardBindMember(ctx context.Context, cardCode, member string) (*ent.Card, error) {
	cobj, err := c.GetCardDetailByCode(ctx, cardCode)
	if err != nil {
		return nil, err
	}
	if cobj.MemberCode != "" {
		return nil, fmt.Errorf("card already binded")
	}

	_, err = c.Client.Card.UpdateOne(cobj).SetMemberCode(member).SetBindTime(time.Now()).Save(ctx)
	return cobj, err
}

// CardActive 激活卡
//
//	@param ctx
//	@param cardCode
//	@return error
func (c *Consume) CardActive(ctx context.Context, cardCode string) (*ent.Card, error) {
	cobj, err := c.GetCardDetailByCode(ctx, cardCode)
	if err != nil {
		return nil, err
	}

	_, err = c.Client.Card.UpdateOne(cobj).SetStatus(1).Save(ctx)
	return cobj, err
}

// GetCardListPagin 分页查询
//
//	@param ctx
//	@param cursor
//	@param order
//	@param limit
//	@return []*ent.Card
//	@return error
func (c *Consume) GetCardListPagin(ctx context.Context, cursor, order *string, limit *int) ([]*ent.Card, error) {
	q := c.Client.Card.Query()
	if cursor != nil {
		if order != nil && *order == "asc" {
			q.Where(card.CodeGT(*cursor))
		} else {
			q.Where(card.CodeLT(*cursor))
		}
	}

	if limit != nil {
		q.Limit(*limit)
	} else {
		q.Limit(10)
	}

	return q.All(ctx)
}

// GetAllCardByMember  用户所有卡
//
//	@param ctx
//	@param member
//	@return []*ent.Card
//	@return error
func (c *Consume) GetAllCardByMember(ctx context.Context, member string) ([]*ent.Card, error) {
	return c.Client.Card.Query().Where(card.MemberCodeEQ(member)).All(ctx)
}
