package types

// config
type AccountConfig struct {
	IServHost string
	Username  string
	Password  string
}

type ClientOptions struct {
	EnableModules map[string]bool
	AgentString   string
}
