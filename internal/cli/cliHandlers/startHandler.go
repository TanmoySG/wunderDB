package cliHandlers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/shutdown"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/startup"
	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"
)

var StartOptFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "override",
		Aliases: []string{"o"},
		Usage:   "override configurations",
	},
	&cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Usage:   "set port to run instance",
	},
	&cli.StringFlag{
		Name:    "storage",
		Aliases: []string{"s"},
		Usage:   "set persistant storage path",
	},
}

func (ch cliHandler) StartOptHandler(ctx *cli.Context) error {
	overrideConfigFlag := ctx.Bool("override")
	if overrideConfigFlag {
		portOverride := ctx.String("port")
		persistantStoragePathOverride := ctx.String("storage")
		setEnvs(portOverride, persistantStoragePathOverride, overrideConfigFlag)
	}

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).PaddingTop(1)

	fmt.Println(style.Render("Starting wunderDb..."))

	c, err := config.Load()
	if err != nil {
		return fmt.Errorf("error loading configurations: %s", err)
	}

	w, err := startup.Prepare(*c)
	if err != nil {
		return fmt.Errorf("error starting wdb server: %s", err)
	}

	shutdown.Listen(*w, *c) // listens to shutdown signals
	startup.Start(w, c)     // starts server and initial setup

	return nil
}

func setEnvs(port, persistantStoragePath string, override bool) {
	os.Setenv(config.PORT, port)
	os.Setenv(config.PERSISTANT_STORAGE_PATH, persistantStoragePath)
	os.Setenv(config.OVERRIDE_CONFIG, strconv.FormatBool(override))
}
