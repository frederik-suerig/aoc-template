package ui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/manifoldco/promptui"
)

var (
	// Color definitions
	successColor  = lipgloss.Color("2") // Green
	errorColor    = lipgloss.Color("1") // Red
	infoColor     = lipgloss.Color("6") // Cyan
	warningColor  = lipgloss.Color("3") // Yellow
	fileColor     = lipgloss.Color("4") // Blue
	downloadColor = lipgloss.Color("5") // Magenta
	dimColor      = lipgloss.Color("8") // Gray

	// Base styles
	successStyle = lipgloss.NewStyle().Foreground(successColor).Bold(true)
	errorStyle   = lipgloss.NewStyle().Foreground(errorColor).Bold(true)
	infoStyle    = lipgloss.NewStyle().Foreground(infoColor).Bold(true)
	warningStyle = lipgloss.NewStyle().Foreground(warningColor).Bold(true)
	fileStyle    = lipgloss.NewStyle().Foreground(fileColor)
	dimStyle     = lipgloss.NewStyle().Foreground(dimColor)

	// Icons
	successIcon  = "✓"
	errorIcon    = "✗"
	infoIcon     = "ℹ"
	warningIcon  = "⚠"
	fileIcon     = "📄"
	dirIcon      = "📁"
	downloadIcon = "⬇"
)

// MakeRelative converts an absolute path to a relative path from the current working directory
func MakeRelative(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		// If we can't get working dir, return the path as-is
		return path
	}
	rel, err := filepath.Rel(wd, path)
	if err != nil {
		// If we can't make it relative, return the path as-is
		return path
	}
	// Prepend ./ to make it explicit it's a relative path
	return "./" + rel
}

// Success prints a success message with icon
func Success(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	icon := successStyle.Render(successIcon)
	text := successStyle.Render(message)
	_, _ = fmt.Fprintf(os.Stdout, "\n%s %s\n", icon, text)
}

// Error prints an error message with icon
func Error(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	icon := errorStyle.Render(errorIcon)
	text := errorStyle.Render(message)
	_, _ = fmt.Fprintf(os.Stderr, "%s %s\n", icon, text)
}

// Info prints an info message with icon
func Info(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	icon := infoStyle.Render(infoIcon)
	text := infoStyle.Render(message)
	_, _ = fmt.Fprintf(os.Stdout, "%s %s\n", icon, text)
}

// Warning prints a warning message with icon
func Warning(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	icon := warningStyle.Render(warningIcon)
	text := warningStyle.Render(message)
	_, _ = fmt.Fprintf(os.Stdout, "%s %s\n", icon, text)
}

// FileCreated prints a message for a created file
func FileCreated(path string) {
	icon := successStyle.Render(successIcon)
	fileIconStyled := lipgloss.NewStyle().Foreground(fileColor).Bold(true).Render(fileIcon)
	relPath := MakeRelative(path)
	pathStyled := fileStyle.Render(relPath)
	_, _ = fmt.Fprintf(os.Stdout, "  %s %s %s\n\n", icon, fileIconStyled, pathStyled)
}

// DirCreated prints a message for a created directory
func DirCreated(path string) {
	icon := successStyle.Render(successIcon)
	dirIconStyled := lipgloss.NewStyle().Foreground(fileColor).Bold(true).Render(dirIcon)
	relPath := MakeRelative(path)
	pathStyled := fileStyle.Render(relPath)
	_, _ = fmt.Fprintf(os.Stdout, "  %s %s %s\n\n", icon, dirIconStyled, pathStyled)
}

// Download prints a download message
func Download(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	icon := lipgloss.NewStyle().Foreground(downloadColor).Bold(true).Render(downloadIcon)
	text := lipgloss.NewStyle().Foreground(downloadColor).Render(message)
	_, _ = fmt.Fprintf(os.Stdout, "  %s %s\n\n", icon, text)
}

// DimText prints dimmed text
func DimText(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	styled := dimStyle.Render(message)
	_, _ = fmt.Fprintf(os.Stdout, "%s\n", styled)
}

// Header prints a styled header with a border
func Header(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")). // White
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("6")). // Cyan border
		Padding(0, 1).
		Margin(1, 0)

	styled := headerStyle.Render(message)
	_, _ = fmt.Fprintf(os.Stdout, "\n%s\n", styled)
}

// HighlightInfo prints a highly visible info message with extra styling
func HighlightInfo(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	icon := infoStyle.Render(infoIcon)
	// Make it more prominent with bold, brighter color, and spacing
	highlightStyle := lipgloss.NewStyle().
		Foreground(infoColor).
		Bold(true)
	text := highlightStyle.Render(message)
	_, _ = fmt.Fprintf(os.Stdout, "\n%s %s\n", icon, text)
}

// ConfirmOverwrite prompts the user to confirm if they want to delete and recreate existing files
func ConfirmOverwrite(year, day int, relPath string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Day %d of %d already exists at %s. Delete and recreate? (y/N)", day, year, relPath),
		IsConfirm: true,
		Default:   "N",
	}

	result, err := prompt.Run()
	if err != nil {
		// User cancelled or said no
		if err == promptui.ErrInterrupt || err == promptui.ErrAbort {
			return false, nil
		}
		return false, err
	}

	// Add spacing after prompt response
	_, _ = fmt.Fprintf(os.Stdout, "\n")

	return result == "y" || result == "Y" || result == "yes" || result == "Yes", nil
}
