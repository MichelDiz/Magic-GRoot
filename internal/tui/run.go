package tui

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func runTUI(model tea.Model) {
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Println("Erro ao iniciar a interface:", err)
		os.Exit(1)
	}
}

func RunTUI(projectPath string, scripts []string) {
	runTUI(NewModel(projectPath, scripts))
}

func RunAliasTUI(projects []string) {
	runTUI(NewAliasModel(projects))
}

func RunListTUI(projects []string) {
	runTUI(NewListModel(projects))
}

func RunSetRootTUI() {
	runTUI(NewSetRootModel())
}

func RunAliasManagerTUI() {
	runTUI(NewAliasManagerModel())
}

func clearScreen() {
	cmd := exec.Command("clear") //! ver se funciona para todos OSes Unix/Linux/MacOS
	cmd.Stdout = os.Stdout
	cmd.Run()
}
