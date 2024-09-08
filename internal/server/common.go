package server

import (
	"fmt"
	"strings"

	"github.com/TanmoySG/wunderDB/internal/version"
	"github.com/charmbracelet/lipgloss"
)

var (
	defaultPanicMessage = "wunderDB panicked on request"

	// startup dialog configuration
	width          = 96
	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
	subtle   = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	docStyle = lipgloss.NewStyle().Padding(1, 0, 1, 0)
)

func (ws wdbServer) startupMessage(port string) {
	dialogBuilder := strings.Builder{}
	port = strings.Trim(port, ":")

	lineOne := lipgloss.NewStyle().
		Width(50).Bold(true).
		Align(lipgloss.Center).
		Render(fmt.Sprintf("wunderDb %s", version.WDB_VERSION))

	lineTwo := lipgloss.NewStyle().
		Width(50).Align(lipgloss.Center).
		Foreground(lipgloss.Color("#F25D94")).
		Render(fmt.Sprintf("Running on Port %s", port))

	sectionTwo := lipgloss.JoinHorizontal(lipgloss.Top, lineTwo)
	ui := lipgloss.JoinVertical(lipgloss.Center, lineOne, sectionTwo)

	dialog := lipgloss.Place(width, 10,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	dialogBuilder.WriteString(dialog + "\n")
	fmt.Println(docStyle.Render(dialogBuilder.String()))
}
