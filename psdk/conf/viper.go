package conf

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

type ConfigCtx struct {
	*viper.Viper
	context.Context
	fmt.Stringer
}

func WithContext(ctx context.Context) *ConfigCtx {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	return &ConfigCtx{Viper: v, Context: ctx}
}

func (c *ConfigCtx) AddConfigPaths(paths ...string) {
	for _, path := range paths {
		c.Viper.AddConfigPath(path)
	}
}

func (c *ConfigCtx) Value(key any) any {
	vk, ok := key.(string)

	if !ok {
		return c.Context.Value(key)
	}

	if val := c.Viper.Get(vk); val != nil {
		return val
	}
	return c.Context.Value(key)
}

func (c *ConfigCtx) String() string {
	return "ctx.Config"
}
