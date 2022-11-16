package wfs

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
