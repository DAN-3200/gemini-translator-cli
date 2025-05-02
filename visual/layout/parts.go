package layout

import (
	css "github.com/charmbracelet/lipgloss"
)

// -- Styles and Components -----------------------------------------------------
const (
	blue      = "#87CEEB" // #87CEEB
	green     = "#50C878" // #50C878
	yellow    = "#FADA5E" // #FADA5E
	white     = "#FAF9F6" // #FAF9F6
	steelBlue = "#4682B4" // #4682B4
)
var infoHelp = css.NewStyle().
	Foreground(css.Color("#606060")).
	Render("[Ctrl+x] Clear • [Ctrl+a] Swicth mode " + "• [Ctrl+d] Switch language " + "• [Esc] Exit")

var box = css.NewStyle().Padding(1)

var prefixStyle = func(text string) string {
	return css.NewStyle().
		Foreground(css.Color("#606060")).
		Render(text)
}

var titleApp = css.NewStyle().
	Bold(true).
	Padding(0, 2).
	Foreground(css.Color(white)).
	Background(css.Color(steelBlue)).
	Render("LoroApp")

var fontColor = func(text string, color string) string {
	return css.NewStyle().
		Foreground(css.Color(color)).
		Render(text)
}
