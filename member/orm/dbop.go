package orm

import (
	"context"
	"errors"

	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/member/orm/ent/member"
)

type MemberDBOP struct {
	client *ent.Client
}

func NewMemberDBOP(client *ent.Client) *MemberDBOP {
	return &MemberDBOP{client: client}
}

func (op *MemberDBOP) InsertNewMember(ctx context.Context, param Param) (*ent.Member, error) {
	cr := op.client.Member.Create()
	cr.SetCode(param.Code)
	cr.SetWxOpenID(param.WxOpenID)
	return cr.Save(ctx)
}

func (op *MemberDBOP) SelectByCode(ctx context.Context, param Param) (*ent.Member, error) {
	if param.Code == "" {
		return nil, errors.New("用户为空")
	}

	q := EntClient(ctx).Member.Query()
	q.Where(member.CodeEQ(param.Code))
	return q.Only(ctx)
}

func (op *MemberDBOP) SelectByWxID(ctx context.Context, param Param) (*ent.Member, error) {
	if param.WxOpenID == "" {
		return nil, errors.New("用户为空")
	}

	q := op.client.Member.Query()
	q.Where(member.CodeEQ(param.WxOpenID))
	return q.Only(ctx)
}
