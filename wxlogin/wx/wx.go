package wx

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/wechat/pkg/data"
	"github.com/twiglab/crm/wechat/pkg/low"
)

type WxCli struct {
	client graphql.Client
}

func NewWxCli(client graphql.Client) *WxCli {
	return &WxCli{client: client}
}

func (c *WxCli) AuthUser(ctx context.Context, jsCode string) error {
	_, err := low.AuthUser(ctx, c.client, data.JsCodeReq{JsCode: jsCode})
	if err != nil {
		return err
	}
	return nil
}
