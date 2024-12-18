package sns

type AuthUserResp struct {
	OpenID  string `json:"open_id"`
	UnionID string `json:"union_id"`
}

type JsCodeReq struct {
	JsCode string `json:"js_code"`
}

type WeChatCommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e WeChatCommonError) Error() string {
	return e.ErrMsg
}
