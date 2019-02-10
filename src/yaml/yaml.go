package yaml

import (
	"io/ioutil"
	"log"
	"sort"

	"gopkg.in/yaml.v2"
)

type dockerCompose struct {
	Services map[string]service `yaml:"services"`
}

type service struct {
	DependsOn []string `yaml:"depends_on"`
}

// LoadPossibleServices _
func LoadPossibleServices(showDependencies bool) []string {
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
	var dependencyServices = make([]string, 0)

	for service := range dc.Services {
		dependencyServices = append(dependencyServices, dc.Services[service].DependsOn...)
		possibleServices = append(possibleServices, service)
	}

	var filteredServices = possibleServices
	if !showDependencies {
		filteredServices = filterOutDependencies(possibleServices, dependencyServices)
	}
	sort.Strings(filteredServices)
	return filteredServices
}

func filterOutDependencies(services []string, dependencies []string) []string {
	var filteredServices = make([]string, 0)
	for _, service := range services {
		if !contains(dependencies, service) {
			filteredServices = append(filteredServices, service)
		}
	}

	return filteredServices
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
