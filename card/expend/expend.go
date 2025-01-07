package expend

import (
	"github.com/twiglab/crm/card/orm/ent"
)

type extend struct {
	IExpendCache
	ICalcTask
}

var (
	cardCache *expendVar
	expendIns *extend
)

func InitModel(client *ent.Client) *extend {
	cardCache = &expendVar{
		client: client,
	}
	expendIns = &extend{
		IExpendCache: cardCache,
	}
	return expendIns
}
