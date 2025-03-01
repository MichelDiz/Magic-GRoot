package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestListModel(t *testing.T) {
	projects := []string{"project1", "project2"}
	model := NewListModel(projects)

	// Test initial state
	if model.cursor != 0 {
		t.Errorf("Expected initial cursor to be 0, got %d", model.cursor)
	}

	// Test cursor movement
	msg := tea.KeyMsg{Type: tea.KeyUp}
	newModel, _ := model.Update(msg)
	m := newModel.(listModel)
	if m.cursor != 0 {
		t.Errorf("Expected cursor to stay at 0 when already at top")
	}

	msg = tea.KeyMsg{Type: tea.KeyDown}
	newModel, _ = m.Update(msg)
	m = newModel.(listModel)
	if m.cursor != 1 {
		t.Errorf("Expected cursor to move to 1")
	}
}
