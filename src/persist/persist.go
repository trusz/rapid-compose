package persist

import (
	"log"
	"os"

	"github.com/trusz/rapid-compose/src/storage"
)

// Selection _
type Selection = storage.Selection

// SaveSelections _
func SaveSelections(services []string) {

	var selections = storage.Read()
	var si = findIndexByPath(selections, currentDir())
	var newSelection = Selection{
		Path:     currentDir(),
		Services: services,
	}

	if si >= 0 {
		selections[si] = newSelection
	} else {
		selections = append(selections, newSelection)
	}

	storage.Write(selections)

}

// LoadSelections _
func LoadSelections() []string {
	var selections = storage.Read()
	var si = findIndexByPath(selections, currentDir())
	if si >= 0 {
		return selections[si].Services
	}
	return []string{}
}

func currentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func findIndexByPath(
	selections []Selection,
	path string,
) int {
	for si, selection := range selections {
		if selection.Path == path {
			return si
		}
	}

	return -1
}
