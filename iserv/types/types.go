package types

import "time"

// config
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

// other
type IServExercise struct {
	Type    string
	Title   string
	DueDate time.Time
	Tags    []string
	Teacher string
}
