package cliHandlers

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"
)

// ALWAYS UPDATE WHILE PUBLISHING NEW VERSION
const WDB_VERSION = "v2.0.0-test"
const CLI_VERSION = "v0.0.1-test"

func VersionOptHandler(ctx *cli.Context) error {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4"))

	fmt.Printf("wDB Version: %s\n", style.Render(WDB_VERSION))
	fmt.Printf("CLI Version: %s\n", style.Render(CLI_VERSION))

	return nil
}
