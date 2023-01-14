package main

import (
	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/fsLoader"
	"github.com/TanmoySG/wunderDB/internal/roles"
	s "github.com/TanmoySG/wunderDB/internal/server"
	"github.com/TanmoySG/wunderDB/internal/users"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
	wdbClient "github.com/TanmoySG/wunderDB/pkg/wdb"
	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configurations: %s", err)
	}
	fs := fsLoader.NewWFileSystem("wfs")

	loadedDatabase, _ := fs.LoadDatabases()
	loadedRoles, _ := fs.LoadRoles()
	loadedUsers, _ := fs.LoadUsers()

	db := databases.Use(loadedDatabase)
	rl := roles.Use(loadedRoles)
	us := users.Use(loadedUsers)

	halg := authentication.MD5

	wdbc := wdbClient.NewWdbClient(model.Configurations{}, db, rl, us, halg)

	wdbc.InitializeAdmin(c)

	Shutdown(db, rl, us)

	server := s.NewWdbServer(wdbc, c.Port)

	server.Start()
}
