package mb

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/member/pkg/data"
	"github.com/twiglab/crm/member/pkg/low"
)

type Member struct {
	Code     string
	WxOpenID string
	Status   int32
	Level    int32
}

type MemberCli struct {
	Client graphql.Client
}

func (c *MemberCli) GetMemberByWxOpenID(ctx context.Context, openID string) (*Member, error) {
	resp, err := low.QryMemberByWxOpenID(ctx, c.Client, data.OpenIDReq{WxOpenID: openID})
	if err != nil {
		return nil, err
	}
	m := resp.GetQryMemberByWxOpenID()
	return &Member{Code: m.Code, WxOpenID: m.WxOpenID, Status: m.Status, Level: m.Level}, nil
}
