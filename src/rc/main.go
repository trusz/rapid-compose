package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/trusz/rapid-compose/src/dc"
	"github.com/trusz/rapid-compose/src/prompt"
	"github.com/trusz/rapid-compose/src/yaml"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Rapid Compose (rc) starts selected services.\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}

	var showAll = flag.Bool("a", false, "Show all service.")
	var inverse = flag.Bool("i", false, "Inverse selection. Start everything except selected ones.")
	flag.Parse()

	possibleServices := yaml.LoadPossibleServices(*showAll)
	services := prompt.Question(possibleServices, *inverse)

	if len(services) > 0 {
		dc.Start(services)
	}

}
