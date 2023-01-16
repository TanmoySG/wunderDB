package main

import (
	"fmt"
	"os"
	"os/signal"

	fsLoader "github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
	log "github.com/sirupsen/logrus"
)

// move clean start and exit to a different file

func cleanExit(persistantStoragePath string, db map[model.Identifier]*model.Database, rl map[model.Identifier]*model.Role, us map[model.Identifier]*model.User) error {
	fs := fsLoader.NewWFileSystem(persistantStoragePath)

	err := fs.UnloadDatabases(db)
	if err != nil {
		return fmt.Errorf("error in graceful shutdown: %s", err)
	}

	err = fs.UnloadRoles(rl)
	if err != nil {
		return fmt.Errorf("error in graceful shutdown: %s", err)
	}

	err = fs.UnloadUsers(us)
	if err != nil {
		return fmt.Errorf("error in graceful shutdown: %s", err)
	}

	return nil
}

func Shutdown(persistantStoragePath string, db map[model.Identifier]*model.Database, rl map[model.Identifier]*model.Role, us map[model.Identifier]*model.User) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Infof("Gracefully shutting down...")
		err := cleanExit(persistantStoragePath,db, rl, us)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}()
}
