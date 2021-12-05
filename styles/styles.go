package styles

import "github.com/charmbracelet/lipgloss"

// var style = lipgloss.NewStyle().Bold(true)

var InitStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

var Error = lipgloss.NewStyle().
	Foreground(lipgloss.Color("1"))

var Success = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#30da00"))

var Warning = lipgloss.NewStyle().Foreground(lipgloss.Color("13"))

var PromptStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00BDFF"))

var StatusSuccess = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#74FF33"))

var StatusPrompt = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#00BDFF"))

var StatusError = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#FF3A3A"))

var Description = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#969696"))
