package cli

import (
	"flag"
	"os"
	"strconv"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/shutdown"
	"github.com/TanmoySG/wunderDB/internal/server/lifecycle/startup"
	log "github.com/sirupsen/logrus"
)

func Listen() {
	start := flag.Bool("start", false, "Starts wDB instance")
	port := flag.String("p", "8086", "Port of wDB instance")
	persistantStoragePath := flag.String("s", "", "Persistant FS of wDB instance")
	override := flag.Bool("o", false, "Override Configurations for wDB instance")

	flag.Parse()

	if *start {
		setEnvs(*port, *persistantStoragePath, *override)
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

func setEnvs(port, persistantStoragePath string, override bool) {
	os.Setenv(config.PORT, port)
	os.Setenv(config.PERSISTANT_STORAGE_PATH, persistantStoragePath)
	os.Setenv(config.OVERRIDE_CONFIG, strconv.FormatBool(override))
}
