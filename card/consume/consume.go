package consume

import (
	"context"
	"fmt"

	"github.com/twiglab/crm/card/orm/ent"
)

type IConsume interface {
	IOPT
	IDBX
}

type IOPT interface {
	GetBalance(ctx context.Context, cardCode string) (int64, error)
	GetPayCode(ctx context.Context, cardCode string) (string, error)
	CardExpend(ctx context.Context, payCode string, consume int64) error
}

type IDBX interface {
	GetCardDetailByCode(ctx context.Context, cardCode string) (*ent.Card, error)
	GetCardListPagin(ctx context.Context, cursor, order *string, limit *int) ([]*ent.Card, error)
	GetAllCardByMember(ctx context.Context, member string) ([]*ent.Card, error)
	GetChangeRecordDetailByCode(ctx context.Context, code string) (*ent.ChargeRecord, error)
	GetNotConsumeRecords(ctx context.Context, cardCode string) ([]*ent.ChargeRecord, error)

	CardBindMember(ctx context.Context, cardCode, member string) (*ent.Card, error)
	CardActive(ctx context.Context, cardCode string) (*ent.Card, error)

	AddConsumeRecord(ctx context.Context, cardCode string) (string, error)
	ChangeCardExpend(ctx context.Context, cardCode, code string, consume int64) error
}

type Consume struct {
	Client *ent.Client
}

// GetBalance 卡余额查询
//
//	@param ctx
//	@param cardCode
//	@return int64
//	@return error
func (c *Consume) GetBalance(ctx context.Context, cardCode string) (int64, error) {
	cobj, err := c.GetCardDetailByCode(ctx, cardCode)
	if err != nil {
		return 0, err
	}

	crl, err := c.GetNotConsumeRecords(ctx, cardCode)
	if err != nil {
		return 0, err
	}

	balance := cobj.LastCleanBalance
	for _, cr := range crl {
		balance -= cr.Deduct
	}

	return balance, nil
}

// GetPayCode 获得支付码
//
//	@param ctx
//	@param cardCode
//	@return string
//	@return error
func (c *Consume) GetPayCode(ctx context.Context, cardCode string) (string, error) {
	balance, err := c.GetBalance(ctx, cardCode)
	if err != nil {
		return "", err
	}

	d, err := c.AddConsumeRecord(ctx, cardCode, balance)
	if err != nil {
		return "", err
	}

	return d.PayCode, nil
}

// CardExpend 消费
//
//	@param ctx
//	@param payCode
//	@param consume
//	@return error
func (c *Consume) CardExpend(ctx context.Context, payCode string, consume int64) error {
	var balance int64
	_, code, balance, err := splitPayCode(payCode)
	if err != nil {
		return err
	}

	d, err := c.GetChangeRecordDetailByCode(ctx, code)
	if err != nil {
		return err
	}

	nb, err := c.GetBalance(ctx, d.CardCode)
	if err != nil {
		return err
	}

	if balance != nb {
		// TODO add log
		balance = nb
	}

	if balance < consume {
		return fmt.Errorf("balance not enough")
	}

	return c.ChangeCardExpend(ctx, d.CardCode, code, consume)
}
