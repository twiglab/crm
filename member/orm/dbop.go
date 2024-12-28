package orm

import (
	"context"

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
