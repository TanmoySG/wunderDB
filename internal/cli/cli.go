package cli

import (
	"os"

	"github.com/TanmoySG/wunderDB/internal/cli/cliHandlers"
	"github.com/urfave/cli/v2"
)

func Run() {
	app := cli.NewApp()
	app.Name = "wdbCLI"
	app.Usage = "Command Line Interface for wdb"

	app.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "starts the wdb instance",
			Flags:  cliHandlers.StartOptFlags,
			Action: cliHandlers.StartOptHandler,
		},
		{
			Name:   "version",
			Usage:  "version of CLI and wunderDb",
			Action: cliHandlers.VersionOptHandler,
		},
	}

	_ = app.Run(os.Args)

}
