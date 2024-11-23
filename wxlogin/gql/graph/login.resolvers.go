package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"

	"github.com/twiglab/crm/wxlogin/gql/graph/model"
)

// Login is the resolver for the Login field.
func (r *queryResolver) Login(ctx context.Context, input model.JsCodeReq) (*model.LoginUserResp, error) {
	x, err := r.AuthCli.Login2(ctx, input.JsCode)
	if err != nil {
		return nil, err
	}

	return &model.LoginUserResp{Code: x.Code, Jwt: x.Jwt}, err
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
