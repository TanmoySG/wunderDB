package cliHandlers

import (
	"github.com/urfave/cli/v2"
)

type cliHandler struct {
	wdbVersion string
	cliVersion string
}

type Client interface {
	VersionOptHandler(ctx *cli.Context) error
	StartOptHandler(ctx *cli.Context) error
}

func GetCliHandlers(wdbVersion, cliVersion string) Client {
	return &cliHandler{
		wdbVersion: wdbVersion,
		cliVersion: cliVersion,
	}
}
