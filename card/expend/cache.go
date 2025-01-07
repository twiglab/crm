package expend

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/twiglab/crm/card/orm/ent"
	"github.com/twiglab/crm/card/orm/ent/card"
)

type IExpendCache interface {
	GetPayCode(ctx context.Context, cardCode string) (string, error)
	CardExpend(ctx context.Context, payCode string, consume int64) error
	ClearCache() error
	GetAllBalance() string

	// getBalance(ctx context.Context, code string) (int64, error)
	getClient() *ent.Client
	getCache() *sync.Map
}

type expendVar struct {
	client *ent.Client
	cache  sync.Map
}

// GetPayCode 生成付款码
//
//	@param ctx
//	@param cardCode
//	@return string
//	@return error
func (v *expendVar) GetPayCode(ctx context.Context, cardCode string) (string, error) {
	b, err := getBalance(ctx, v, cardCode)
	if err != nil {
		return "", err
	}

	pc, err := genPayCode(ctx, v.client, b, cardCode)
	if err != nil {
		return "", err
	}

	return pc, nil
}

// CardExpend 消费
//
//	@param ctx
//	@param payCode
//	@param consume
//	@return error
func (v *expendVar) CardExpend(ctx context.Context, payCode string, consume int64) error {
	// 查找
	cr, err := getRecordByPaycode(ctx, v.client, payCode)
	if err != nil {
		return err
	}
	_, code, _, err := splitPayCode(payCode)
	if err != nil {
		return err
	}

	ob, ok := v.cache.Load(code)
	if !ok {
		// TODO log记录异常
		if cr.Status >= 1 {
			return nil
		}
	}

	if ob.(int64) < consume {
		return ErrCardBalanceNotEnough
	}

	changeCardExpend(ctx, v.client, code, payCode, consume)

	v.cache.Store(code, ob.(int64)-consume)

	return nil
}
func (v *expendVar) getClient() *ent.Client {
	return v.client
}
func (v *expendVar) getCache() *sync.Map {
	return &v.cache
}

// ClearCache 清空缓存
//
//	@return error
func (v *expendVar) ClearCache() error {
	v.cache = sync.Map{}
	return nil
}

// GetAllBalance 返回所有缓存
//
//	@return string
func (v *expendVar) GetAllBalance() string {
	m := make(map[string]int64)
	v.cache.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.(int64)
		return true
	})
	jStr, _ := json.Marshal(m)
	return string(jStr)
}

func getBalance(ctx context.Context, v *expendVar, code string) (int64, error) {
	if v, ok := v.cache.Load(code); ok {
		return v.(int64), nil
	}

	cobj, err := v.client.Card.Query().Where(card.CodeEQ(code)).First(ctx)
	if err != nil {
		return 0, err
	}

	v.cache.Store(code, cobj.Amount)

	return cobj.Amount, nil
}
