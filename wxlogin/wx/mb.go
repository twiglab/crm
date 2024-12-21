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
	graphql.Client
}

// LoginOrCr 根据微信openid查询用户信息，如果用户不存在则创建新用户
//
//	@param ctx
//	@param wxOpenID
//	@return *Member
//	@return error
func (c *MemberCli) LoginOrCr(ctx context.Context, wxOpenID string) (*Member, error) {
	// 根据微信openid查询用户信息
	req, err := low.QueryWxMemberByOpenID(ctx, c.Client, data.OpenIDReq{WxOpenID: wxOpenID})
	if err != nil {
		return nil, err
	}

	m := req.GetQueryWxMemberByOpenID()
	if m.Found {
		return &Member{Code: m.Code, WxOpenID: m.WxOpenID}, nil
	}

	// 创建用户
	r, err := low.CreateWxMember(ctx, c.Client, data.CreateWxMemberReq{Code: "xx", WxOpenID: wxOpenID})
	if err != nil {
		return nil, err
	}
	qm := r.GetCreateWxMember()
	return &Member{Code: qm.Code, WxOpenID: qm.WxOpenID}, nil
}
