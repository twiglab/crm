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
//	@param amount
//	@return error
func changeCardExpend(ctx context.Context, client *ent.Client, code, paycode string, amount int64) error {

	obj, err := client.ChargeRecord.Query().Where(chargerecord.CodeEQ(code)).First(ctx)
	if err != nil {
		return err
	}
	if obj.PayCode != paycode {
		return ErrPayCodeFormat
	}

	// 当前时间戳
	now := time.Now().Unix()

	// 更新信息
	_, err = client.ChargeRecord.Update().
		Where(chargerecord.CodeEQ(code)).
		SetPayTs(now).
		SetDeduct(amount).
		Save(ctx)

	return nil
}
