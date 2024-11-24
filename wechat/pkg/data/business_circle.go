package data

import (
	"fmt"
	"github.com/bytedance/sonic"
)

type BusinessCircleNotifyBase struct {
	Id           string `json:"id"`            // 通知的唯一ID
	CreateTime   string `json:"create_time"`   // 通知时间(遵循rfc3339标准格式)
	EventType    string `json:"event_type"`    // 通知的类型， 会员积分服务授权通知的类型为：MALL_AUTH.ACTIVATE_CARD
	ResourceType string `json:"resource_type"` // 通知的资源数据类型
	Summary      string `json:"summary"`       // 回调摘要
}

// 微信通知数据结构
type OriginBusinessCircleNotifySource struct {
	Algorithm      string `json:"algorithm"`       // 对支付结果数据进行加密的加密算法
	Ciphertext     string `json:"ciphertext"`      // Base64编码后的支付结果数据密文
	OriginalType   string `json:"original_type"`   // 原始回调类型为：discount_card
	AssociatedData string `json:"associated_Data"` // 附加数据
	Nonce          string `json:"nonce"`           // 加密使用的随机串
}

type OriginBusinessCircleAuthor struct {
	BusinessCircleNotifyBase
	Source OriginBusinessCircleNotifySource `json:"source"`
}

/*
商圈会员积分服务授权结果
*/

type BusinessCircleAuthorType string

var (
	// 会员开卡(进卡包) + 未授权会员积分服务
	BusinessCircleAuthorType_Reg BusinessCircleAuthorType = "REGISTERED_MODE"
	// 会员开卡(进卡包）+ 授权会员积分服务
	BusinessCircleAuthorType_RegAndAuth BusinessCircleAuthorType = "REGISTERED_AND_AUTHORIZATION_MODE"
)

type BusinessCircleAuthorSource struct {
	Openid   string                   `json:"openid"`    // 顾客授权时使用的小程序上的OpenID
	Code     string                   `json:"code"`      // 用户开会员卡时的商圈商户号
	Mchid    string                   `json:"mchid"`     // 用户开会员卡时的商圈商户号
	AuthType BusinessCircleAuthorType `json:"auth_type"` // 用户授权类型
}

type BusinessCircleAuthor struct {
	BusinessCircleNotifyBase
	Source BusinessCircleAuthorSource `json:"source"`
}

func EncodeBusinessCircleAuthor(author *BusinessCircleAuthor) (string, error) {
	return sonic.MarshalString(author)
}

func DecodeBusinessCircleAuthor(json string) (*BusinessCircleAuthor, error) {
	var author BusinessCircleAuthor
	if err := sonic.UnmarshalString(json, &author); err != nil {
		return nil, fmt.Errorf("UnmarshalString err:%w", err)
	}
	return &author, nil
}

type BusinessCircleShopBase struct {
	Mchid        string `json:"mchid"`         // 微信支付分配的商户号
	MerchantName string `json:"merchant_name"` // 商圈商户名称
	ShopName     string `json:"shop_name"`     // 门店名称，商圈在商圈小程序上圈店时填写的门店名称
	ShopNumber   string `json:"shop_number"`   // 门店编号，商圈在商圈小程序上圈店时填写的门店编号，用于跟商圈自身已有的商户识别码对齐
}

/*
商圈会员场内支付结果
*/
type BusinessCirclePaymentSource struct {
	BusinessCircleShopBase
	Appid         string `json:"appid"`          // 顾客授权积分时使用的小程序的AppID
	Openid        string `json:"openid"`         // 顾客授权时使用的小程序上的OpenID
	TimeEnd       string `json:"time_end"`       // 交易完成时间，遵循[rfc3339]标准格式
	Amount        int    `json:"amount"`         // 用户实际消费金额，单位（分）
	TransactionId string `json:"transaction_id"` // 微信支付订单号
	CommitTag     string `json:"commit_tag"`     // 手动提交积分标记，自动提交时无该字段，用于区分用户手动申请后推送的积分数据
}

type BusinessCirclePayment struct {
	BusinessCircleNotifyBase
	Source BusinessCirclePaymentSource `json:"source"`
}

func EncodeBusinessCirclePayment(payment *BusinessCirclePayment) (string, error) {
	return sonic.MarshalString(payment)
}

func DecodeBusinessCirclePayment(json string) (*BusinessCirclePayment, error) {
	var author BusinessCirclePayment
	if err := sonic.UnmarshalString(json, &author); err != nil {
		return nil, fmt.Errorf("UnmarshalString err:%w", err)
	}
	return &author, nil
}

/*
商圈会员场内退款结果
*/
type BusinessCircleRefundSource struct {
	BusinessCircleShopBase
	Appid         string `json:"appid"`          // 顾客授权积分时使用的小程序的AppID
	Openid        string `json:"openid"`         // 顾客授权时使用的小程序上的OpenID
	RefundTime    string `json:"time_end"`       // 交易完成时间，遵循[rfc3339]标准格式
	PayAmount     int    `json:"pay_amount"`     // 用户实际消费金额，单位（分）
	RefundAmount  int    `json:"refund_amount"`  // 用户退款金额，单位（分）
	TransactionId string `json:"transaction_id"` // 微信支付订单号
	RefundId      string `json:"refund_id"`      // 微信支付退款单号
}

type BusinessCircleRefund struct {
	BusinessCircleNotifyBase
	Source BusinessCircleRefundSource `json:"source"`
}

func EncodeBusinessCircleRefund(refund *BusinessCircleRefund) (string, error) {
	return sonic.MarshalString(refund)
}

func DecodeBusinessCircleRefund(json string) (*BusinessCircleRefund, error) {
	var author BusinessCircleRefund
	if err := sonic.UnmarshalString(json, &author); err != nil {
		return nil, fmt.Errorf("UnmarshalString err:%w", err)
	}
	return &author, nil
}
