package wfs

import (
	"fmt"
)

type WFileSystem struct {
	wfsBasePath        string
	namespacesBasePath string
}

const (
	namespacesBasePathFormat = "%s/namespaces/namespaces_persisted.json"
)

func NewWFileSystem(basePath string) WFileSystem {
	return WFileSystem{
		wfsBasePath:        basePath,
		namespacesBasePath: fmt.Sprintf(namespacesBasePathFormat, basePath),
	}
}
