package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
)

const fileName = ".rapid-compose"
const defaultContent = "[]"

var filePath = homeFolder() + "/" + fileName

// Write _
func Write(selections []Selection) {
	content, _ := json.Marshal(selections)
	writeFile(filePath, content)
}

// Read _
func Read() []Selection {
	var content = readFile(filePath)
	var selections = make([]Selection, 0)
	err := json.Unmarshal(content, &selections)
	if err != nil {
		panic(err)
	}
	return selections
}

// Selection _
type Selection struct {
	Path     string   `json:"path"`
	Services []string `json:"services"`
}

func writeFile(path string, content []byte) {

	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func readFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte(defaultContent)
	}
	return content
}

func homeFolder() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
