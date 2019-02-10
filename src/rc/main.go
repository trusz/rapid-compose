package main

import (
	"flag"

	"github.com/trusz/rapid-compose/src/dc"
	"github.com/trusz/rapid-compose/src/prompt"
	"github.com/trusz/rapid-compose/src/yaml"
)

func main() {

	var showDependencies = flag.Bool("d", false, "Show dependencies")
	flag.Parse()

	possibleServices := yaml.LoadPossibleServices(*showDependencies)
	services := prompt.Question(possibleServices)

	if len(services) > 0 {
		dc.Start(services)
	}

}
