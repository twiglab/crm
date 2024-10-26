package x

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

func ID() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}

func Code36() string {
	return uuid.Must(uuid.NewV7()).String()
}

// CheckMobile 检验手机号
func CheckMobile(phone string) error {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	if !reg.MatchString(phone) {
		return errors.New("dadas")
	}

	return nil
}

func AddDay(i int) func() time.Time {
	return func() time.Time {
		now := time.Now()
		return now.AddDate(0, 0, i)
	}
}
