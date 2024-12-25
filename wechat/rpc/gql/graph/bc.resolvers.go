package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"

	"github.com/twiglab/crm/wechat/pkg/bc"
)

// BcPointSync is the resolver for the bcPointSync field.
func (r *queryResolver) BcPointSync(ctx context.Context, input bc.BusinessCirclePointsSync) (*bc.BcVoid, error) {
	err := r.BcClinet.BCPointsSync(ctx, input)
	return &bc.BcVoid{}, err
}