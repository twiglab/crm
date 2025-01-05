package config

type App struct {
	ID  string `yaml:"id" mapstructure:"id"`
	Web struct {
		Addr string `mapstructure:"addr"`
	} `yaml:"web" mapstructure:"web"`
	Jwt struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`

	Wechat struct {
		RpcUrl string `mapstructure:"rpc-url"`
	} `mapstructure:"wechat"`

	Member struct {
		RpcUrl string `mapstructure:"rpc-url"`
	} `mapstructure:"member"`
}
