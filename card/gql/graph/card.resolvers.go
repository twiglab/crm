package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"
	"fmt"
	"time"

	"github.com/twiglab/crm/card/gql/graph/model"
	"github.com/twiglab/crm/card/orm/ent"
	"github.com/twiglab/crm/card/orm/ent/card"
)

// BindCard is the resolver for the bindCard field.
func (r *mutationResolver) BindCard(ctx context.Context, input model.BindCardReq) (*model.CardResp, error) {
	q, err := r.Client.Card.Query().Where(card.CodeEQ(input.Code)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("card code not found")
		}
		return nil, err
	}
	if q.MemberCode != "" {
		return nil, fmt.Errorf("card already binded")
	}

	c := r.Client.Card.UpdateOne(q)
	c.SetMemberCode(input.MemberCode)
	c.SetBindTime(time.Now())
	cobj, err := c.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.CardResp{
		Code:          cobj.Code,
		CodeBin:       cobj.CardBin,
		Type:          cobj.Type,
		Pic1:          cobj.Pic1,
		Pic2:          cobj.Pic2,
		Amount:        int(cobj.Amount),
		MemberCode:    &cobj.MemberCode,
		LastCleanTime: nil,
		Status:        cobj.Status,
	}, nil
}

// ActiveCard is the resolver for the activeCard field.
func (r *mutationResolver) ActiveCard(ctx context.Context, input model.ActiveCardReq) (*model.CardResp, error) {
	q, err := r.Client.Card.Query().Where(card.CardBinEQ(input.CodeBin)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("card code not found")
		}
		return nil, err
	}

	c := r.Client.Card.UpdateOne(q)
	c.SetStatus(1)
	cobj, err := c.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.CardResp{
		Code:          cobj.Code,
		CodeBin:       cobj.CardBin,
		Type:          cobj.Type,
		Pic1:          cobj.Pic1,
		Pic2:          cobj.Pic2,
		Amount:        int(cobj.Amount),
		MemberCode:    &cobj.MemberCode,
		LastCleanTime: nil,
		Status:        cobj.Status,
	}, nil
}

// GetChargeRecordCode is the resolver for the getChargeRecordCode field.
func (r *mutationResolver) GetChargeRecordCode(ctx context.Context, input model.ChargeRecordCodeReq) (*model.ChargeRecordCodeResp, error) {
	// 验证card 和 用户关系
	cobj, err := r.Client.Card.Query().Where(card.CodeEQ(input.CardCode)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("card code not found")
		}
		return nil, err
	}
	if cobj.MemberCode != input.MemberCode {
		return nil, fmt.Errorf("card not binded to member")
	}

	pc, err := r.Cache.GetPayCode(ctx, r.Client, input.CardCode)
	if err != nil {
		return nil, err
	}

	return &model.ChargeRecordCodeResp{Code: pc}, nil
}

// QueryCardDetail is the resolver for the queryCardDetail field.
func (r *queryResolver) QueryCardDetail(ctx context.Context, input *model.QueryCardByCode) (*model.CardResp, error) {
	cobj, err := r.Client.Card.Query().Where(card.CodeEQ(input.Code)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("card code not found")
		}
		return nil, err
	}

	return &model.CardResp{
		Code:          cobj.Code,
		CodeBin:       cobj.CardBin,
		Type:          cobj.Type,
		Pic1:          cobj.Pic1,
		Pic2:          cobj.Pic2,
		Amount:        int(cobj.Amount),
		MemberCode:    &cobj.MemberCode,
		LastCleanTime: nil,
		Status:        cobj.Status,
	}, nil
}

// QueryCardList is the resolver for the queryCardList field.
func (r *queryResolver) QueryCardList(ctx context.Context, input *model.PaginationReq) ([]*model.CardResp, error) {
	q := r.Client.Card.Query()
	// q.Where(card.)
	if input.Order != nil && *input.Order == "desc" {
		q.Order(ent.Desc(card.FieldCode)).Where(card.CodeGT(*input.Cursor))
	} else {
		q.Order(ent.Asc(card.FieldCode)).Where(card.CodeLT(*input.Cursor))
	}
	if input.Limit != nil {
		q.Limit(*input.Limit)
	} else {
		q.Limit(10)
	}
	cobjs, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	var rets []*model.CardResp
	for _, cobj := range cobjs {
		obj := &model.CardResp{
			Code:          cobj.Code,
			CodeBin:       cobj.CardBin,
			Type:          cobj.Type,
			Pic1:          cobj.Pic1,
			Pic2:          cobj.Pic2,
			Amount:        int(cobj.Amount),
			MemberCode:    &cobj.MemberCode,
			LastCleanTime: nil,
			Status:        cobj.Status,
		}
		rets = append(rets, obj)
	}

	return rets, nil
}

// QueryCardByMemberCode is the resolver for the queryCardByMemberCode field.
func (r *queryResolver) QueryCardByMemberCode(ctx context.Context, input *model.QueryCardByMemberCode) ([]*model.CardResp, error) {
	cobjs, err := r.Client.Card.Query().
		Where(card.MemberCodeEQ(input.MemberCode)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var rets []*model.CardResp
	for _, cobj := range cobjs {
		obj := &model.CardResp{
			Code:       cobj.Code,
			CodeBin:    cobj.CardBin,
			MemberCode: &cobj.MemberCode,
			Type:       cobj.Type,
			Amount:     int(cobj.Amount),
		}
		rets = append(rets, obj)
	}
	return rets, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
