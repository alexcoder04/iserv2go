package iserv

import "time"

type IServExercise struct {
	Type    string
	Title   string
	DueDate time.Time
	Tags    []string
	Teacher string
}
