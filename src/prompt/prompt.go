package prompt

import (
	"gopkg.in/AlecAivazis/survey.v1"
)

// QuestionForStart _
func QuestionForStart(
	possibleServices []string,
	oldSelection []string,
	inverse bool,
) []string {
	var message = "RAPID COMPOSE(RC) \nSelect services to start:"
	if inverse {
		message = "RAPID COMPOSE(RC) \nSelect services NOT to start:"
	}

	chosenServices := Question(possibleServices, oldSelection, message)

	if inverse {
		return inverseSelection(possibleServices, chosenServices)
	}

	return chosenServices
}

// Question _
func Question(
	possibleServices []string,
	oldSelection []string,
	message string,
) []string {

	chosenServices := []string{}
	prompt := &survey.MultiSelect{
		Message: message,
		Options: possibleServices,
		Default: oldSelection,
	}
	survey.AskOne(prompt, &chosenServices, nil)

	return chosenServices
}

func inverseSelection(
	possibleServices []string,
	chosenServices []string,
) []string {
	var services = make([]string, 0)
	for _, possibleService := range possibleServices {
		if !contains(chosenServices, possibleService) {
			services = append(services, possibleService)
		}
	}

	return services
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
