package balance

import (
	"context"

	"github.com/twiglab/crm/bonus/orm/ent"
	"github.com/twiglab/crm/bonus/orm/ent/bonusacc"
	"github.com/twiglab/crm/bonus/orm/ent/bonusitem"
)

type BonusAcc struct {
	MemberCode       string
	LastCleanBalance int
	LastCleanTs      int64
}

type Balance struct {
	Client *ent.Client
}

func (b *Balance) calc(ctx context.Context, bacc BonusAcc) (int, error) {
	balance := bacc.LastCleanBalance
	lastCleanTs := bacc.LastCleanTs

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

func (b *Balance) CalcBalance(ctx context.Context, acc string) (int, error) {
	accq := b.Client.BonusAcc.Query()
	accq.Where(bonusacc.MemberCodeEQ(acc))
	bacc, err := accq.Only(ctx)

	if ent.IsNotFound(err) {
		return b.calc(ctx, BonusAcc{
			MemberCode: acc,
		})
	}
	if err != nil {
		return 0, nil
	}

	return b.calc(ctx, BonusAcc{
		MemberCode:       bacc.MemberCode,
		LastCleanBalance: bacc.LastCleanBalance,
		LastCleanTs:      bacc.LastCleanTs,
	})
}
