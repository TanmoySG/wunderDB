package wfs

import (
	"fmt"
)

type WFileSystem struct {
	wfsBasePath        string
	namespacesBasePath string
	databasesBasePath  string
}

const (
	namespacesBasePathFormat = "%s/namespaces/namespaces_persisted.json"
	databasesBasePathFormat  = "%s/databases/databases_persisted.json"
)

func NewWFileSystem(basePath string) WFileSystem {
	return WFileSystem{
		wfsBasePath:        basePath,
		namespacesBasePath: fmt.Sprintf(namespacesBasePathFormat, basePath),
		databasesBasePath:  fmt.Sprintf(databasesBasePathFormat, basePath),
	}
}
