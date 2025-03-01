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
	cmd, handled := UpdateListHandler(msg, &m.cursor, m.choices, &m.quitting, &m.inputMode, func(selected string) tea.Cmd {
		alias := strings.Split(selected, " → ")[0]
		m.editAlias = alias
		m.inputMode = true
		m.newName = ""
		clearScreen()
		fmt.Printf("\n Digite um novo nome para o alias '%s': ", alias)
		return nil
	})

	if handled {
		return m, cmd
	}

	if m.inputMode {
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			switch keyMsg.String() {
			case "enter":
				if m.newName != "" {
					projectPath, exists := m.aliases[m.editAlias]
					if exists {
						config.SetAlias(m.newName, projectPath)
						delete(m.aliases, m.editAlias)
						m.aliases[m.newName] = projectPath
						m.recarregarLista()
					}
				}
				m.inputMode = false
				m.editAlias = ""
				return m, nil
			case "backspace":
				if len(m.newName) > 0 {
					m.newName = m.newName[:len(m.newName)-1]
				}
			default:
				if len(keyMsg.String()) == 1 {
					m.newName += keyMsg.String()
				}
			}
			return m, nil
		}
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {

		if len(m.choices) > 0 {
			switch keyMsg.String() {
			case "delete", "backspace", "d":
				alias := strings.Split(m.choices[m.cursor], " → ")[0]
				delete(m.aliases, alias)
				config.DeleteAlias(alias)
				m.recarregarLista()
				return m, nil
			}
		}
	}

	return m, nil
}

func (m aliasManagerModel) View() string {
	if m.inputMode {
		return fmt.Sprintf("\n Novo nome para '%s': %s\n(Pressione Enter para confirmar)", m.editAlias, m.newName)
	}
	return RenderList("Gerenciador de Aliases (Pressione Enter para editar, Delete para excluir):", m.choices, m.cursor)
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
