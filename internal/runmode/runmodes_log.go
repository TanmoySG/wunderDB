package runmode

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	title    strings.Builder
	subtle   = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	docStyle = lipgloss.NewStyle().Padding(1, 0, 1, 0)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			Width(50).
			BorderForeground(subtle)
	descStyle = lipgloss.NewStyle().MarginTop(1)
	divider   = lipgloss.NewStyle().
			SetString("•").
			Padding(0, 1).
			Foreground(subtle).
			String()
	special = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	url     = lipgloss.NewStyle().Foreground(special).Render
)

func runModeStartupMessage(message, port string) {
	port = strings.Trim(port, ":")
	doc := strings.Builder{}

	desc := lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("wunderDB is running in MAINTENANCE_MODE"),
		infoStyle.Render("Running on Port "+url(port)),
	)

	row := lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
	doc.WriteString(row + "\n\n")
	fmt.Println(docStyle.Render(doc.String()))
}
