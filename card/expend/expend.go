package expend

import (
	"context"
	"errors"

	"github.com/twiglab/crm/card/orm/ent"
)

type IExpend interface {
	GetPayCode(ctx context.Context, client *ent.Client, cardCode string) (string, error)
	CardExpend(ctx context.Context, client *ent.Client, payCode string, amount int64) error
}

type Expend struct{}

// 获得paycode
func (e *Expend) GetPayCode(ctx context.Context, client *ent.Client, cardCode string) (string, error) {
	var b int64
	b, err := getCacheBalance(cardCode)
	if err != nil {
		if errors.Is(err, ErrCardCodeNotInMap) {
			// 从数据库中获取
			q, err := getCardDetailByCode(ctx, client, cardCode)
			if err != nil {
				return "", err
			}
			setCacheBalance(cardCode, q.Amount)
			b = q.Amount
		} else {
			return "", nil
		}
	}

	pc, err := genPayCode(ctx, client, b, cardCode)
	if err != nil {
		return "", err
	}

	return pc, nil
}

// CardExpend 卡消费
//
//	@param cardCode
//	@param amount
//	@return error
func (e *Expend) CardExpend(ctx context.Context, client *ent.Client, payCode string, amount int64) error {
	_, code, _, err := splitPayCode(payCode)
	if err != nil {
		return err
	}

	cobj, err := getChangeRecordDetailByCode(ctx, client, code)
	if err != nil {
		return err
	}

	b, err := getCacheBalance(cobj.CardCode)
	if err != nil {
		return err
	}

	if b < amount {
		return ErrCardBalanceNotEnough
	}

	err = changeCardExpend(ctx, client, code, payCode, amount)
	if err != nil {
		return err
	}

	setCacheBalance(cobj.CardCode, b-amount)

	return nil
}
