package dc

import (
	"strings"

	"github.com/trusz/rapid-compose/src/cmd"
)

// Start _
func Start(services []string) {
	command := "docker-compose up " + strings.Join(services[:], " ")
	upCommand := cmd.Run(command)
	<-waitForInterrupt()
	upCommand.Wait()
	dcDown()
}
