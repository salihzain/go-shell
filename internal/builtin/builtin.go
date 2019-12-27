package builtin

import (
	"errors"
	"fmt"
	"os"
)

// Helper is an alias for builtin functions
type Helper func([]string) error

// All returns all the available shell helper functions
func All() map[string]Helper {
	return map[string]Helper{
		"cd":   changeDir,
		"help": help,
		"exit": exit,
	}
}

// changeDir is a wrapper around cd "new dir"
// not that useful, but good example of using Go builtin os functions
func changeDir(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: new dir must be specified")
	}
	return os.Chdir(args[1])
}

func help(_ []string) error {
	fmt.Println("to run an executable, enter its name and arguments")
	fmt.Println("example: ls -a")
	return nil
}

func exit(_ []string) error {
	os.Exit(0)
	// unreachable code, but keeping it to satisfy the type
	return nil
}
