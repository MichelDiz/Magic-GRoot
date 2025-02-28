package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var DryRun bool = false

func RunScript(projectPath, script string) {
	runner := GetPreferredRunner()

	var cmdArgs []string
	if NeedsRunPrefix(runner) {
		cmdArgs = append(cmdArgs, "run", script)
	} else {
		cmdArgs = append(cmdArgs, script)
	}

	if runner == "bash" && !strings.HasPrefix(script, "./") {
		script = "./" + script
		cmdArgs = []string{script}
	}

	cmd := exec.Command(runner, cmdArgs...)
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if DryRun {
		fmt.Printf(" (Dry-run) Executaria: cd %s && %s %s\n", projectPath, runner, strings.Join(cmdArgs, " "))
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf(" Erro ao executar '%s' com %s em %s\n", script, runner, projectPath)
	}
}
