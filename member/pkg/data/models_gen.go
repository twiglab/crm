package data

type CreateWxMemberReq struct {
	Code     string `json:"code"`
	WxOpenID string `json:"wxOpenID"`
}

type MemberResp struct {
	Code     string `json:"code"`
	WxOpenID string `json:"wxOpenID"`
}

type OpenIDReq struct {
	WxOpenID string `json:"wxOpenID"`
}
