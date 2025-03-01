package tui

import (
	"fmt"
	"mgr/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices     []string
	cursor      int
	selected    string
	projectPath string
	quitting    bool
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
		case "q", "esc":
			m.quitting = true
			return m, tea.Quit
		case "ctrl+c":
			m.quitting = true
			return m, tea.Interrupt
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			clearScreen()
			m.selected = m.choices[m.cursor]
			fmt.Println("\nRodando script:", m.selected)
			utils.RunScript(m.projectPath, m.selected)
			// clearScreen()
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return RenderList("Selecione um script para executar:", m.choices, m.cursor)
}
