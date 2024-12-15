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
	RefundTime    string `json:"refund_time"`    // 交易完成时间，遵循[rfc3339]标准格式
	PayAmount     int    `json:"pay_amount"`     // 用户实际消费金额，单位（分）
	RefundAmount  int    `json:"refund_amount"`  // 用户退款金额，单位（分）
	TransactionID string `json:"transaction_id"` // 微信支付订单号
	RefundID      string `json:"refund_id"`      // 微信支付退款单号
}

// 商圈会员积分同步
type BusinessCirclePointsSync struct {
	TransactionID    string `json:"transaction_id"`     // 微信支付订单号
	AppID            string `json:"appid"`              // 顾客授权积分时使用的小程序的AppID
	OpenID           string `json:"openid"`             // 顾客授权时使用的小程序上的OpenID
	EarnPoints       bool   `json:"earn_points"`        // 用于标明此单是否获得积分，true为获得积分，false为未获得
	IncreasedPoints  int    `json:"increased_points"`   // 顾客此笔交易新增的积分值
	PointsUpdateTime int64  `json:"points_update_time"` // 为顾客此笔交易成功积分的时间
}

// 商圈会员积分服务授权查询
type BusinessCircleAuthorQuery struct {
	OpenID string `json:"openid"` // 顾客授权时使用的小程序上的OpenID
	AppID  string `json:"appid"`  // 顾客授权积分时使用的小程序的AppID
}

const (
	AuthorizeStateUnauthorized = "UNAUTHORIZED" // 未授权
	AuthorizeStateAuthorized   = "AUTHORIZED"   // 已授权
	AuthorizeStateDeauthorized = "DEAUTHORIZED" // 已取消授权
)

// 商圈会员待积分状态查询
type BusinessCirclePointsStatusQuery struct {
	OpenID  string `json:"openid"`  // 用户在商户对应AppID下的唯一标识
	AppID   string `json:"appid"`   // 支持服务号、小程序等类型的AppID，需已与brandid完成下单AppID绑定
	BrandID int    `json:"brandid"` // 调用方商户号对应的品牌brandid，调用方商户号需为此品牌brandid的品牌主商户号或品牌服务商商户号
}

const (
	ParkingStateIn  = "IN"  // 用户开车进入商圈
	ParkingStateOut = "OUT" // 用户开车离开商圈
)

// 商圈会员停车状态同步
type BusinessCircleParkingSyncQuery struct {
	OpenID      string `json:"openid"`       // 用户在商户对应AppID下的唯一标识
	AppID       string `json:"appid"`        // 支持服务号、小程序等类型的AppID，需已与brandid完成下单AppID绑定
	BrandID     int    `json:"brandid"`      // 调用方商户号对应的品牌brandid，调用方商户号需为此品牌brandid的品牌主商户号或品牌服务商商户号
	PlateNumber string `json:"plate_number"` // 首位需为省份的中文简称，第二位起支持大写字母、数字、中文
	State       string `json:"state"`        // 停车状态，服务商模式下必传
	Time        int64  `json:"time"`         // 在场状态更新时间，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE
}
