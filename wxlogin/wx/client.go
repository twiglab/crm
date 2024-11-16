package wx

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/member/pkg/data"
	"github.com/twiglab/crm/member/pkg/low"
)

type MemberCli struct {
	client graphql.Client
}

func NewMemberCli(client graphql.Client) *MemberCli {
	return &MemberCli{client: client}
}

func (c *MemberCli) QueryWxMember(ctx context.Context, wxOpenID string) error {
	_, err := low.QueryWxMember(ctx, c.client, data.OpenIDReq{WxOpenID: wxOpenID})
	if err != nil {
		return err
	}
	return nil
}
