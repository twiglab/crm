package shop

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/erp/pkg/shop"
	"github.com/twiglab/crm/erp/pkg/shop/low"
)

type ShopDesc struct {
	MallCode string
	MallName string

	ShopCode string
	ShopName string
}

type ShopCli struct {
	Client graphql.Client
}

func (s *ShopCli) Qry(ctx context.Context, shopCode string) (ShopDesc, error) {
	resp, err := low.QryShopByCode(ctx, s.Client, shop.QryShopReq{ShopCode: shopCode})
	if err != nil {
		return ShopDesc{}, err
	}

	desc := resp.GetQryShopByCode()
	return ShopDesc{
		MallCode: desc.MallCode,
		MallName: desc.MallName,
		ShopCode: desc.ShopCode,
		ShopName: desc.ShopName,
	}, nil
}
