package tui

import (
	"termlock/internal/styles"
	// "github.com/charmbracelet/lipgloss"
)

func RenderFooter(mode string) string {
	switch mode {
	case "search":
		return 	styles.Bright.Render("↑/↓") + " " +
				styles.Faded.Render("Move") + "   " +
				styles.Bright.Render("Enter") + " " +
				styles.Faded.Render("Copy") + "   " +
				styles.Bright.Render("Esc") + " " +
				styles.Faded.Render("Cancel")
	default:
		return 	styles.Bright.Render("↑/↓") + " " +
				styles.Faded.Render("Move") + "   " +
				styles.Bright.Render("Enter") + " " +
				styles.Faded.Render("Copy") + "   " +
				styles.Bright.Render("/") + " " +
				styles.Faded.Render("Search") + "   " +
				styles.Bright.Render("i") + " " +
				styles.Faded.Render("Import") + "   " +
				styles.Bright.Render("q") + " " +
				styles.Faded.Render("Quit")
	}
}
