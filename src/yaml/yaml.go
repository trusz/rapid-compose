package yaml

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type dockerCompose struct {
	Services map[string]interface{} `yaml:"services"`
}

// LoadPossibleServices _
func LoadPossibleServices() []string {
	yamlFile, err := ioutil.ReadFile("docker-compose.yaml")
	if err != nil {
		yamlFile, err = ioutil.ReadFile("docker-compose.yml")
		if err != nil {
			log.Printf("Error -> Yaml file read: #%v ", err)
		}
	}

	var dc = &dockerCompose{}
	err = yaml.Unmarshal(yamlFile, dc)
	if err != nil {
		log.Fatalf("Error -> Unmarshal: %v", err)
	}

	var possibleServices = make([]string, 0)
	for service := range dc.Services {
		possibleServices = append(possibleServices, service)
	}

	return possibleServices
}
