package expend

import (
	"context"
	"time"

	"github.com/twiglab/crm/card/orm/ent"
	"github.com/twiglab/crm/card/orm/ent/card"
	"github.com/twiglab/crm/card/orm/ent/chargerecord"
)

// getCardDetailByCode 根据卡号获取卡详情
//
//	@param ctx
//	@param client
//	@param cardCode
//	@return *ent.Card
//	@return error
func getCardDetailByCode(ctx context.Context, client *ent.Client, cardCode string) (*ent.Card, error) {
	return client.Card.Query().Where(card.CodeEQ(cardCode)).First(ctx)
}

// getChangeRecordDetailByCode 根据code获取消费记录详情
//
//	@param ctx
//	@param client
//	@param code
//	@return *ent.ChargeRecord
//	@return error
func getChangeRecordDetailByCode(ctx context.Context, client *ent.Client, code string) (*ent.ChargeRecord, error) {
	return client.ChargeRecord.Query().Where(chargerecord.CodeEQ(code)).First(ctx)
}

// addRecord 数据库添加支付请求，获得code
//
//	@param ctx
//	@param client
//	@return string
//	@return error
func addRecord(ctx context.Context, client *ent.Client, cardCode string) (string, error) {
	cobj, err := client.ChargeRecord.Create().SetCardCode(cardCode).Save(ctx)
	if err != nil {
		return "", err
	}
	return cobj.Code, nil
}

// 根据paycode 查找记录
func getRecordByPaycode(ctx context.Context, client *ent.Client, paycode string) (*ent.ChargeRecord, error) {
	return client.ChargeRecord.Query().Where(chargerecord.PayCodeEQ(paycode)).First(ctx)
}

// updateRecordPaycode
//
//	@param ctx
//	@param client
//	@param code
//	@param paycode
//	@return error
func updateRecordPaycode(ctx context.Context, client *ent.Client, code, paycode string) error {
	_, err := client.ChargeRecord.Update().Where(chargerecord.CodeEQ(code)).SetPayCode(paycode).Save(ctx)
	return err
}

// changeCardExpend 添加卡的消费记录
//
//	@param ctx
//	@param client
//	@param code
//	@param paycode
//	@param consume
//	@return error
func changeCardExpend(ctx context.Context, client *ent.Client, code, paycode string, consume int64) error {

	obj, err := client.ChargeRecord.Query().Where(chargerecord.CodeEQ(code)).First(ctx)
	if err != nil {
		return err
	}
	if obj.PayCode != paycode {
		return ErrPayCodeFormat
	}

	// 当前时间戳
	now := time.Now().Unix()

	// 开启事务
	txn, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	// 更新信息
	_, err = txn.ChargeRecord.Update().
		Where(chargerecord.CodeEQ(code)).
		SetPayTs(now).
		SetDeduct(consume).
		SetStatus(1).
		Save(ctx)

	if err != nil {
		txn.Rollback()
	}

	_, err = txn.Card.Update().Where(card.CodeEQ(code)).SetLastUseTs(now).Save(ctx)
	if err != nil {
		txn.Rollback()
	}

	txn.Commit()

	return err
}

// clearCardExpend 卡消费记录清算
//
//	@param ctx
//	@param client
//	@param code
//	@param balance
//	@return error
func clearCardExpend(ctx context.Context, client *ent.Client, code string, balance int64) error {

	// 查找card信息
	cobj, err := client.Card.Query().Where(card.CodeEQ(code)).First(ctx)
	if err != nil {
		return err
	}

	// 开始事务
	txn, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	// 查找所有未清算的记录
	changeList, err := txn.ChargeRecord.Query().Where(chargerecord.CardCodeEQ(code), chargerecord.StatusEQ(0)).All(ctx)
	if err != nil {
		return nil
	}

	// 计算
	ok := calcBalance(balance, cobj.LastCleanBalance, changeList)
	if !ok {
		// TODO log
		return ErrClearingDataNotMatch
	}

	// TODO ent设置不可变，数据变动部分后续研究
	// _, err = txn.Card.Update().Where(card.CodeEQ(code)).last_clean_balance(balance).Save(ctx)

	return nil
}
