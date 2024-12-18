package sns

type AuthUserResp struct {
	OpenID  string `json:"openID"`
	UnionID string `json:"unionID"`
}

type JsCodeReq struct {
	JsCode string `json:"jsCode"`
}

type WeChatCommonError struct {
	ErrCode int64
	ErrMsg  string
}

func (e WeChatCommonError) Error() string {
	return e.ErrMsg
}
