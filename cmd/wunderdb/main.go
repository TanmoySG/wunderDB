package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/fsLoader"
	"github.com/TanmoySG/wunderDB/internal/roles"
	s "github.com/TanmoySG/wunderDB/internal/server"
	"github.com/TanmoySG/wunderDB/model"
	wdbClient "github.com/TanmoySG/wunderDB/pkg/wdb"
	log "github.com/sirupsen/logrus"
)

func main() {
	fs := fsLoader.NewWFileSystem("wfs")

	loadedDatabase, _ := fs.LoadDatabases()
	loadedRoles, _ := fs.LoadRoles()

	db := databases.Use(loadedDatabase)
	rl := roles.Use(loadedRoles)

	wdbc := wdbClient.NewWdbClient(db, rl)

	Shutdown(db, rl)

	server := s.NewWdbServer(wdbc)

	server.Start()
}

// move clean start and exit to a different file

func cleanExit(db map[model.Identifier]*model.Database, rl map[model.Identifier]*model.Role) error {
	fs := fsLoader.NewWFileSystem("wfs")

	err := fs.UnloadDatabases(db)
	if err != nil {
		return fmt.Errorf("error in graceful shutdown: %s", err)
	}

	err = fs.UnloadRoles(rl)
	if err != nil {
		return fmt.Errorf("error in graceful shutdown: %s", err)
	}

	return nil
}

func Shutdown(db map[model.Identifier]*model.Database, rl map[model.Identifier]*model.Role) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Infof("Gracefully shutting down...")
		err := cleanExit(db, rl)
		if err != nil {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}()
}
