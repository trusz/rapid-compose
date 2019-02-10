package main

import (
	"github.com/trusz/rapid-compose/src/dc"
	"github.com/trusz/rapid-compose/src/prompt"
	"github.com/trusz/rapid-compose/src/yaml"
)

func main() {

	possibleServices := yaml.LoadPossibleServices()
	services := prompt.Question(possibleServices)

	if len(services) > 0 {
		dc.Start(services)
	}

}
