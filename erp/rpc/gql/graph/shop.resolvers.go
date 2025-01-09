package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"

	"github.com/twiglab/crm/erp/orm/ent/shop"
	sp "github.com/twiglab/crm/erp/pkg/shop"
)

// QryShopByCode is the resolver for the qryShopByCode field.
func (r *queryResolver) QryShopByCode(ctx context.Context, input sp.QryShopReq) (*sp.ShopResp, error) {
	q := r.Client.Shop.Query()
	q.Where(shop.ShopCodeEQ(input.ShopCode))
	s, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	return &sp.ShopResp{
		ShopCode: s.ShopCode,
		ShopName: s.ShopName,
		MallCode: s.MallCode,
		MallName: s.MallName,

		Floor: s.Floor,
		Pos:   s.Pos,

		ContractCode: s.ContractCode,

		BizClass1:     s.BizClass1,
		BizClassName1: s.BizClassName1,
		BizClass2:     s.BizClass2,
		BizClassName2: s.BizClassName2,
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
