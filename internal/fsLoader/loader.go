package fsLoader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/fs"
)

func (w WFileSystem) LoadNamespaces() (map[model.Identifier]*model.Namespace, error) {

	var namespaces map[model.Identifier]*model.Namespace

	if fs.CheckFileExists(w.namespacesBasePath) {
		persitedNamespacesBytes, err := os.ReadFile(w.namespacesBasePath)
		if err != nil {
			return nil, fmt.Errorf("error reading namespace file: %s", err)
		}

		err = json.Unmarshal(persitedNamespacesBytes, &namespaces)
		if err != nil {
			return nil, fmt.Errorf("error marshaling namespace file: %s", err)
		}
	}

	return namespaces, nil
}

func (w WFileSystem) LoadDatabases() (map[model.Identifier]*model.Database, error) {

	var databases map[model.Identifier]*model.Database

	if fs.CheckFileExists(w.databasesBasePath) {
		persitedDatabasesBytes, err := os.ReadFile(w.databasesBasePath)
		if err != nil {
			return nil, fmt.Errorf("error reading database file: %s", err)
		}

		err = json.Unmarshal(persitedDatabasesBytes, &databases)
		if err != nil {
			return nil, fmt.Errorf("error marshaling database file: %s", err)
		}
	}

	return databases, nil
}

func (w WFileSystem) LoadUsers() (map[model.Identifier]*model.User, error) {

	var users map[model.Identifier]*model.User

	if fs.CheckFileExists(w.usersBasePath) {
		persitedUsersBytes, err := os.ReadFile(w.usersBasePath)
		if err != nil {
			return nil, fmt.Errorf("error reading users persisted file: %s", err)
		}

		err = json.Unmarshal(persitedUsersBytes, &users)
		if err != nil {
			return nil, fmt.Errorf("error marshaling users file: %s", err)
		}
	}

	return users, nil
}

func (w WFileSystem) LoadRoles() (map[model.Identifier]*model.Role, error) {

	var roles map[model.Identifier]*model.Role

	if fs.CheckFileExists(w.rolesBasePath) {
		persitedRolesBytes, err := os.ReadFile(w.rolesBasePath)
		if err != nil {
			return nil, fmt.Errorf("error reading persisted persisted file: %s", err)
		}

		err = json.Unmarshal(persitedRolesBytes, &roles)
		if err != nil {
			return nil, fmt.Errorf("error marshaling roles file: %s", err)
		}
	}

	return roles, nil
}
