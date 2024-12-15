package bc

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/twiglab/crm/wechat/pkg/werr"
	"github.com/twiglab/crm/wechat/sns"
)

func BCPointsSync(ctx context.Context, query BusinessCirclePointsSync) error {
	var bm gopay.BodyMap
	bm.Set("transaction_id", query.TransactionID)
	bm.Set("appid", query.AppID)
	bm.Set("openid", query.OpenID)
	bm.Set("earn_points", query.EarnPoints)
	bm.Set("increased_points", query.IncreasedPoints)
	bm.Set("points_update_time", query.PointsUpdateTime)
	result, err := sns.GetClient().V3BusinessPointsSync(ctx, bm)
	if err != nil {
		return &werr.WeChatCommonError{
			ErrCode: int64(result.Code),
			ErrMsg:  result.Error,
		}
	}
	return nil
}

func BCAuthPointsQuery(ctx context.Context, query BusinessCircleAuthorQuery) (*wechat.BusinessAuthPointsQuery, error) {
	var bm gopay.BodyMap
	bm.Set("appid", query.AppID)
	result, err := sns.GetClient().V3BusinessAuthPointsQuery(ctx, query.OpenID, bm)
	if err != nil {
		return nil, &werr.WeChatCommonError{
			ErrCode: int64(result.Code),
			ErrMsg:  result.Error,
		}
	}
	return result.Response, nil
}

func BCPointsStatusQuery(ctx context.Context, query BusinessCirclePointsStatusQuery) (*wechat.BusinessPointsStatusQuery, error) {
	var bm gopay.BodyMap
	bm.Set("brandid", query.BrandID)
	bm.Set("appid", query.AppID)
	result, err := sns.GetClient().V3BusinessPointsStatusQuery(ctx, query.OpenID, bm)
	if err != nil {
		return nil, &werr.WeChatCommonError{
			ErrCode: int64(result.Code),
			ErrMsg:  result.Error,
		}
	}
	return result.Response, nil
}

func BCParkingSync(ctx context.Context, query BusinessCircleParkingSyncQuery) error {
	var bm gopay.BodyMap
	bm.Set("brandid", query.BrandID)
	bm.Set("appid", query.AppID)
	bm.Set("openid", query.OpenID)
	bm.Set("plate_number", query.PlateNumber)
	bm.Set("state", query.State)
	bm.Set("time", query.Time)
	result, err := sns.GetClient().V3BusinessParkingSync(ctx, bm)
	if err != nil {
		return &werr.WeChatCommonError{
			ErrCode: int64(result.Code),
			ErrMsg:  result.Error,
		}
	}
	return nil
}