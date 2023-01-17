package main

import (
	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/lifecycle/shutdown"
	"github.com/TanmoySG/wunderDB/internal/lifecycle/startup"
	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatalf("error loading configurations: %s", err)
	}

	w, err := startup.Prepare(*c)
	if err != nil {
		log.Fatalf("error starting wdb server: %s", err)
	}

	shutdown.Listen(*w, *c) // listens to shutdown signals
	startup.Start(w, c)     // starts server and initial setup
}
