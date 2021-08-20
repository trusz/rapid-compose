package dc

import (
	"strings"

	"github.com/trusz/rapid-compose/src/cmd"
)

// StartDetached _
func StartDetached(services []string) {
	command := "docker compose up " + strings.Join(services[:], " ")
	cmd.Background(command)
}
