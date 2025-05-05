package main

import (
	c "github.com/TanmoySG/wunderDB/internal/cli"
	"github.com/TanmoySG/wunderDB/internal/version"
)

// ALWAYS UPDATE WHILE PUBLISHING NEW VERSION
const WDB_VERSION = version.WDB_VERSION
const CLI_VERSION = version.CLI_VERSION

func main() {
	//  go run ./cmd/wdbctl/cli.go start -p 8089 -o
	cli := c.CreateCLI(WDB_VERSION, CLI_VERSION)
	cli.Run()
}
