package tui

import (
	"fmt"
	"mgr/internal/config"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type setRootModel struct {
	choices     []string
	cursor      int
	selected    string
	manualInput string
	inputMode   bool
}

func NewSetRootModel() setRootModel {
	return setRootModel{
		choices: []string{
			"Usar diretório atual",
			"Inserir diretório manualmente",
			"Cancelar",
		},
		cursor: 0,
	}
}

func (m setRootModel) Init() tea.Cmd {
	return nil
}

func (m setRootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputMode {

			if msg.String() == "enter" {
				if m.manualInput != "" {
					m.selected = m.manualInput
					config.SetConfig("root_path", m.selected)
					fmt.Println("\n Diretório root definido como:", m.selected)
				}
				return m, tea.Quit
			} else if msg.String() == "backspace" && len(m.manualInput) > 0 {
				m.manualInput = m.manualInput[:len(m.manualInput)-1]
			} else if len(msg.String()) == 1 {
				m.manualInput += msg.String()
			}
			return m, nil
		}

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
			switch m.cursor {
			case 0:
				currentDir, _ := os.Getwd()
				config.SetConfig("root_path", currentDir)
				fmt.Println("\n Diretório root definido como:", currentDir)
				return m, tea.Quit
			case 1:
				m.inputMode = true
			case 2:
				fmt.Println("\nOperação cancelada.")
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m setRootModel) View() string {
	if m.inputMode {
		return fmt.Sprintf("\n📝 Digite o caminho manualmente:\n%s\n(Pressione Enter para confirmar)", m.manualInput)
	}

	s := "\n Selecione o diretório root:\n\n"
	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "=>"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nUse as setas para navegar, Enter para selecionar, Q para sair."
	return s
}

func RunSetRootTUI() {
	p := tea.NewProgram(NewSetRootModel())
	if err := p.Start(); err != nil {
		fmt.Println("Erro ao iniciar a interface:", err)
		os.Exit(1)
	}
}
