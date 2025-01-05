package code

import "github.com/google/uuid"

// 必须用这个函数生成code，保证uud有序
func Code36() string {
	u := uuid.Must(uuid.NewV7())
	return u.String()
}

func NewV7() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}
