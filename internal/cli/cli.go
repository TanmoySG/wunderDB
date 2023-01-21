package cli

import (
	"log"
	"os"

	"github.com/TanmoySG/wunderDB/internal/cli/cliHandlers"
	"github.com/urfave/cli/v2"
)

type cliClient struct {
	cliHandlers cliHandlers.Client
}

type Client interface {
	Run()
}

func CreateCLI(wdbVersion, cliVersion string) Client {
	return cliClient{
		cliHandlers: cliHandlers.GetCliHandlers(wdbVersion, cliVersion),
	}
}

func (cc cliClient) Run() {
	app := cli.NewApp()
	app.Name = "wdbctl"
	app.Usage = "Command Line Interface for wdb"

	app.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "starts the wdb instance",
			Flags:  cliHandlers.StartOptFlags,
			Action: cc.cliHandlers.StartOptHandler,
		},
		{
			Name:   "version",
			Usage:  "version of CLI and wunderDb",
			Action: cc.cliHandlers.VersionOptHandler,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("CLI Action Failed : %s", err)
	}
}
