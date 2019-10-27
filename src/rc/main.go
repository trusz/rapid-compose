package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/trusz/rapid-compose/src/persist"

	"github.com/trusz/rapid-compose/src/dc"
	"github.com/trusz/rapid-compose/src/prompt"
	"github.com/trusz/rapid-compose/src/yaml"
)

func main() {

	showAll, inverse, reset, restart := parseFlags()

	if reset {
		persist.SaveSelections(emptyServiceList)
		return
	}

	if restart {
		InitRestart(showAll)
		return
	}

	oldSelection := persist.LoadSelections()
	possibleServices := yaml.LoadPossibleServicesNames(showAll)
	services := prompt.QuestionForStart(possibleServices, oldSelection, inverse)

	persist.SaveSelections(services)

	if len(services) > 0 {
		dc.Start(services)
	}

}

var emptyServiceList = []string{}

func parseFlags() (showAll bool, inverse bool, reset bool, restart bool) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Rapid Compose (rc) starts selected services.\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}

	_showAll := flag.Bool("a", false, "Show all service.")
	_inverse := flag.Bool("i", false, "Inverse selection. Start everything except selected ones.")
	_reset := flag.Bool("r", false, "Resets selected services.")
	_restart := flag.Bool("restart", false, "Restart selected services")
	flag.Parse()

	return *_showAll, *_inverse, *_reset, *_restart

}
