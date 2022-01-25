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
		switch i {
		case 0:
			fmt.Printf("Good Morning ! %v \n", r.Name)
		case 1:
			fmt.Printf("Good Afternoon ! %v \n", r.Name)
		default:
			fmt.Printf("Good Night ! %s \n", r.Name)
		}

		time.Sleep(2 * time.Second)
	}
	fmt.Printf("Bye Bye ! %v \n", r.Name)
	return nil
}
