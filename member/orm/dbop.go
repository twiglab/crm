package orm

import (
	"context"
	"errors"

	"github.com/it512/box"
	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/member/orm/ent/member"
)

type MemberDBOP struct {
	client *ent.Client
}

func NewMemberDBOP(client *ent.Client) *MemberDBOP {
	return &MemberDBOP{client: client}
}

var EntClientKey = box.RndKey()

func EntClient(ctx context.Context) *ent.Client {
	return box.MustFrom[*ent.Client](ctx, EntClientKey)
}

type Param struct {
	Code     string
	WxOpenID string
}

func InsertNewMember(ctx context.Context, param Param) (*ent.Member, error) {
	cr := EntClient(ctx).Member.Create()
	cr.SetCode(param.Code)
	cr.SetWxOpenID(param.WxOpenID)
	return cr.Save(ctx)
}

func SelectByCode(ctx context.Context, param Param) (*ent.Member, error) {
	if param.Code == "" {
		return nil, errors.New("用户为空")
	}

	q := EntClient(ctx).Member.Query()
	q.Where(member.CodeEQ(param.Code))
	return q.Only(ctx)
}

func UpdatePhone(ctx context.Context, code, phone string) error {
	u := EntClient(ctx).Member.Update()
	u.SetPhone(phone)
	u.Where(member.CodeEQ(code))
	return u.Exec(ctx)
}
