package graph

import (
	"github.com/twiglab/crm/wechat/bc"
	"github.com/twiglab/crm/wechat/sns"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Auth *sns.Auth
	BcClinet *bc.BcClient
}
