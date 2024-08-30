package cliHandlers

import (
	"os"
	"strings"

	wdb "github.com/TanmoySG/wunderDB/cmd/entrypoint"
	"github.com/TanmoySG/wunderDB/internal/config"
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
	&cli.StringFlag{
		Name:    "admin",
		Aliases: []string{"a"},
		Usage:   "Admin Username and password. Value should be passed as username:password",
	},
}

func (ch cliHandler) StartOptHandler(ctx *cli.Context) error {
	overrideConfigFlag := ctx.Bool("override")
	if overrideConfigFlag {
		os.Setenv(config.OVERRIDE_CONFIG, "true")
	}

	setEnv(ctx)

	if err := wdb.EntryPoint(); err != nil {
		return err
	}

	return nil
}

func setEnv(ctx *cli.Context) {
	portOverride := ctx.String("port")
	if portOverride != "" {
		os.Setenv(config.PORT, portOverride)
	}

	persistantStoragePathOverride := ctx.String("storage")
	if persistantStoragePathOverride != "" {
		os.Setenv(config.PERSISTANT_STORAGE_PATH, persistantStoragePathOverride)
	}

	adminCredentials := ctx.String("admin")
	if adminCredentials != "" {
		adminCredentialSlice := strings.Split(adminCredentials, ":")
		adminUserID, adminPassword := adminCredentialSlice[0], adminCredentialSlice[1]
		if adminUserID != "" {
			os.Setenv(config.ADMIN_ID, adminUserID)
			if adminPassword != "" {
				os.Setenv(config.ADMIN_PASSWORD, adminPassword)
			}
		}
	}
}
