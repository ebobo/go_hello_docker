package main

import (
	"fmt"
	"time"
)

type greetingCmd struct {
	Name string `short:"n" long:"name" description:"Your name, for a greeting" default:"Unknown"`
}

// Execute export
func (r *greetingCmd) Execute(args []string) error {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello", r.Name, i+1)
		time.Sleep(3 * time.Second)
	}
	return nil
}
