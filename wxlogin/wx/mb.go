package wx

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/google/uuid"
	"github.com/twiglab/crm/member/pkg/data"
	"github.com/twiglab/crm/member/pkg/low"
	"github.com/twiglab/crm/psdk/code"
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
	req, err := low.WxLogin(ctx, c.Client, data.OpenIDReq{WxOpenID: wxOpenID})
	if err != nil {
		return nil, err
	}

	m := req.GetWxLogin()
	if m.Found {
		return &Member{Code: m.Code, WxOpenID: m.WxOpenID}, nil
	}

	// 没找到, 直接新建
	code := code.Code36()
	if wxOpenID == "0000000000-0000000000" {
		code = uuid.Nil.String()
	}
	// 创建用户
	r, err := low.CreateWxMember(ctx, c.Client, data.CreateWxMemberReq{Code: code, WxOpenID: wxOpenID})
	if err != nil {
		return nil, err
	}
	qm := r.GetCreateWxMember()
	return &Member{Code: qm.Code, WxOpenID: qm.WxOpenID}, nil
}
