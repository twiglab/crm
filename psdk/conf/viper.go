package conf

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

type ConfigCtx struct {
	viper *viper.Viper
	context.Context
	fmt.Stringer
}

func WithContext(ctx context.Context) *ConfigCtx {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	return &ConfigCtx{viper: v, Context: ctx}
}

func (c *ConfigCtx) SetConfigName(name string) {
	c.viper.SetConfigName(name)
}

func (c *ConfigCtx) SetConfigType(typ string) {
	c.viper.SetConfigType(typ)
}

func (c *ConfigCtx) AddConfigPath(paths ...string) {
	for _, path := range paths {
		c.viper.AddConfigPath(path)
	}
}

func (c *ConfigCtx) ReadInConfig() error {
	return c.viper.ReadInConfig()
}

func (c *ConfigCtx) Unmarshal(v any) error {
	return c.viper.Unmarshal(v)
}

func (c *ConfigCtx) Value(key any) any {
	vk, ok := key.(string)

	if !ok {
		return c.Context.Value(key)
	}

	if val := c.viper.Get(vk); val != nil {
		return val
	}
	return c.Context.Value(key)
}

func (c *ConfigCtx) String() string {
	return "ctx.Config"
}
