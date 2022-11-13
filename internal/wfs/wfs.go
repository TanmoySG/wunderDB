package wfs

import (
	"fmt"
)

type WFileSystem struct {
	wfsBasePath        string
	namespacesBasePath string
}

const (
	wfsNamespacesPathFormat = "%s/namespaces/"
	namespaceFilePathFormat = "%s/%s/%s_persisted.json"
)

func NewWFileSystem(basePath string) WFileSystem {
	return WFileSystem{
		wfsBasePath:        basePath,
		namespacesBasePath: fmt.Sprintf(wfsNamespacesPathFormat, basePath),
	}
}

func (w WFileSystem) getPersistedNamespaceFilePath(identifier string) string {
	return fmt.Sprintf(namespaceFilePathFormat, w.namespacesBasePath, identifier, identifier)
}
