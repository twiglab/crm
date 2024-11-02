package werr

type WeChatCommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e WeChatCommonError) Error() string {
	return ""
}
