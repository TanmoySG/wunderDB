package cli

import (
	"flag"
	"os"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/lifecycle/shutdown"
	"github.com/TanmoySG/wunderDB/internal/lifecycle/startup"
	log "github.com/sirupsen/logrus"
)

func Listen() {
	start := flag.Bool("start", false, "Starts wDB instance")
	port := flag.String("p", "8086", "Port of wDB instance")
	persistantStoragePath := flag.String("s", "", "Persistant FS of wDB instance")
	flag.Parse()

	if *start {
		setEnvs(*port, *persistantStoragePath)
		startWdb()
	}
}

func startWdb() {
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

func setEnvs(port, persistantStoragePath string) {
	os.Setenv(config.PORT, port)
	os.Setenv(config.PERSISTANT_STORAGE_PATH, persistantStoragePath)
}
