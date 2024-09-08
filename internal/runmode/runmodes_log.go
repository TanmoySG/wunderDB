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
	special1  = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	mode      = lipgloss.NewStyle().Foreground(special1).Render
	special2  = lipgloss.AdaptiveColor{Light: "#6979e3", Dark: "#2940d7"}
	portStyle = lipgloss.NewStyle().Foreground(special2).Render
)

func runModeStartupMessage(modeName, port string) {
	port = strings.Trim(port, ":")
	doc := strings.Builder{}

	desc := lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("wunderDB is running in "+mode(modeName)),
		infoStyle.Render("Running on Port "+portStyle(port)),
	)

	row := lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
	doc.WriteString(row + "\n\n")
	fmt.Println(docStyle.Render(doc.String()))
}
