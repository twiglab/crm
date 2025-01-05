package orm

import (
	"context"
	"time"

	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/member/orm/ent/member"
)

type MemberDBOP struct {
	Client *ent.Client
}

func (op *MemberDBOP) GetMemberByCode(ctx context.Context, code string) (*ent.Member, error) {
	q := op.Client.Member.Query()
	q.Where(member.CodeEQ(code))
	m, err := q.Only(ctx)
	return m, err
}

func (op *MemberDBOP) WxLogin(ctx context.Context, openID string) (*ent.Member, bool, error) {
	q := op.Client.Member.Query()
	q.Where(member.WxOpenIDEQ(openID))
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

func (r *MemberDBOP) CreateWxMember(ctx context.Context, code, openID string) (*ent.Member, error) {
	c := r.Client.Member.Create().
		SetCode(code).
		SetStatus(1). //设置有效
		SetLastTime(time.Now()).
		SetWxOpenID(openID)
	return c.Save(ctx)
}
