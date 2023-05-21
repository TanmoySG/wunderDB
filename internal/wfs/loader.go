package wfs

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/fs"
	"github.com/TanmoySG/wunderDB/pkg/tools"
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

	users := map[model.Identifier]*model.User{}

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

	roles := map[model.Identifier]*model.Role{}

	if fs.CheckFileExists(w.rolesBasePath) {
		persitedRolesBytes, err := os.ReadFile(w.rolesBasePath)
		if err != nil {
			return nil, fmt.Errorf("error reading persisted persisted file: %s", err)
		}

		updateRolesAspectFlag := os.Getenv("UPDATE_ROLE_ASPECTS")
		if len(updateRolesAspectFlag) != 0 {
			updateRoles, err := strconv.ParseBool(updateRolesAspectFlag)
			if err != nil {
				return nil, fmt.Errorf("failed to parse flag UPDATE_ROLE_ASPECTS: %s", err)
			}

			if updateRoles {
				updateTool := os.Getenv("UPDATE_TOOL")
				updateValue := os.Getenv("UPDATE_VALUE")

				persitedRolesBytes, err = updateRoleAspect(updateTool, persitedRolesBytes, updateValue)
				if err != nil {
					return nil, fmt.Errorf("failed to update roles aspect: %s", err)
				}

				_ = os.Unsetenv("UPDATE_ROLE_ASPECTS")
				_ = os.Unsetenv("UPDATE_TOOL")
				_ = os.Unsetenv("UPDATE_VALUE")
			}
		}

		err = json.Unmarshal(persitedRolesBytes, &roles)
		if err != nil {
			return nil, fmt.Errorf("error marshaling roles file: %s", err)
		}
	}

	return roles, nil
}

func loadEntity(entityPath string, entity any) (any, error) {
	if !fs.CheckFileExists(entityPath) {
		return nil, fmt.Errorf("%s does not exist", entityPath)
	}

	persitedRolesBytes, err := os.ReadFile(entityPath)
	if err != nil {
		return nil, fmt.Errorf("error reading persisted persisted file: %s", err)
	}

	err = json.Unmarshal(persitedRolesBytes, &entity)
	if err != nil {
		return nil, fmt.Errorf("error marshaling roles file: %s", err)
	}

	return entity, nil
}

func updateRoleAspect(toolToUse string, rolesFileReadBytes []byte, valueToSet string) ([]byte, error) {
	t, err := tools.Use(toolToUse)
	if err != nil {
		return nil, err
	}

	out, err := t.Execute(string(rolesFileReadBytes), valueToSet)
	if err != nil {
		return nil, err
	}

	return out.([]byte), nil
}
