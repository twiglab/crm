package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"

	"github.com/twiglab/crm/wechat/mq"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"github.com/twiglab/crm/wechat/pkg/bc/msg"
	"github.com/twiglab/crm/wechat/testbed/graph/model"
)

// TestAuth is the resolver for the testAuth field.
func (r *queryResolver) TestAuth(ctx context.Context, input msg.BusinessCircleAuthor) (*model.Result, error) {
	body, err := mq.FromMsgp(&input)
	if err != nil {
		return nil, err
	}

	err = r.Q.Send(ctx, bc.MQ_WX_TOC_BC_AUTH, body)

	return &model.Result{MsgID: input.MsgID}, err
}

// TestPayment is the resolver for the testPayment field.
func (r *queryResolver) TestPayment(ctx context.Context, input msg.BusinessCirclePayment) (*model.Result, error) {
	body, err := mq.FromMsgp(&input)
	if err != nil {
		return nil, err
	}

	err = r.Q.Send(ctx, bc.MQ_WX_TOC_BC_PAYMENT, body)

	return &model.Result{MsgID: input.MsgID}, err
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
