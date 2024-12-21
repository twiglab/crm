package msg

import "time"

//go:generate msgp

const (
	REGISTERED_MODE                   = "REGISTERED_MODE"
	REGISTERED_AND_AUTHORIZATION_MODE = "REGISTERED_AND_AUTHORIZATION_MODE"
)

type BusinessCircleMsg struct {
	MsgID      string    `json:"msg_id,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	MsgType    string    `json:"msg_type,omitempty"`
	Summary    string    `json:"summary,omitempty"`
}

// 商圈会员积分服务授权结果通知
type BusinessCircleAuthor struct {
	BusinessCircleMsg
	OpenID   string `json:"open_id,omitempty"`   // 顾客授权时使用的小程序上的OpenID
	Code     string `json:"code,omitempty"`      // 用户开会员卡时的商圈商户号
	MchID    string `json:"mch_id,omitempty"`    // 用户开会员卡时的商圈商户号
	AuthType string `json:"auth_type,omitempty"` // 用户授权类型
}

type BusinessCircleShopBase struct {
	MchID        string // 微信支付分配的商户号
	MerchantName string // 商圈商户名称
	ShopName     string // 门店名称，商圈在商圈小程序上圈店时填写的门店名称
	ShopNumber   string // 门店编号，商圈在商圈小程序上圈店时填写的门店编号，用于跟商圈自身已有的商户识别码对齐
}

// 商圈会员场内支付结果通知
type BusinessCirclePayment struct {
	BusinessCircleMsg
	BusinessCircleShopBase
	AppID         string    // 顾客授权积分时使用的小程序的AppID
	OpenID        string    // 顾客授权时使用的小程序上的OpenID
	Amount        int       // 用户实际消费金额，单位（分）
	TimeEnd       time.Time // 交易完成时间，遵循[rfc3339]标准格式
	TransactionID string    // 微信支付订单号
}

// 商圈会员场内退款结果通知
type BusinessCircleRefund struct {
	BusinessCircleMsg
	BusinessCircleShopBase
	AppID         string    // 顾客授权积分时使用的小程序的AppID
	OpenID        string    // 顾客授权时使用的小程序上的OpenID
	RefundTime    time.Time // 交易完成时间，遵循[rfc3339]标准格式
	PayAmount     int       // 用户实际消费金额，单位（分）
	RefundAmount  int       // 用户退款金额，单位（分）
	TransactionID string    // 微信支付订单号
	RefundID      string    // 微信支付退款单号
}
