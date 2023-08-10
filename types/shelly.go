package types

import (
	"net"
)

type ShellyDevice struct {
	Name       string `json:"name" yaml:"name"`
	FQName     string `json:"fq_name" yaml:"fq_name"`
	App        string `json:"app" yaml:"app"`
	Version    string `json:"ver" yaml:"ver"`
	Generation string `json:"gen" yaml:"gen"`
	IPv4       net.IP `json:"ipv4,omitempty" yaml:"ipv4,omitempty"`
	IPv6       net.IP `json:"ipv6,omitempty" yaml:"ipv6,omitempty"`
}
