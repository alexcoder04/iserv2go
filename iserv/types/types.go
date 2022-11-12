package types

import "time"

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

// other
type IServExercise struct {
	Type    string
	Title   string
	DueDate time.Time
	Tags    []string
	Teacher string
}
