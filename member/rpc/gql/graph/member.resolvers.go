package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"fmt"

	"github.com/twiglab/crm/member/orm"
	"github.com/twiglab/crm/member/pkg/data"
)

// CreateWxMember is the resolver for the CreateWxMember field.
func (r *mutationResolver) CreateWxMember(ctx context.Context, input data.CreateWxMemberReq) (*data.MemberResp, error) {
	panic(fmt.Errorf("not implemented: CreateWxMember - CreateWxMember"))
}

// QueryWxMember is the resolver for the QueryWxMember field.
func (r *queryResolver) QueryWxMember(ctx context.Context, input data.OpenIDReq) (*data.MemberResp, error) {
	m, err := r.dbop.SelectByWxID(ctx, orm.Param{WxOpenID: input.WxOpenID})
	if m != nil && err == nil {
		return &data.MemberResp{Code: m.Code}, nil
	}

	return nil, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
