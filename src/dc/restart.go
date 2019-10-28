package dc

import (
	"strings"

	"github.com/trusz/rapid-compose/src/cmd"
)

// Restart _
func Restart(containerIDs []string) {
	command := "docker restart " + strings.Join(containerIDs[:], " ")
	restartCommand := cmd.Run(command)
	restartCommand.Wait()
}
