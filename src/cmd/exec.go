package cmd

import (
	"log"
	"os/exec"
)

// Exec runs the command and returns the output
func Exec(command string) string {
	app := "sh"
	arg0 := "-c"

	out, err := exec.Command(app, arg0, command).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
