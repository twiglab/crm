package cmdb

type JWTConfig struct {
	Alg    string `json:"alg" yaml:"alg" mapstructure:"alg"`
	Secret string `json:"secret" yaml:"secret" mapstructure:"secret"`
}
