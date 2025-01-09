package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"

	"github.com/twiglab/crm/bonus/gql/graph/model"
	"github.com/twiglab/crm/bonus/orm"
)

// BonusPage is the resolver for the bonusPage field.
func (r *queryResolver) BonusPage(ctx context.Context, input model.BonusPageReq) (*model.BonusPageResp, error) {
	a, err := r.Items.ItemsPage(ctx, orm.ItemsPageParam{
		MemberCode: input.MemberCode,
		Last:       input.Last,
	})
	if err != nil {
		return nil, err
	}
	size := len(a)
	is := make([]*model.BonusItem, size)
	for i, item := range a {
		is[i] = &model.BonusItem{
			Code:   item.Code,
			Amount: item.Amount,
		}
	}

	last := input.Last
	if size > 0 {
		last = is[size-1].Code
	}

	return &model.BonusPageResp{
		BonusItems: is,
		Limit:      input.Limit,
		Last:       last,
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
