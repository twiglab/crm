package orm

import (
	"context"
	"time"

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

	BcmbWxMsgID string
	BcmbType    int
	BcmbRegTime time.Time
	BcmbCode    string
}

func insertNewMember(ctx context.Context, client *ent.Client, param Param) (*ent.Member, error) {
	cr := client.Member.Create()
	cr.SetCode(param.Code)
	cr.SetWxOpenID(param.WxOpenID)
	return cr.Save(ctx)
}

func selectByWxID(ctx context.Context, client *ent.Client, param Param) (*ent.Member, error) {
	q := client.Member.Query()
	q.Where(member.CodeEQ(param.WxOpenID))
	return q.Only(ctx)
}

func registWxMember(ctx context.Context, client *ent.Client, param Param) (*ent.Member, error) {
	update := client.Member.Update()

	update.SetLevel(1).
		Where(member.CodeEQ(param.WxOpenID))

	update.Exec(ctx)

	return nil, nil
}
