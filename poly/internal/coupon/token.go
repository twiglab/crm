package coupon

import (
	"math/rand"
	"time"
)

var tempGlobalToken map[string]int = make(map[string]int)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// token 生成
func GenerateToken() (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 4)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	t := string(b)
	tempGlobalToken[t] = 1
	return t, nil
}

// token 验证
func CheckToken(token string) bool {
	_, ok := tempGlobalToken[token]
	return ok
}

// token 删除
func DeleteToken(token string) {
	delete(tempGlobalToken, token)
}

// token 预占用
func PreUseToken(token string) {
	tempGlobalToken[token] = 1
}
