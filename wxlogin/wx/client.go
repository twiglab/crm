package wx

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/member/pkg/data"
	"github.com/twiglab/crm/member/pkg/low"
)

type Member struct {
	Code     string `json:"code"`
	WxOpenID string `json:"wxOpenID"`
}

type MemberCli struct {
	client graphql.Client
}

func NewMemberCli(client graphql.Client) *MemberCli {
	return &MemberCli{client: client}
}

/*
func (c *MemberCli) QueryWxMember(ctx context.Context, wxOpenID string) (*Member, error) {
	req, err := low.QueryWxMember(ctx, c.client, data.OpenIDReq{WxOpenID: wxOpenID})
	if err != nil {
		return nil, err
	}
	m := req.GetQueryWxMember()
	return &Member{Code: m.Code, WxOpenID: m.WxOpenID}, nil
}

func (c *MemberCli) Cr(ctx context.Context) {
	low.CreateWxMember(ctx, c.client, data.CreateWxMemberReq{})
}
*/

func (c *MemberCli) LoginOrCr(ctx context.Context, wxOpenID string) (*Member, error) {
	req, err := low.QueryWxMember(ctx, c.client, data.OpenIDReq{WxOpenID: wxOpenID})
	if err != nil {
		return nil, err

	}
	found := true
	if found {
		m := req.GetQueryWxMember()
		return &Member{Code: m.Code, WxOpenID: m.WxOpenID}, nil
	}
	r, err := low.CreateWxMember(ctx, c.client, data.CreateWxMemberReq{Code: "xx", WxOpenID: wxOpenID})
	if err != nil {
		return nil, err

	}
	m := r.GetCreateWxMember()
	return &Member{Code: m.Code, WxOpenID: m.WxOpenID}, nil
}
