package data

type AuthUserResp struct {
	OpenID  string `json:"openID"`
	UnionID string `json:"unionID"`
}

type JsCodeReq struct {
	JsCode string `json:"jsCode"`
}
