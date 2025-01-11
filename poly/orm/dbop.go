package orm

import (
	"context"

	"github.com/twiglab/crm/poly/orm/ent"
	"github.com/twiglab/crm/poly/orm/ent/poly"
)

type PolyDBOP struct {
	Client *ent.Client
}

// 获取活动列表
func (dbop *PolyDBOP) GetPolyPage(ctx context.Context, mallCode string) ([]*ent.Poly, error) {
	return dbop.Client.Poly.Query().Where(poly.MallCodeEQ(mallCode)).All(ctx)
}

// 活动详情
func (dbop *PolyDBOP) PolyDetail(ctx context.Context, code string) (*ent.Poly, error) {
	return dbop.Client.Poly.Query().Where(poly.CodeEQ(code)).Only(ctx)
}
