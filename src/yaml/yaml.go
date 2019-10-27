package yaml

import (
	"io/ioutil"
	"log"
	"sort"

	"gopkg.in/yaml.v2"
)

type dockerCompose struct {
	Services Services `yaml:"services"`
}

// Services _
type Services map[string]Service

// Service _
type Service struct {
	DependsOn []string `yaml:"depends_on"`
	Image     string   `yaml:"image"`
}

// LoadPossibleServices _
func LoadPossibleServices() Services {
	dc := parseDCFile()
	return dc.Services
}

// LoadPossibleServicesNames _
func LoadPossibleServicesNames(showAll bool) []string {

	dc := parseDCFile()

	var possibleServices = make([]string, 0)
	var dependencyServices = make([]string, 0)

	for service := range dc.Services {
		dependencyServices = append(dependencyServices, dc.Services[service].DependsOn...)
		possibleServices = append(possibleServices, service)
	}

	var filteredServices = possibleServices
	if !showAll {
		filteredServices = filterOutDependencies(possibleServices, dependencyServices)
	}
	sort.Strings(filteredServices)
	return filteredServices
}

func parseDCFile() dockerCompose {
	dcFile := readDockerComposeFile()

	var dc = &dockerCompose{}
	err := yaml.Unmarshal(dcFile, dc)
	if err != nil {
		log.Fatalf("Error -> Unmarshal: %v", err)
	}

	return *dc
}

func readDockerComposeFile() []byte {
	yamlFile, err := ioutil.ReadFile("docker-compose.yaml")
	if err != nil {
		yamlFile, err = ioutil.ReadFile("docker-compose.yml")
		if err != nil {
			log.Printf("Error -> Yaml file read: #%v ", err)
		}
	}

	return yamlFile
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
