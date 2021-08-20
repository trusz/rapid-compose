package dc

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/trusz/rapid-compose/src/cmd"
)

// FindRunningContainers _
func FindRunningContainers() RunningContainers {
	command := "docker ps --format \"{{.ID}}\t{{.Image}}\""
	output := cmd.Exec(command)
	lines := strings.Split(output, "\n")

	containers := make(RunningContainers)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		containerID := fields[0]
		imageName := fields[1]

		containers[imageName] = containerID

	}

	return containers
}

// RunningContainers _
type RunningContainers map[ImageName]ContainerID

// ImageName _
type ImageName = string

// ContainerID _
type ContainerID = string

func dcDown() {
	command := "docker compose down"
	cmd.Run(command).Wait()
}

func waitForInterrupt() chan bool {
	signal.Ignore(os.Interrupt)
	signal.Ignore(syscall.SIGINT)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	wait := make(chan bool, 1)

	go func() {
		for range sigc {
			wait <- true
		}
	}()

	return wait

}
