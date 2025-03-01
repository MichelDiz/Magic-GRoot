package tui

import (
	"fmt"
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
