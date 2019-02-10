package prompt

import (
	"fmt"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Question _
func Question(possibleServices []string) []string {
	choosenServices := []string{}
	prompt := &survey.MultiSelect{
		Message: "RAPID COMPOSE(RC) \nSelect services to start:",
		Options: possibleServices,
	}
	fmt.Println("Rapid Compose")
	survey.AskOne(prompt, &choosenServices, nil)

	return choosenServices
}
