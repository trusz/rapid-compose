package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/yaml.v2"
)

func main() {

	possibleServices := loadPossibleServices()
	services := question(possibleServices)
	start(services)

}

type dockerCompose struct {
	Services map[string]interface{} `yaml:"services"`
}

func loadPossibleServices() []string {
	yamlFile, err := ioutil.ReadFile("docker-compose.yaml")
	if err != nil {
		yamlFile, err = ioutil.ReadFile("docker-compose.yml")
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}
	}

	var dc = &dockerCompose{}
	err = yaml.Unmarshal(yamlFile, dc)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	var possibleServices = make([]string, 0)
	for service := range dc.Services {
		possibleServices = append(possibleServices, service)
	}

	return possibleServices
}

func question(possibleServices []string) []string {
	choosenServices := []string{}
	prompt := &survey.MultiSelect{
		Message: "RAPID COMPOSE(RC) \nSelect services to start:",
		Options: possibleServices,
	}
	fmt.Println("Rapid Compose")
	survey.AskOne(prompt, &choosenServices, nil)

	return choosenServices
}

func start(services []string) {
	app := "sh"
	arg0 := "-c"
	arg1 := "docker-compose up -d " + strings.Join(services[:], " ")

	cmd := exec.Command(app, arg0, arg1)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		println("error")
		fmt.Printf("log %q \n", stderr.String())
		log.Fatal(err)
	}

	fmt.Printf("log %q\n", stdout.String())

	// cmd := exec.Command("tr", "a-z", "A-Z")
	// cmd.Stdin = strings.NewReader("some input")
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("in all caps: %q\n", out.String())
}
