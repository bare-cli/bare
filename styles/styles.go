package styles

import "github.com/charmbracelet/lipgloss"

// var style = lipgloss.NewStyle().Bold(true)

var InitStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4"))

var InitError = lipgloss.NewStyle().	
	Foreground(lipgloss.Color("#fe4343"))

var InitSuccess = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#30da00"))

var AddFileStlyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#b600da"))