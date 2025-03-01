package tui

import (
	"fmt"
	"mgr/internal/scanner"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type listModel struct {
	choices []string
	cursor  int
}

func NewListModel(projects []string) listModel {
	return listModel{
		choices: projects,
		cursor:  0,
	}
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			fmt.Println("\n Operação cancelada.")
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
			selectedProject := m.choices[m.cursor]
			fmt.Println("\n Projeto selecionado:", selectedProject)

			scripts := scanner.ScanForScripts(selectedProject)
			if scriptList, exists := scripts[selectedProject]; exists && len(scriptList) > 0 {
				RunTUI(selectedProject, scriptList)
			} else {
				fmt.Println("\n Nenhum script encontrado para esse projeto.")
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m listModel) View() string {
	s := "\n Selecione um projeto:\n\n"

	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "=>"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n Use as setas para navegar, Enter para selecionar, Q para sair."
	return s
}
