package config

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/bonus/bc"
	"github.com/twiglab/crm/bonus/mb"
	"github.com/twiglab/crm/bonus/shop"
)

func BcCli(rpc GqlRpc) *bc.WxBcCli {
	return &bc.WxBcCli{
		Client: graphql.NewClient(rpc.RpcUrl, http.DefaultClient),
	}
}
func MbCli(rpc GqlRpc) *mb.MemberCli {
	return &mb.MemberCli{
		Client: graphql.NewClient(rpc.RpcUrl, http.DefaultClient),
	}
}

func ShopCli(rpc GqlRpc) *shop.ShopCli {
	return &shop.ShopCli{
		Client: graphql.NewClient(rpc.RpcUrl, http.DefaultClient),
	}
}
