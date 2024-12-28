package cmdb

import (
	"log"
	"net"
)

type Binding struct {
	Address  string            `json:"address" yaml:"address" mapstructure:"address"`
	Protocol string            `json:"protocol" yaml:"protocol" mapstructure:"protocol"`
	Configs  map[string]string `json:"configs" mapstructure:"configs" yaml:"configs"`
}

func NewListen(binding Binding) net.Listener {
	s := "tcp"
	if binding.Protocol != "" {
		s = binding.Protocol
	}

	l, err := net.Listen(s, binding.Address)
	if err != nil {
		log.Fatal(err)
	}

	return l
}
