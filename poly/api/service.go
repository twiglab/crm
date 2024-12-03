package api

import (
	"context"
	"fmt"

	"github.com/twiglab/crm/poly/internal/coupon"
	"github.com/twiglab/crm/poly/orm"
	"github.com/twiglab/crm/poly/orm/ent"
)

// TODO FIX
var (
	p   = orm.NewPolyDBOP(&ent.Client{})
	ctx = context.Background()
)

type PolyService struct{}

func NewPolyService() *PolyService {
	return &PolyService{}
}

// ActivityList 活动列表
//
//	@param mallCode 商场code
//	@return []*ent.Poly 活动信息
//	@return error
func (s *PolyService) ActivityList(mallCode string) ([]*ent.Poly, error) {
	l, err := p.GetPolyList(ctx, mallCode)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// ActivityDetail 活动详情
//
//	@param code 活动code
//	@return *ent.Poly 活动信息
//	@return error
func (s *PolyService) ActivityDetail(code string) (*ent.Poly, error) {
	d, err := p.GetPolyDetail(ctx, code)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// ActivityCreate 活动创建
//
//	@param param
//	@return *ent.Poly
//	@return error
func (s *PolyService) ActivityCreate(param *orm.CreatePolyParam) (*ent.Poly, error) {
	d, err := p.CreatePoly(ctx, param)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// ActivityQualificationVerify 活动参与资格验证
//
//	@param activityCode 活动code
//	@param input
//	@return bool 可否领劵
//	@return string 卷token
//	@return int 卷id
//	@return error
func (s *PolyService) ActivityQualificationVerify(activityCode string, input any) (bool, string, int, error) {
	d, err := p.GetPolyDetail(ctx, activityCode)
	if err != nil {
		return false, "", 0, err
	}

	if d.Status != orm.ACTIVITYSTATUSOPEN {
		return false, "", 0, fmt.Errorf("活动未开放")
	}

	couponToken, err := coupon.GenerateToken()
	if err != nil {
		return false, "", 0, fmt.Errorf("活动未开放")
	}

	couponId := 123

	return true, couponToken, couponId, nil
}

// ActivityCouponTokenVerify 活动卷token验证
//
//	@param couponToken
//	@return bool
//	@return error
func (s *PolyService) ActivityCouponTokenVerify(couponToken string) (bool, error) {
	b := coupon.CheckToken(couponToken)
	return b, nil
}

// ActivityCouponPreOccupy 活动卷预占用
//
//	@param couponToken
//	@return bool
//	@return error
func (s *PolyService) ActivityCouponPreOccupy(couponToken string) (bool, error) {
	coupon.PreUseToken(couponToken)
	return true, nil
}
