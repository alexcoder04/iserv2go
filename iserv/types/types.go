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
	Types       []string
	Title       string
	DueDate     time.Time
	StartDate   time.Time
	Description string
	Files       []string
	Tags        []string
	Teacher     string
}
