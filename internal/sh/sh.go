package sh

import (
	"github.com/sigmazain/go-shell/internal/builtin"

	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"syscall"
)

var (
	reader       = bufio.NewReader(os.Stdin)
	builtins     = builtin.All()
	childProcess *exec.Cmd
)

func readLine() (string, error) {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		return "", err
	}

	// get rid of '\n'
	return string(line[:len(line)-1]), nil
}

func splitLine(line string) ([]string, bool) {
	// -1 to return all substrings
	s := regexp.MustCompile("(\"[^\"]+\"|[^ ]+)+").FindAllString(line, -1)
	return s, len(s) > 0
}

func execute(args []string) error {
	if len(args) == 0 {
		return errors.New("program name must be passed")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Env = os.Environ()

	// fork and execute the process in a child process
	if err := cmd.Start(); err != nil {
		return err
	}

	// wait until the child process finish executing
	childProcess = cmd
	err := cmd.Wait()

	return err
}

func handleSignals() {
	// kill the child process if an interupt or terminate or quit signals are received
	// this ensures that the parent process (the shell) isn't killed
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	for {
		s := <-signalChannel
		if (s == syscall.SIGINT ||
			s == syscall.SIGTERM ||
			s == syscall.SIGQUIT) &&
			childProcess != nil {
		} else {
			fmt.Printf(" to exit go-shell, enter \"exit\"\npress enter/return to continue\n")
		}
	}

}

// Loop is the main loop for go-shell
func Loop() {
	var (
		line   string
		args   []string
		argsOK bool
		err    error
	)

	// setup the signal handler
	go handleSignals()

	for {
		fmt.Printf("go-shell-> ")

		line, err = readLine()
		if err != nil {
			log.Printf("failed to read input: %v\n", err)
			continue
		}

		args, argsOK = splitLine(line)
		if !argsOK {
			log.Println("invalid input, try again")
			continue
		}

		// check if the command is for a builtin function before executing it
		if f, ok := builtins[args[0]]; ok {
			err = f(args)
			if err != nil {
				log.Printf("failed to execute builtin: %v", err)
			}
			continue
		}

		err = execute(args)
		if err != nil {
			log.Printf("process execution failed with err: %v", err)
			continue
		}
	}
}
