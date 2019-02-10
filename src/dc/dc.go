package dc

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

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

func dcDown() {
	command := "docker-compose down"
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
