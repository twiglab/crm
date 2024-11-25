package bc

const (
	REGISTERED_MODE                   = "REGISTERED_MODE"
	REGISTERED_AND_AUTHORIZATION_MODE = "REGISTERED_AND_AUTHORIZATION_MODE"
)

// 商圈会员积分服务授权结果通知
type BusinessCircleAuthorSource struct {
	OpenID   string `json:"openid"`    // 顾客授权时使用的小程序上的OpenID
	Code     string `json:"code"`      // 用户开会员卡时的商圈商户号
	MchID    string `json:"mchid"`     // 用户开会员卡时的商圈商户号
	AuthType string `json:"auth_type"` // 用户授权类型
}

type BusinessCircleShopBase struct {
	MchID        string `json:"mchid"`         // 微信支付分配的商户号
	MerchantName string `json:"merchant_name"` // 商圈商户名称
	ShopName     string `json:"shop_name"`     // 门店名称，商圈在商圈小程序上圈店时填写的门店名称
	ShopNumber   string `json:"shop_number"`   // 门店编号，商圈在商圈小程序上圈店时填写的门店编号，用于跟商圈自身已有的商户识别码对齐
}

// 商圈会员场内支付结果通知
type BusinessCirclePaymentSource struct {
	BusinessCircleShopBase
	AppID         string `json:"appid"`          // 顾客授权积分时使用的小程序的AppID
	OpenID        string `json:"openid"`         // 顾客授权时使用的小程序上的OpenID
	Amount        int    `json:"amount"`         // 用户实际消费金额，单位（分）
	TimeEnd       string `json:"time_end"`       // 交易完成时间，遵循[rfc3339]标准格式
	TransactionID string `json:"transaction_id"` // 微信支付订单号
}

// 商圈会员场内退款结果通知
type BusinessCircleRefundSource struct {
	BusinessCircleShopBase
	AppID         string `json:"appid"`          // 顾客授权积分时使用的小程序的AppID
	OpenID        string `json:"openid"`         // 顾客授权时使用的小程序上的OpenID
	RefundTime    string `json:"time_end"`       // 交易完成时间，遵循[rfc3339]标准格式
	PayAmount     int    `json:"pay_amount"`     // 用户实际消费金额，单位（分）
	RefundAmount  int    `json:"refund_amount"`  // 用户退款金额，单位（分）
	TransactionID string `json:"transaction_id"` // 微信支付订单号
	RefundID      string `json:"refund_id"`      // 微信支付退款单号
}
