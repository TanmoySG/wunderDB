package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/fsLoader"
	s "github.com/TanmoySG/wunderDB/internal/server"
	"github.com/TanmoySG/wunderDB/model"
	wdbClient "github.com/TanmoySG/wunderDB/pkg/wdb"
	log "github.com/sirupsen/logrus"
)

func main() {
	fs := fsLoader.NewWFileSystem("wfs")

	loadedDatabase, _ := fs.LoadDatabases()
	db := databases.WithWDB(loadedDatabase)
	wdbc := wdbClient.NewWdbClient(db)

	Shutdown(db)

	server := s.NewWdbServer(wdbc)

	server.Start()
}

// move clean start and exit to a different file

func cleanExit(db map[model.Identifier]*model.Database) error {
	fs := fsLoader.NewWFileSystem("wfs")

	err := fs.UnloadDatabases(db)
	if err != nil {
		return fmt.Errorf("error in graceful shutdown: %s", err)
	}

	return nil
}

func Shutdown(db map[model.Identifier]*model.Database) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Infof("Gracefully shutting down...")
		err := cleanExit(db)
		if err != nil {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}()
}
