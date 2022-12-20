package fsLoader

import (
	"encoding/json"
	"os"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/fs"
)

func (w WFileSystem) UnloadNamespaces(namespaces map[model.Identifier]*model.Namespace) error {
	namespacesJson, err := json.Marshal(namespaces)
	if err != nil {
		return err
	}

	if !fs.CheckFileExists(w.namespacesBasePath) {
		os.Create(w.namespacesBasePath)
	}

	err = os.WriteFile(w.namespacesBasePath, namespacesJson, 0740)
	if err != nil {
		return err
	}
	return nil
}

func (w WFileSystem) UnloadDatabases(databases map[model.Identifier]*model.Database) error {
	databasesJson, err := json.Marshal(databases)
	if err != nil {
		return err
	}

	if !fs.CheckFileExists(w.databasesBasePath) {
		os.Create(w.databasesBasePath)
	}

	err = os.WriteFile(w.databasesBasePath, databasesJson, 0740)
	if err != nil {
		return err
	}
	return nil
}

func (w WFileSystem) UnloadRoles(roles map[model.Identifier]*model.Role) error {
	rolesJson, err := json.Marshal(roles)
	if err != nil {
		return err
	}

	if !fs.CheckFileExists(w.rolesBasePath) {
		os.Create(w.rolesBasePath)
	}

	err = os.WriteFile(w.rolesBasePath, rolesJson, 0740)
	if err != nil {
		return err
	}
	return nil
}

func (w WFileSystem) UnloadUsers(users map[model.Identifier]*model.User) error {
	usersJson, err := json.Marshal(users)
	if err != nil {
		return err
	}

	if !fs.CheckFileExists(w.usersBasePath) {
		os.Create(w.usersBasePath)
	}

	err = os.WriteFile(w.usersBasePath, usersJson, 0740)
	if err != nil {
		return err
	}
	return nil
}
