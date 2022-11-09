package iserv

type IServAccountConfig struct {
	IServHost string
	Username  string
	Password  string
}

type IServClientOptions struct {
	EnableEmail bool
	AgentString string
}
