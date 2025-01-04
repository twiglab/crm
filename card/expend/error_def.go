package expend

import "errors"

var (
	ErrCardCodeNotInMap     = errors.New("缓存没有该卡号")
	ErrCardBalanceNotEnough = errors.New("可用余额不足")
	ErrBalanceLessThanZero  = errors.New("余额不能小于0")
	ErrPayCodeFormat        = errors.New("paycode格式错误")
)
