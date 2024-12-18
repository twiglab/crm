package sns

type AuthUserResp struct {
	OpenID  string
	UnionID string
}

type JsCodeReq struct {
	JsCode string
}

type WeChatCommonError struct {
	ErrCode int64
	ErrMsg  string
}

func (e WeChatCommonError) Error() string {
	return e.ErrMsg
}
