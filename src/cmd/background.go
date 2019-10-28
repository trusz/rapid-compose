package cmd

import (
	"os/exec"
)

// Background starts the given command in a background process
func Background(command string) {

	app := "sh"
	arg0 := "-c"

	cmd := exec.Command(app, arg0, command)
	// stdout, _ := cmd.StdoutPipe()
	// stderr, _ := cmd.StderrPipe()

	cmd.Start()

	// go func() {
	// 	scanner := bufio.NewScanner(stdout)
	// 	for scanner.Scan() {
	// 		m := scanner.Text()
	// 		fmt.Println(m)
	// 	}
	// }()

	// go func() {
	// 	scanner := bufio.NewScanner(stderr)
	// 	for scanner.Scan() {
	// 		m := scanner.Text()
	// 		fmt.Println(m)
	// 	}
	// }()

	// return cmd
}
