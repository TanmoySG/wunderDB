package main

import c "github.com/TanmoySG/wunderDB/internal/cli"

// ALWAYS UPDATE WHILE PUBLISHING NEW VERSION
const WDB_VERSION = "v2.0.0-beta"
const CLI_VERSION = "v0.0.1-beta"

func main() {
	//  go run ./cmd/wdbctl/cli.go start -p 8089 -o
	cli := c.CreateCLI(WDB_VERSION, CLI_VERSION)
	cli.Run()
}
