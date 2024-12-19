package bc

import (
	"context"
	"errors"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
)

var (
	bcClient *BcClient
)

type BcClient struct {
	client *wechat.ClientV3
}

func GetClient() *BcClient {
	return bcClient
}

func InitWeChatClient(MchId, SerialNo, APIv3Key, PrivateKey string) error {
	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	client, err := wechat.NewClientV3(MchId, SerialNo, APIv3Key, PrivateKey)
	if err != nil {
		return err
	}

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		return err
	}

	bcClient = new(BcClient)
	bcClient.client = client
	return nil
}

func NewBcClient(client *wechat.ClientV3) *BcClient {
	return &BcClient{client: client}
}

func (c *BcClient) BCPointsSync(ctx context.Context, query BusinessCirclePointsSync) error {
	var bm gopay.BodyMap
	bm.Set("transaction_id", query.TransactionID)
	bm.Set("appid", query.AppID)
	bm.Set("openid", query.OpenID)
	bm.Set("earn_points", query.EarnPoints)
	bm.Set("increased_points", query.IncreasedPoints)
	bm.Set("points_update_time", query.PointsUpdateTime)
	result, err := c.client.V3BusinessPointsSync(ctx, bm)
	if result.Code != 0 {
		return errors.New(result.Error)
	}
	return err
}

/*
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
*/
