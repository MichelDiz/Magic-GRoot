package tui

import (
	"fmt"
	"mgr/internal/scanner"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
)

type listModel struct {
	choices  []string
	cursor   int
	quitting bool
}

func NewListModel(projects []string) listModel {
	sort.Strings(projects) // Ordena os projetos em ordem alfabética
	return listModel{
		choices: projects,
		cursor:  0,
	}
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd, handled := UpdateListHandler(msg, &m.cursor, m.choices, &m.quitting, nil, func(selected string) tea.Cmd {
		clearScreen()
		fmt.Println("\n Projeto selecionado:", selected)
		scripts := scanner.ScanForScripts(selected)
		if scriptList, exists := scripts[selected]; exists && len(scriptList) > 0 {
			scriptArray := make([]string, 0, len(scriptList))
			for script, command := range scriptList {
				scriptArray = append(scriptArray, fmt.Sprintf("%s: %s", script, command))
			}
			sort.Strings(scriptArray)
			RunTUI(selected, scriptArray)
		} else {
			fmt.Println("\n Nenhum script encontrado para esse projeto.")
		}
		return tea.Quit
	})

	if handled {
		return m, cmd
	}

	return m, nil
}

func (m listModel) View() string {
	return RenderList("Selecione um projeto:", m.choices, m.cursor, m.quitting)
}
