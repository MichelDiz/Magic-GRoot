package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func RenderList(title string, choices []string, cursor int, quitting bool) string {
	s := fmt.Sprintf("\n%s\n\n", title) // Título da lista

	for i, choice := range choices {
		cursorMarker := "  "
		if cursor == i {
			cursorMarker = "=>"
		}
		s += fmt.Sprintf("%s %s\n", cursorMarker, choice)
	}

	if quitting {
		return ""
	}

	s += "\n\nPressione ↑/↓, Enter para selecionar, Q para sair."
	return s
}

func UpdateListHandler(
	msg tea.Msg,
	cursor *int,
	choices []string,
	quitting *bool,
	inputMode *bool,
	onSelect func(selected string) tea.Cmd,
) (tea.Cmd, bool) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if inputMode != nil && *inputMode {
			if msg.String() == "enter" {
				*inputMode = false
				return tea.Quit, true
			}
			return nil, true
		}

		switch msg.String() {
		case "q", "esc":
			*quitting = true
			return tea.Quit, true
		case "ctrl+c":
			*quitting = true
			return tea.Interrupt, true
		case "up":
			if *cursor > 0 {
				*cursor--
			}
			return nil, true
		case "down":
			if *cursor < len(choices)-1 {
				*cursor++
			}
			return nil, true
		case "enter":
			if onSelect != nil {
				return onSelect(choices[*cursor]), true
			}

			return tea.Quit, true
		}
	}
	return nil, false
}
