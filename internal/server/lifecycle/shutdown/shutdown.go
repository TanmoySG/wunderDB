package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TanmoySG/wunderDB/internal/config"
	fsLoader "github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
	log "github.com/sirupsen/logrus"
)

func Listen(w model.WDB, c config.Config) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		log.Infof("Gracefully shutting down...")
		err := cleanExit(c.PersistantStoragePath, w.Databases, w.Roles, w.Users)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}()
}

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
