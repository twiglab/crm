package mb

import (
	"context"
	"errors"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/member/pkg/data"
	"github.com/twiglab/crm/member/pkg/low"
)

type Member struct {
	Code     string
	WxOpenID string
}

type MemberCli struct {
	Client graphql.Client
}

func (c *MemberCli) GetMemberByWxOpenID(ctx context.Context, openID string) (*Member, error) {
	resp, err := low.QueryWxMemberByOpenID(ctx, c.Client, data.OpenIDReq{WxOpenID: openID})
	if err != nil {
		return nil, err
	}

	m := resp.GetQueryWxMemberByOpenID()
	if !m.Found {
		return nil, errors.New("not found")
	}

	return &Member{Code: m.Code, WxOpenID: m.WxOpenID}, nil
}
