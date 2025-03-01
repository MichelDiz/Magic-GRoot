package utils

import "testing"

func TestNeedsRunPrefix(t *testing.T) {
	tests := []struct {
		name     string
		runner   string
		expected bool
	}{
		{"npm needs prefix", "npm", true},
		{"yarn no prefix", "yarn", false},
		{"pnpm needs prefix", "pnpm", true},
		{"bash no prefix", "bash", false},
		{"unknown needs prefix", "unknown", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NeedsRunPrefix(tt.runner); got != tt.expected {
				t.Errorf("NeedsRunPrefix(%s) = %v, want %v", tt.runner, got, tt.expected)
			}
		})
	}
}
