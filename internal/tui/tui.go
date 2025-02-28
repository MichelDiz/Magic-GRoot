package tui

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices     []string
	cursor      int
	selected    string
	projectPath string
}

func NewModel(projectPath string, scripts []string) model {
	return model{
		choices:     scripts,
		cursor:      0,
		projectPath: projectPath,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.choices[m.cursor]
			fmt.Println("\nRodando script:", m.selected)
			runScript(m.projectPath, m.selected)
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Selecione um script para executar:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPressione ↑/↓ para navegar, Enter para executar, Q para sair."
	return s
}

func RunTUI(projectPath string, scripts []string) {
	p := tea.NewProgram(NewModel(projectPath, scripts))
	if err := p.Start(); err != nil {
		fmt.Println("Erro ao iniciar a interface:", err)
		os.Exit(1)
	}
}

func runScript(projectPath, script string) {
	cmd := exec.Command("npm", "run", script)
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
