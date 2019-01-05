package ssh

import "fmt"

type Endpoint struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

type SshTunnelConfig struct {
	Local  *Endpoint `yaml:"local"`
	Target *Endpoint `yaml:"target"`
	Remote *Endpoint `yaml:"remote"`
}

type YamlConfig struct {
	Tunnels []*SshTunnelConfig `yaml:"tunnels"`
}
