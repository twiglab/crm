package mini

import (
	"context"

	"github.com/twiglab/crm/wechat/pkg/data"
	"github.com/twiglab/crm/wechat/pkg/werr"
)

func Code2Session(ctx context.Context, jsCode string) (resp *data.AuthUserResp, e error) {
	result, err := GetAuth().Code2SessionContext(ctx, jsCode)
	resp = &data.AuthUserResp{
		OpenID: result.OpenID,
		// SessionKey: result.SessionKey,
		UnionID: result.UnionID,
	}

	if err != nil {
		e = &werr.WeChatCommonError{
			ErrCode: result.ErrCode,
			ErrMsg:  result.Error(),
		}
	}
	return
}
