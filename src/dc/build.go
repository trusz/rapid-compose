package dc

import (
	"strings"

	"github.com/trusz/rapid-compose/src/cmd"
)

// StartDetached _
func Build(services []string) {
	command := "docker compose build " + strings.Join(services[:], " ")
	buildCommand := cmd.Run(command)
	buildCommand.Wait()
}
