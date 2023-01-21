package main

import c "github.com/TanmoySG/wunderDB/internal/cli"

// ALWAYS UPDATE WHILE PUBLISHING NEW VERSION
const WDB_VERSION = "v1.0.3-test"
const CLI_VERSION = "v0.0.1-test"

func main() {
	//  go run ./cmd/wdbcli/cli.go -start -p 8089 -o
	cli := c.CreateCLI(WDB_VERSION, CLI_VERSION)
	cli.Run()
}
