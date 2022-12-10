package types

import "fmt"

type ClientConfig struct {
	EnableModules map[string]bool
	AgentString   string
	SaveSessions  bool

	IServHost string
	Username  string
	Password  string
}

func (c *ClientConfig) IServUrl() string {
	return fmt.Sprintf("https://%s/iserv", c.IServHost)
}
