package startup

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/roles"
	"github.com/TanmoySG/wunderDB/internal/users"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
	wdbClient "github.com/TanmoySG/wunderDB/pkg/wdb"

	wdbs "github.com/TanmoySG/wunderDB/internal/server"
)

func Prepare(c config.Config) (*model.WDB, error) {
	fs := wfs.NewWFileSystem(c.PersistantStoragePath)

	err := fs.InitializeWFS()
	if err != nil {
		return nil, fmt.Errorf("error loading wfs: %s", err)
	}

	loadedDatabase, err := fs.LoadDatabases()
	if err != nil {
		return nil, fmt.Errorf("error loading wfs: %s", err)
	}

	loadedRoles, err := fs.LoadRoles()
	if err != nil {
		return nil, fmt.Errorf("error loading wfs: %s", err)
	}

	loadedUsers, err := fs.LoadUsers()
	if err != nil {
		return nil, fmt.Errorf("error loading wfs: %s", err)
	}

	wdb := model.WDB{
		Databases: databases.Use(loadedDatabase),
		Roles:     roles.Use(loadedRoles),
		Users:     users.Use(loadedUsers),
	}

	return &wdb, nil
}

func Start(w *model.WDB, c *config.Config) {
	wdbClientConfigurations := model.Configurations{
		Admin: (*model.Identifier)(&c.AdminID),
	}

	wdbc := wdbClient.NewWdbClient(wdbClientConfigurations, w.Databases, w.Roles, w.Users, authentication.MD5)
	wdbc.InitializeAdmin(c)

	server := wdbs.NewWdbServer(wdbc, c.Port, c.RootDirectoryPath)
	server.Start()
}
