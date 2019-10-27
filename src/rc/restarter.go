package main

import (
	"github.com/trusz/rapid-compose/src/dc"
	"github.com/trusz/rapid-compose/src/prompt"
	"github.com/trusz/rapid-compose/src/yaml"
)

// InitRestart _
func InitRestart(showAll bool) {
	services := yaml.LoadPossibleServices()
	rcs := dc.FindRunningContainers()
	resServices := mapRestartableServices(services, rcs)

	message := "RAPID COMPOSE(RC) \nSelect services to re-start:"
	serviceNames := mapToServiceNames(resServices)
	emptySelection := []string{}
	servicesToRestart := prompt.Question(serviceNames, emptySelection, message)

	if len(servicesToRestart) == 0 {
		return
	}

	ids := mapToContainerIDs(resServices, servicesToRestart)

	dc.Restart(ids)
}

func mapToServiceNames(services RestartableServices) []string {
	names := make([]string, 0)

	for serviceName := range services {
		names = append(names, serviceName)
	}

	return names
}

func mapToContainerIDs(services RestartableServices, serviceNames []string) []string {
	ids := make([]string, 0)

	for _, name := range serviceNames {
		service := services[name]
		if service == nil {
			continue
		}

		ids = append(ids, service.ContainerID)
	}

	return ids
}

// MapRestartableServices _
func mapRestartableServices(
	services yaml.Services,
	rcs dc.RunningContainers,
) RestartableServices {
	rss := make(RestartableServices)

	for serviceName := range services {
		service := services[serviceName]
		containerID := rcs[service.Image]
		if containerID == "" {
			continue
		}

		rss[serviceName] = &RestartableService{
			Service:     service,
			ContainerID: containerID,
		}
	}

	return rss
}

// RestartableServices _
type RestartableServices map[string]*RestartableService

// RestartableService _
type RestartableService struct {
	yaml.Service
	ContainerID string
}

func filterOutNotOwnedImageNames(
	ownedImageNames []string,
	rcs dc.RunningContainers,
) dc.RunningContainers {
	ownedContainers := make(dc.RunningContainers)

	for _, im := range ownedImageNames {
		containerID := rcs[im]
		if containerID == "" {
			continue
		}

		ownedContainers[im] = containerID
	}

	return ownedContainers
}
