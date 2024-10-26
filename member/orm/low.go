package orm

import (
	"context"
	"errors"

	"github.com/it512/box"
	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/member/orm/ent/member"
)

var EntClientKey = box.RndKey()

func EntClient(ctx context.Context) *ent.Client {
	return box.MustFrom[*ent.Client](ctx, EntClientKey)
}

type Param struct {
	Code     string
	WxOpenID string
}

func insertNewMember(ctx context.Context, client *ent.Client, param Param) (*ent.Member, error) {
	cr := client.Member.Create()
	cr.SetCode(param.Code)
	cr.SetWxOpenID(param.WxOpenID)
	return cr.Save(ctx)
}

func selectByWxID(ctx context.Context, client *ent.Client, param Param) (*ent.Member, error) {
	if param.WxOpenID == "" {
		return nil, errors.New("用户为空")
	}

	q := client.Member.Query()
	q.Where(member.CodeEQ(param.WxOpenID))
	return q.Only(ctx)
}
