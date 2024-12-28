package bc

import (
	"time"
)

type BcVoid struct {
	Result int `json:"result"`
}

type BusinessCirclePointsSync struct {
	TransactionID string `json:"transactionID"`
	//AppID            string    `json:"appID"`
	OpenID           string    `json:"openID"`
	EarnPoints       bool      `json:"earnPoints"`
	IncreasedPoints  int       `json:"increasedPoints"`
	PointsUpdateTime time.Time `json:"pointsUpdateTime"`
}
