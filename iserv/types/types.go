package types

// config
type IServAccountConfig struct {
	IServHost string
	Username  string
	Password  string
}

type IServClientOptions struct {
	EnableModules map[string]bool
	AgentString   string
}
