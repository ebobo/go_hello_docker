package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Greeting greetingCmd `command:"greeting" description:"Say hello to user"`
}

// use go run . to start not go run main.go
func main() {
	var opt Options
	p := flags.NewParser(&opt, flags.Default)
	_, err := p.Parse()
	switch flagsErr := err.(type) {
	case flags.ErrorType:
		if flagsErr == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)

	default:
		os.Exit(1)
	}
}
