package orm

import (
	"context"
	"time"

	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/member/orm/ent/member"
)

type Param struct {
	OpenID  string
	UnionID string

	Code string
}

type MemberDBOP struct {
	Client *ent.Client
}

func (op *MemberDBOP) GetMemberByCode(ctx context.Context, param Param) (*ent.Member, error) {
	q := op.Client.Member.Query()
	q.Where(member.CodeEQ(param.Code))
	m, err := q.Only(ctx)
	return m, err
}

func (op *MemberDBOP) WxLogin(ctx context.Context, param Param) (*ent.Member, bool, error) {
	q := op.Client.Member.Query()
	q.Where(member.WxOpenIDEQ(param.OpenID))
	m, err := q.Only(ctx)
	if ent.IsNotFound(err) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	u := m.Update()
	u.SetLastTime(time.Now())
	mb, err := u.Save(ctx)

	return mb, true, err
}

func (r *MemberDBOP) WxCreateMember(ctx context.Context, param Param) (*ent.Member, error) {
	c := r.Client.Member.Create().
		SetCode(param.Code).
		SetStatus(1). //设置有效
		SetLastTime(time.Now()).
		SetWxOpenID(param.OpenID).
		SetWxUnionID(param.UnionID)
	return c.Save(ctx)
}
