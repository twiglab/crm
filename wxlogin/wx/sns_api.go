package wx

import (
	"github.com/imroc/req/v3"
)

const (
	baseURL = "https://api.weixin.qq.com"
	sns     = "/sns/jscode2session"
)

type WeChatClient struct {
	wxappid  string
	wxsecret string
	client   *req.Client
}

func NewWeChatClient(wxappid, wxsecret string) *WeChatClient {
	c := req.C().EnableDumpAll().
		SetBaseURL(baseURL).
		SetCommonQueryParam("appid", wxappid).
		SetCommonQueryParam("secret", wxsecret)

	return &WeChatClient{
		wxappid:  wxappid,
		wxsecret: wxsecret,
		client:   c,
	}
}
func (c *WeChatClient) WxCode2Session(code string) (res Code2SessionResp, err error) {
	_, err = c.client.R().
		SetQueryParam("js_code", code).
		SetQueryParam("grant_type", "authorization_code").
		SetSuccessResult(&res).
		Get(sns)
	return
}

type WxCode2SessionBaseInfo struct {
	OpenID     string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	UnionID    string `json:"unionid,omitempty"`
}

type Code2SessionResp struct {
	WxCode2SessionBaseInfo
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
