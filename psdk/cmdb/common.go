package cmdb

type App struct {
	ID      string            `json:"id" yaml:"id" mapstructure:"id"`
	Debug   string            `json:"debug" yaml:"debug" mapstructure:"debug"`
	LogFile string            `json:"log-file" yaml:"log-file" mapstructure:"log-file"`
	Configs map[string]string `json:"configs" mapstructure:"configs" yaml:"configs"`
}

type Endpoint struct {
	Address string            `json:"address" yaml:"address" mapstructure:"address"`
	Configs map[string]string `json:"configs" mapstructure:"configs" yaml:"configs"`
}
