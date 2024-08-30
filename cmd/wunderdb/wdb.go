package main

import (
	wdb "github.com/TanmoySG/wunderDB/cmd/entrypoint"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := wdb.EntryPoint(); err != nil {
		log.Fatalf("error starting wdb: %s", err)
	}
}
