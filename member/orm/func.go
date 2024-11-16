package orm

import (
	"context"
	"errors"

	"github.com/twiglab/crm/member/orm/ent"
	"github.com/twiglab/crm/member/orm/ent/member"
)

func InsertNewMember(ctx context.Context, param Param) (*ent.Member, error) {
	client := EntClient(ctx)
	return insertNewMember(ctx, client, param)
}

func SelectByWxID(ctx context.Context, param Param) (*ent.Member, error) {
	if param.WxOpenID == "" {
		return nil, errors.New("用户为空")
	}

	client := EntClient(ctx)
	return selectByWxID(ctx, client, param)
}

func Cr(ctx context.Context, param Param) (*ent.Member, error) {
	client := EntClient(ctx)
	m, err := selectByWxID(ctx, client, param)
	if ent.IsNotFound(err) {
		return insertNewMember(ctx, client, param)
	}
	return m, err
}

func SelectByCode(ctx context.Context, param Param) (*ent.Member, error) {
	if param.Code == "" {
		return nil, errors.New("用户为空")
	}

	q := EntClient(ctx).Member.Query()
	q.Where(member.CodeEQ(param.Code))
	return q.Only(ctx)
}

func UpdatePhone(ctx context.Context, code, phone string) error {
	u := EntClient(ctx).Member.Update()
	u.SetPhone(phone)
	u.Where(member.CodeEQ(code))
	return u.Exec(ctx)
}
