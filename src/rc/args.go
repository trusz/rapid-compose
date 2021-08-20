package main

import (
	"os"
	"strings"
)

func hasBuildCommand() bool {
	return hasCommand("build")
}

func hasCommand(cmd string) bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if strings.ToLower(arg) == cmd {
			return true
		}
	}

	return false
}
