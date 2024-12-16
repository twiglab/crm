package sns

import (
	"context"

	"github.com/it512/box"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"github.com/twiglab/crm/wechat/pkg/sns"
)

var AuthKey = box.RndKey()

type Auth struct {
	wxAuth *auth.Auth
}

func NewAuth(a *auth.Auth) *Auth {
	return &Auth{wxAuth: a}
}

func (a *Auth) Code2Session(ctx context.Context, jsCode string) (resp *sns.AuthUserResp, e error) {
	result, err := a.wxAuth.Code2SessionContext(ctx, jsCode)
	resp = &sns.AuthUserResp{
		OpenID:  result.OpenID,
		UnionID: result.UnionID,
	}

	if err != nil {
		e = &sns.WeChatCommonError{
			ErrCode: result.ErrCode,
			ErrMsg:  result.Error(),
		}
	}
	return
}
