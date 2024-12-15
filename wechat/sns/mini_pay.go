package sns

import (
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
)

var (
	wx_client *wechat.ClientV3
)

func GetClient() *wechat.ClientV3 {
	return wx_client
}

func InitWeChatClient(MchId, SerialNo, APIv3Key, PrivateKey string) {
	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	client, err := wechat.NewClientV3(MchId, SerialNo, APIv3Key, PrivateKey)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		xlog.Error(err)
		return
	}

	wx_client = client
}
