package tui

import (
	"fmt"
	"mgr/internal/config"

	tea "github.com/charmbracelet/bubbletea"
)

type aliasModel struct {
	choices   []string
	cursor    int
	selected  string
	inputMode bool
	aliasName string
}

func NewAliasModel(projects []string) aliasModel {
	return aliasModel{
		choices: projects,
		cursor:  0,
	}
}

func (m aliasModel) Init() tea.Cmd {
	return nil
}

func (m aliasModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputMode {

			if msg.String() == "enter" {
				if m.aliasName != "" {
					config.SetAlias(m.aliasName, m.selected)
					fmt.Printf("\n Alias '%s' criado para o projeto: %s\n", m.aliasName, m.selected)
				} else {
					fmt.Println("\n Alias inválido. Operação cancelada.")
				}
				return m, tea.Quit
			} else if msg.String() == "backspace" && len(m.aliasName) > 0 {
				m.aliasName = m.aliasName[:len(m.aliasName)-1]
			} else if len(msg.String()) == 1 {
				m.aliasName += msg.String()
			}
			return m, nil
		}

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
			m.selected = m.choices[m.cursor]
			clearScreen()
			fmt.Printf("\n Digite um nome para o alias do projeto '%s': ", m.selected)
			m.inputMode = true
		}
	}
	return m, nil
}

func (m aliasModel) View() string {
	if m.inputMode {
		return fmt.Sprintf("\n Nome do alias: %s\n(Pressione Enter para confirmar)", m.aliasName)
	}

	s := "\n Escolha um projeto para criar um alias:\n\n"
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
