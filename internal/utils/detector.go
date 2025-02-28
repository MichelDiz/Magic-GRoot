package utils

import (
	"fmt"
	"mgr/internal/config"
	"os/exec"
)

func IsCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func DetectAndSaveRunners() {
	availableRunners := []string{}

	if IsCommandAvailable("npm") {
		availableRunners = append(availableRunners, "npm")
	}
	if IsCommandAvailable("yarn") {
		availableRunners = append(availableRunners, "yarn")
	}
	if IsCommandAvailable("pnpm") {
		availableRunners = append(availableRunners, "pnpm")
	}
	if IsCommandAvailable("bash") {
		availableRunners = append(availableRunners, "bash")
	}

	if len(availableRunners) > 0 {
		config.SetConfig("available_runners", fmt.Sprintf("%v", availableRunners))
		fmt.Println(" Runners detectados e salvos:", availableRunners)
	} else {
		fmt.Println(" Nenhum gerenciador de scripts detectado.")
	}
}

func GetPreferredRunner() string {
	runner := config.GetConfig("preferred_runner")
	if runner == "" {

		available := config.GetConfig("available_runners")
		if available != "" {
			runner = available[0:3]
		} else {
			runner = "npm"
		}
		config.SetConfig("preferred_runner", runner)
	}
	return runner
}

func NeedsRunPrefix(runner string) bool {
	switch runner {
	case "npm", "pnpm":
		return true
	case "yarn", "bash":
		return false
	default:
		return true
	}
}
