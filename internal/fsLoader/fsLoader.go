package fsLoader

import (
	"fmt"
)

type WFileSystem struct {
	wfsBasePath        string
	namespacesBasePath string
	databasesBasePath  string
	usersBasePath      string
	rolesBasePath      string
}

const (
	namespacesBasePathFormat = "%s/namespaces/namespaces_persisted.json"
	databasesBasePathFormat  = "%s/databases/databases_persisted.json"
	usersBasePathFormat      = "%s/users/users_persisted.json"
	rolesBasePathFormat      = "%s/roles/roles_persisted.json"
)

func NewWFileSystem(basePath string) WFileSystem {
	return WFileSystem{
		wfsBasePath:        basePath,
		namespacesBasePath: fmt.Sprintf(namespacesBasePathFormat, basePath),
		databasesBasePath:  fmt.Sprintf(databasesBasePathFormat, basePath),
		usersBasePath:      fmt.Sprintf(usersBasePathFormat, basePath),
		rolesBasePath:      fmt.Sprintf(rolesBasePathFormat, basePath),
	}
}
