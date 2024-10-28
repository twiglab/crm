package wxservice

import "fmt"

type LoginResp struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
	WeChatCommonError
}

func login(jsCode string) (*LoginResp, error) {
	result, err := wxMiniProg.GetAuth().Code2Session(jsCode)
	if err != nil {
		return nil, fmt.Errorf("code2session:%w", err)
	}
	resp := LoginResp{
		OpenID:     result.OpenID,
		SessionKey: result.SessionKey,
		UnionID:    result.UnionID,
		WeChatCommonError: WeChatCommonError{
			ErrCode: result.ErrCode,
			ErrMsg:  result.ErrMsg,
		},
	}
	return &resp, nil
}
