package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"

	"github.com/twiglab/crm/bonus/gql/graph/model"
)

// BalanceDetail is the resolver for the balanceDetail field.
func (r *queryResolver) BalanceDetail(ctx context.Context, input *model.BalanceReq) (*model.BalanceDetailResp, error) {
	b, err := r.Balance.CalcBalance(ctx, input.MemberCode)
	return &model.BalanceDetailResp{MemberCode: input.MemberCode, Balance: b}, err
}