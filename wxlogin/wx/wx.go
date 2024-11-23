package wx

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/wechat/pkg/data"
	"github.com/twiglab/crm/wechat/pkg/low"
)

type WxCli struct {
	graphql.Client
}

type Codes struct {
	OpenID  string `json:"openid"`
	Unionid string `json:"unionid"`
}

// AuthUser 微信jscode->(openid,unionid)
//
//	@param ctx
//	@param jsCode
//	@return *Codes
//	@return error
func (c *WxCli) AuthUser(ctx context.Context, jsCode string) (*Codes, error) {
	resp, err := low.AuthUser(ctx, c.Client, data.JsCodeReq{JsCode: jsCode})
	if err != nil {
		return nil, err
	}
	return &Codes{OpenID: resp.AuthUser.OpenID, Unionid: resp.AuthUser.UnionID}, nil
}
