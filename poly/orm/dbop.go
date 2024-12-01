package orm

import (
	"context"
	"encoding/json"
	"time"

	"github.com/twiglab/crm/poly/orm/ent"
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

type CreateActivityParam struct {
	MallCode  string
	Operator  string
	Name      string
	Desc      string
	Budget    int64
	StartTime *time.Time
	EndTime   *time.Time
}

// 创建活动
func (dbop *PolyDBOP) CreateActivity(ctx context.Context, param CreateActivityParam) (*ent.Activity, error) {
	a, err := dbop.client.Activity.
		Create().
		SetMallCode(param.MallCode).
		SetOperator(param.Operator).
		SetActivityName(param.Name).
		SetActivityAddTime(time.Now()).
		SetActivityDesc(param.Desc).
		SetActivityBudget(param.Budget).
		SetActivityStartTime(*param.StartTime).
		SetActivityEndTime(*param.EndTime).
		SetActivityStatus(ACTIVITYSTATUSNEEDAPPROVE).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 获取活动列表
func (dbop *PolyDBOP) GetActivityList(ctx context.Context, mallCode string) ([]*ent.Activity, error) {
	return dbop.client.Activity.Query().Where(activity.MallCodeEQ(mallCode)).All(ctx)
}

// 审批
func (dbop *PolyDBOP) ApproveActivity(ctx context.Context, code, approver string, status int) error {
	// TODO 流程部分

	if status == ACTIVITYAPPROVEAGREE {
		// TODO 活动开启计时器

		// 更新数据
		_, err := dbop.client.Activity.Update().Where(activity.CodeEQ(code)).SetActivityApproveTime(time.Now()).SetActivityStatus(ACTIVITYSTATUSWAIT).SetApprover(approver).Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// 变更申请请求
type ChangeActivityParam struct {
	ActivityCode string
	Operator     string
	SubmitTime   time.Time
	Summary      string
	Reason       string

	ChangeData map[string]any
}

// 变更申请
func (dbop *PolyDBOP) ChangeActivity(ctx context.Context, param ChangeActivityParam) error {
	// 活动存在
	if _, err := dbop.client.Activity.Query().Where(activity.CodeEQ(param.ActivityCode)).First(ctx); err != nil {
		return err
	}

	_, err := dbop.client.ActivityChange.Query().Where(activitychange.ActivityCodeEQ(param.ActivityCode), activitychange.Status(ACTIVITYSTATUSNEEDAPPROVE)).First(ctx)
	if err == nil {
		return ErrExistApprovingChangeRequest
	}
	if !ent.IsNotFound(err) {
		return err
	}

	jsonData, err := json.Marshal(param.ChangeData)
	if err != nil {
		return err
	}

	// 添加变更记录
	_, err = dbop.client.ActivityChange.Create().
		SetActivityCode(param.ActivityCode).
		SetOperator(param.Operator).
		SetSubmitTime(param.SubmitTime).
		SetChangeSummary(param.Summary).
		SetChangeReason(param.Reason).
		SetChangeRecord(string(jsonData)).
		SetStatus(ACTIVITYAPPROVEWAIT).
		Save(ctx)

	return nil
}

// 变更申请审批
func (dbop *PolyDBOP) ApproveChangeActivity(ctx context.Context, ActivityCode, approver string, status int) error {
	// TODO 流程部分

	changeObj, err := dbop.client.ActivityChange.Query().Where(activitychange.ActivityCodeEQ(ActivityCode), activitychange.StatusEQ(ACTIVITYAPPROVEWAIT)).First(ctx)
	if err != nil {
		return ErrChangeRequestNotExistOrApproved
	}

	dbop.client.ActivityChange.UpdateOne(changeObj).SetStatus(status).Save(ctx)

	if status == ACTIVITYAPPROVEAGREE {
		dbop.client.ActivityChange.UpdateOne(changeObj).SetApprover(approver).SetApproveTime(time.Now()).Save(ctx)

		// TODO 更新活动数据
	}

	return nil
}

// 活动开始
func (dbop *PolyDBOP) StartActivity(ctx context.Context, code string) error {
	_, err := dbop.client.Activity.Update().Where(activity.CodeEQ(code)).SetActivityStartTime(time.Now()).SetActivityStatus(ACTIVITYSTATUSOPEN).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// 活动结束
func (dbop *PolyDBOP) EndActivity(ctx context.Context, code string) error {
	_, err := dbop.client.Activity.Update().Where(activity.CodeEQ(code)).SetActivityEndTime(time.Now()).SetActivityStatus(ACTIVITYSTATUSEND).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
