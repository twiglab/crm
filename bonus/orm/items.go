package orm

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/twiglab/crm/bonus/orm/ent"
	"github.com/twiglab/crm/bonus/orm/ent/bonusitem"
)

type ItemsPageParam struct {
	MemberCode string
	Last       string
	Size       int
}

type Items struct {
	Client *ent.Client
}

func (i *Items) ItemsPage(ctx context.Context, param ItemsPageParam) ([]*ent.BonusItem, error) {
	if param.Last == "" {
		param.Last = "!"
	}

	if param.Size == 0 {
		param.Size = 10
	}

	q := i.Client.BonusItem.Query()
	q.Where(
		bonusitem.CodeGT(param.Last),
		bonusitem.MemberCodeEQ(param.MemberCode),
	).
		Limit(param.Size).
		Order(bonusitem.ByCode(sql.OrderDesc()))

	return q.All(ctx)
}
