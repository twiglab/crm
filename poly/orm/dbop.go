package orm

import (
	"context"
	"time"

	"github.com/twiglab/crm/poly/orm/ent"
	"github.com/twiglab/crm/poly/orm/ent/poly"
)

type PolyDBOP struct {
	client *ent.Client
}

func NewPolyDBOP(client *ent.Client) *PolyDBOP {
	return &PolyDBOP{client: client}
}

// 活动状态
const (
	ACTIVITYSTATUSNEEDAPPROVE = iota + 1 // 等待审批
	ACTIVITYSTATUSWAIT                   // 未开启
	ACTIVITYSTATUSOPEN                   // 已开启
	ACTIVITYSTATUSEND                    // 已结束
	ACTIVITYSTATUSDISCARD                // 已废弃

)

// 审批状态
const (
	ACTIVITYAPPROVEWAIT     = iota + 1 // 等待审批
	ACTIVITYAPPROVEAGREE               // 审批同意
	ACTIVITYAPPROVEREJECTED            // 审批拒绝
)

type CreatePolyParam struct {
	MallCode  string
	Operator  string
	Name      string
	Desc      string
	Budget    int64
	StartTime *time.Time
	EndTime   *time.Time
}

// 创建活动
func (dbop *PolyDBOP) CreatePoly(ctx context.Context, param *CreatePolyParam) (*ent.Poly, error) {
	a, err := dbop.client.Poly.
		Create().
		SetMallCode(param.MallCode).
		SetOperator(param.Operator).
		SetName(param.Name).
		SetAddTime(time.Now()).
		SetDesc(param.Desc).
		SetBudget(param.Budget).
		SetStartTime(*param.StartTime).
		SetEndTime(*param.EndTime).
		SetStatus(ACTIVITYSTATUSNEEDAPPROVE).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 获取活动列表
func (dbop *PolyDBOP) GetPolyList(ctx context.Context, mallCode string) ([]*ent.Poly, error) {
	return dbop.client.Poly.Query().Where(poly.MallCodeEQ(mallCode)).All(ctx)
}

// 审批
func (dbop *PolyDBOP) ApprovePoly(ctx context.Context, code, approver string, status int) error {
	// TODO 流程部分

	if status == ACTIVITYAPPROVEAGREE {
		// TODO 活动开启计时器

		// 更新数据
		_, err := dbop.client.Poly.Update().Where(poly.CodeEQ(code)).SetStatus(ACTIVITYSTATUSWAIT).Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// // 变更申请请求
// type ChangePolyParam struct {
// 	PolyCode   string
// 	Operator   string
// 	SubmitTime time.Time
// 	Summary    string
// 	Reason     string

// 	ChangeData map[string]any
// }

// // 变更申请 TODO
// func (dbop *PolyDBOP) ChangePoly(ctx context.Context, param ChangePolyParam) error {
// 	// 活动存在
// 	if _, err := dbop.client.Poly.Query().Where(poly.CodeEQ(param.PolyCode)).First(ctx); err != nil {
// 		return err
// 	}

// 	return nil
// }

// 活动开始
func (dbop *PolyDBOP) StartPoly(ctx context.Context, code string) error {
	// _, err := dbop.client.Poly.Update().Where(poly.CodeEQ(code)).SetStartTime(time.Now()).SetStatus(ACTIVITYSTATUSOPEN).Save(ctx)
	_, err := dbop.client.Poly.Update().Where(poly.CodeEQ(code)).SetStartTime(time.Now()).SetStatus(ACTIVITYSTATUSOPEN).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// 活动结束
func (dbop *PolyDBOP) EndPoly(ctx context.Context, code string) error {
	_, err := dbop.client.Poly.Update().Where(poly.CodeEQ(code)).SetEndTime(time.Now()).SetStatus(ACTIVITYSTATUSEND).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// 活动详情
func (dbop *PolyDBOP) GetPolyDetail(ctx context.Context, code string) (*ent.Poly, error) {
	return dbop.client.Poly.Query().Where(poly.CodeEQ(code)).First(ctx)
}
