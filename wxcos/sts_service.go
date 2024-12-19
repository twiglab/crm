package wxcos

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/viper"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

type stsConfig struct {
	// 微信账号
	Appid  string `yaml:"appid" mapstructure:"appid"`
	Bucket string `yaml:"bucket" mapstructure:"bucket"`
	// 子账号密钥
	SecretId  string `yaml:"secret_id" mapstructure:"secret_id"`
	SecretKey string `yaml:"secret_key" mapstructure:"secret_key"`
	// 存储桶
	Host   string `yaml:"host" mapstructure:"host"`
	Region string `yaml:"region" mapstructure:"region"`

	DurationSeconds int64 `yaml:"duration_seconds" mapstructure:"duration_seconds"` // 链接有效时间

	// http.client配置
	ClientTimeout             int `yaml:"client_timeout" mapstructure:"client_timeout"`
	ClientMaxIdleConnsPerHost int `yaml:"client_max_idle_conns_per_host" mapstructure:"client_max_idle_conns_per_host"`
}

func readConfig() *stsConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}

	var c stsConfig
	err = viper.Unmarshal(&c)
	if err != nil {
		panic("配置文件转换失败")
	}

	return &c
}

var once sync.Once

type stsClient struct {
	config *stsConfig
	client *sts.Client
	opt    *sts.CredentialOptions
}

var stsInstance *stsClient

func GetStsInstance() *stsClient {
	once.Do(func() {
		config := readConfig()

		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConnsPerHost = config.ClientMaxIdleConnsPerHost
		hc := http.Client{Timeout: time.Duration(config.ClientTimeout) * time.Second, Transport: t}

		client := sts.NewClient(config.SecretId, config.SecretKey, &hc, sts.Host(config.Host))
		// 配置来源：https://github.com/tencentyun/qcloud-cos-sts-sdk/blob/master/go/example/sts_demo.go
		opt := &sts.CredentialOptions{
			DurationSeconds: config.DurationSeconds,
			Region:          config.Region,
			Policy: &sts.CredentialPolicy{
				Statement: []sts.CredentialPolicyStatement{
					{
						Action: []string{
							"name/cos:PostObject",
						},
						Effect: "allow",
						Resource: []string{
							"qcs::cos:" + config.Region + ":uid/" + config.Appid + ":" + config.Bucket + "/exampleobject",
						},
					},
				},
			},
		}

		stsInstance = &stsClient{config: config, client: client, opt: opt}
	})
	return stsInstance
}

// GetCredentials 获取临时凭证
func (s *stsClient) GetCredentials() *sts.CredentialResult {
	res, err := s.client.GetCredential(s.opt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", res.Credentials)

	return res
}
