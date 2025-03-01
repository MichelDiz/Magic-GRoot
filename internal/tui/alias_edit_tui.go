package tui

import (
	"fmt"
	"mgr/internal/config"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type aliasManagerModel struct {
	aliases   map[string]string
	choices   []string
	cursor    int
	inputMode bool
	editAlias string
	newName   string
	quitting  bool
}

func NewAliasManagerModel() *aliasManagerModel {
	aliases := config.GetAllAliases()
	choices := make([]string, 0, len(aliases))

	for alias, path := range aliases {
		choices = append(choices, fmt.Sprintf("%s → %s", alias, path))
	}

	return &aliasManagerModel{
		aliases: aliases,
		choices: choices,
		cursor:  0,
	}
}

func (m aliasManagerModel) Init() tea.Cmd {
	return nil
}

func (m *aliasManagerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputMode {
			switch msg.String() {
			case "enter":
				clearScreen()
				if m.newName != "" {

					projectPath, exists := m.aliases[m.editAlias]
					if exists {

						config.UpdateAlias(m.newName, m.editAlias)

						delete(m.aliases, m.editAlias)
						m.aliases[m.newName] = projectPath

						m.recarregarLista()
					}

					fmt.Printf("\n Alias '%s' renomeado para: %s\n", m.editAlias, m.newName)
				} else {
					fmt.Println("\n Alias inválido. Operação cancelada.")
				}
				m.inputMode = false
				return m, nil
			case "backspace":
				if len(m.newName) > 0 {
					m.newName = m.newName[:len(m.newName)-1]
				}
			default:
				if len(msg.String()) == 1 {
					m.newName += msg.String()
				}
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
			selected := m.choices[m.cursor]
			m.editAlias = strings.Split(selected, " → ")[0]
			clearScreen()
			fmt.Printf("\n Digite um nome para o alias do projeto '%s': ", m.editAlias)
			m.inputMode = true
			m.newName = ""
		case "delete", "backspace", "d":
			if len(m.choices) > 0 {
				alias := strings.Split(m.choices[m.cursor], " → ")[0]
				delete(m.aliases, alias)
				config.DeleteAlias(alias)
				m.recarregarLista()
			}
		}
	}
	return m, nil
}
func (m aliasManagerModel) View() string {
	if m.inputMode {
		return fmt.Sprintf("\n\n Novo nome para '%s': %s\n\n(Pressione Enter para confirmar)", m.editAlias, m.newName)
	}
	return RenderList("Gerenciador de Aliases (Pressione Enter para editar, Delete para excluir):", m.choices, m.cursor, m.quitting)
}

func (m *aliasManagerModel) recarregarLista() {
	m.choices = make([]string, 0, len(m.aliases))
	for alias, path := range m.aliases {
		m.choices = append(m.choices, fmt.Sprintf("%s → %s", alias, path))
	}
	if m.cursor >= len(m.choices) {
		m.cursor = len(m.choices) - 1
	}
}
