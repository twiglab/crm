package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type MQConfig struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}

type BusinessCircleConfig struct {
	MchId      string `yaml:"mcd_id" mapstructure:"mcd_id"`
	SerialNo   string `yaml:"serial_no" mapstructure:"serial_no"`
	APIKey     string `yaml:"api_key" mapstructure:"api_key"` // 微信商户平台—>账户设置—>API安全—>设置APIv3密钥
	PrivateKey string `yaml:"private_key" mapstructure:"private_key"`
}

type WechatConfig struct {
	AppId     string `yaml:"app_id" mapstructure:"app_id"`
	AppSecret string `yaml:"app_secret" mapstructure:"app_secret"`
}

type AppConfig struct {
	Name string `yaml:"name" mapstructure:"name"`
	Addr string `yaml:"addr" mapstructure:"addr"`
}

type GlobalConfig struct {
	App    AppConfig            `yaml:"app" mapstructure:"app"`
	MQ     MQConfig             `yaml:"mq" mapstructure:"mq"`
	BC     BusinessCircleConfig `yaml:"bc" mapstructure:"bc"`
	Wechat WechatConfig         `yaml:"wechat" mapstructure:"wechat"`
}

var globalConfig *GlobalConfig

func GetConfig() *GlobalConfig {
	return globalConfig
}

func InitConfig() error {
	viper.SetConfigName("notify-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("配置文件读取失败,err:%s", err)
	}

	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		return fmt.Errorf("配置文件转换失败,err:%s", err)
	}
	return nil
}
