package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"
	"fmt"

	"github.com/twiglab/crm/bonus/gql/graph/model"
)

// QueryBonusDetail is the resolver for the queryBonusDetail field.
func (r *queryResolver) QueryBonusDetail(ctx context.Context, input model.BonusDetailReq) (*model.BonusDetailResp, error) {
	b, err := r.Balance.Calc(ctx, input.MemberCode)
	return &model.BonusDetailResp{MemberCode: input.MemberCode, Balance: b}, err
}

// ListBonusItems is the resolver for the listBonusItems field.
func (r *queryResolver) ListBonusItems(ctx context.Context, input model.BonusListReq) (*model.BonusListResp, error) {
	panic(fmt.Errorf("not implemented: ListBonusItems - listBonusItems"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }