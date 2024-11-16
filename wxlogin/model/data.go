package model

type AuthInfo struct {
	JsCode   string `msg:"js_code"`
	WxOpenID string `msg:"wx_open_id"` // wx 的openID
	AuthCode string `msg:"auth_code"`  // 认证码 6 个字符

	Code    string `msg:"code"` // 用户code
	PhoneNo string `msg:"phone_no"`

	UID string `msg:"id"` // 唯一标识

	Ts int64 `msg:"ts"` //时间戳

}
