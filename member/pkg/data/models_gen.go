package data

type WxCreateMemberReq struct {
	Code      string `json:"code"`
	WxUnionID string `json:"wxUnionID"`
	WxOpenID  string `json:"wxOpenID"`
}

type WxLoginResp struct {
	Code     string `json:"code"`
	WxOpenID string `json:"wxOpenID"`
	Found    bool   `json:"found"`
}

type MemberResp struct {
	Code     string `json:"code"`
	WxOpenID string `json:"wxOpenID"`
}
type OpenIDReq struct {
	WxOpenID string `json:"wxOpenID"`
}
