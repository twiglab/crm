package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"

	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/member/orm/ent/member"
	"github.com/twiglab/crm/member/pkg/data"
)

// CreateWxMember is the resolver for the createWxMember field.
func (r *mutationResolver) CreateWxMember(ctx context.Context, input data.CreateWxMemberReq) (*data.MemberResp, error) {
	q := r.Client.Member.Query()
	q.Where(member.WxOpenIDEQ(input.WxOpenID))
	b, err := q.Exist(ctx)
	if err != nil {
		return nil, err
	}

	if b {
		return &data.MemberResp{Code: input.Code, WxOpenID: input.WxOpenID}, nil
	}

	c := r.Client.Member.Create()
	c.SetCode(input.Code)
	c.SetWxOpenID(input.WxOpenID)
	mb, err := c.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &data.MemberResp{Code: mb.Code, WxOpenID: mb.WxOpenID}, nil
}

// QueryWxMemberByOpenID is the resolver for the queryWxMemberByOpenID field.
func (r *queryResolver) QueryWxMemberByOpenID(ctx context.Context, input data.OpenIDReq) (*data.QryMemberResp, error) {
	q := r.Client.Member.Query()
	q.Where(member.WxOpenIDEQ(input.WxOpenID))
	mb, err := q.Only(ctx)
	if ent.IsNotFound(err) {
		return &data.QryMemberResp{Found: false}, nil
	}

	if err != nil {
		return &data.QryMemberResp{Found: false}, err
	}

	return &data.QryMemberResp{Code: mb.Code, WxOpenID: mb.WxOpenID, Found: true}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
