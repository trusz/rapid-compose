package cmd

import (
	"bufio"
	"fmt"
	"os/exec"
)

// Run starts the command in a new thread,
// writes out stdout and stderr outputs
// and returns the command object
func Run(command string) *exec.Cmd {

	app := "sh"
	arg0 := "-c"

	cmd := exec.Command(app, arg0, command)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	cmd.Start()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}
	}()

	return cmd
}
