package graph

import "github.com/twiglab/crm/wechat/mq"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Q *mq.XMQ
}
