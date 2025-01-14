package data

type CodeReq struct {
	Code string `json:"code"`
}

type WxCreateMemberReq struct {
	Code      string `json:"code"`
	WxUnionID string `json:"wxUnionID"`
	WxOpenID  string `json:"wxOpenID"`
}

type WxLoginResp struct {
	Code     string `json:"code"`
	WxOpenID string `json:"wxOpenID"`
	Level    int32  `json:"level"`
	Status   int32  `json:"status"`

	Found bool `json:"found"`
}

type MemberResp struct {
	Code     string `json:"code"`
	WxOpenID string `json:"wxOpenID"`
	Level    int32  `json:"level"`
	Status   int32  `json:"status"`
}
type OpenIDReq struct {
	WxOpenID string `json:"wxOpenID"`
}
