package iserv

type IServAccountConfig struct {
	IServHost string
	Username  string
	Password  string
}

type IServClientOptions struct {
	EnableWeb   bool
	EnableEmail bool
	EnableFiles bool
	AgentString string
}
