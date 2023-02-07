package cliHandlers

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"
)

func (ch cliHandler) VersionOptHandler(ctx *cli.Context) error {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4"))

	fmt.Printf("wDB Version: %s\n", style.Render(ch.wdbVersion))
	fmt.Printf("CLI Version: %s\n", style.Render(ch.cliVersion))

	return nil
}
