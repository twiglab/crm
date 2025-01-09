package graph

import (
	"github.com/twiglab/crm/bonus/balance"
	"github.com/twiglab/crm/bonus/orm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Balance *balance.Balance
	Items   *orm.Items
}
