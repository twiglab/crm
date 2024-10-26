package orm

import (
	"context"
	"time"

	"github.com/twiglab/crm/adv/orm/ent"
	"github.com/twiglab/crm/adv/orm/ent/adv"
)

type Param struct {
	Index7 string
	Now    time.Time
}

type AdvDBOP struct {
	client *ent.Client
}

func NewAdvDBOP(client *ent.Client) *AdvDBOP {
	return &AdvDBOP{client: client}
}

func (dbop *AdvDBOP) SelectAdvNearBy(ctx context.Context, param Param) ([]*ent.Adv, error) {
	q := dbop.client.Adv.Query()
	q.Where(
		adv.H3Index7EQ(param.Index7),
		adv.StartTimeGT(param.Now),
		adv.EndTimeLT(param.Now),
		adv.StatusEQ(1),
	)
	// q.Order()

	return q.All(ctx)
}
