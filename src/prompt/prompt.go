package prompt

import (
	"gopkg.in/AlecAivazis/survey.v1"
)

// Question _
func Question(possibleServices []string, inverse bool) []string {

	var message = "RAPID COMPOSE(RC) \nSelect services to start:"
	if inverse {
		message = "RAPID COMPOSE(RC) \nSelect services NOT to start:"
	}

	choosenServices := []string{}
	prompt := &survey.MultiSelect{
		Message: message,
		Options: possibleServices,
	}
	survey.AskOne(prompt, &choosenServices, nil)

	if inverse {
		return inverseSelection(possibleServices, choosenServices)
	}

	return choosenServices
}

func inverseSelection(
	possibleServices []string,
	choosenServices []string,
) []string {
	var services = make([]string, 0)
	for _, possibleService := range possibleServices {
		if !contains(choosenServices, possibleService) {
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
