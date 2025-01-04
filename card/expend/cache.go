package expend

import (
	"sync"
)

var balanceCache sync.Map

func getCacheBalance(cardCode string) (int64, error) {
	if v, ok := balanceCache.Load(cardCode); ok {
		return v.(int64), nil
	}
	return 0, ErrCardCodeNotInMap
}

func setCacheBalance(cardCode string, balance int64) error {
	if balance < 0 {
		return ErrBalanceLessThanZero
	}
	balanceCache.Store(cardCode, balance)
	return nil
}
