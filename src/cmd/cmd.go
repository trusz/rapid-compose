package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

// Run _
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

// Exec _
func Exec(command string) string {
	app := "sh"
	arg0 := "-c"

	// cmd := exec.Command(app, arg0, command)
	// stdout, _ := cmd.StdoutPipe()
	// stderr, _ := cmd.StderrPipe()

	// cmd.Run()

	out, err := exec.Command(app, arg0, command).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
