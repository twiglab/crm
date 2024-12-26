package balance

import (
	"context"

	"github.com/twiglab/crm/bonus/orm/ent"
	"github.com/twiglab/crm/bonus/orm/ent/bonusacc"
	"github.com/twiglab/crm/bonus/orm/ent/bonusitem"
)

type Balance struct {
	Client *ent.Client
}

func (b *Balance) Calc(ctx context.Context, acc string) (int, error) {
	accq := b.Client.BonusAcc.Query()
	accq.Where(bonusacc.MemberCodeEQ(acc))
	bacc, err := accq.Only(ctx)

	var (
		balance     int
		lastCleanTs int64
	)

	if err != nil && !ent.IsNotFound(err) {
		return balance, err
	} else {
		balance = bacc.LastCleanBalance
		lastCleanTs = bacc.LastCleanTs
	}

	lq := b.Client.BonusItem.Query()
	lq.Where(bonusitem.CreateTsGT(lastCleanTs))
	list, err := lq.All(ctx)

	if ent.IsNotFound(err) {
		return balance, nil
	}
	if err != nil {
		return balance, err
	}

	for _, l := range list {
		balance = balance + l.Bonus
	}

	return balance, nil
}
