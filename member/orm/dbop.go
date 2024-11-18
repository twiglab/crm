package orm

import (
	"context"
	"errors"

	"github.com/twiglab/crm/member/orm/ent"
)

type MemberDBOP struct {
	Client *ent.Client
}

func (op *MemberDBOP) InsertNewMember(ctx context.Context, param Param) (*ent.Member, error) {
	return insertNewMember(ctx, op.Client, param)
}

func (op *MemberDBOP) SelectByWxID(ctx context.Context, param Param) (*ent.Member, error) {
	if param.WxOpenID == "" {
		return nil, errors.New("用户为空")
	}

	m, err := selectByWxID(ctx, op.Client, param)
	if ent.IsNotFound(err) {
		return nil, nil
	}

	return m, err
}

func (op *MemberDBOP) SelectByWxID2(ctx context.Context, param Param) (*ent.Member, bool, error) {

	m, err := op.SelectByWxID(ctx, param)
	if ent.IsNotFound(err) {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	return m, true, nil
}
