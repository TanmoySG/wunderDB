package wfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/TanmoySG/wunderDB/model"
)

func (w WFileSystem) LoadNamespaces() (map[model.Identifier]*model.Namespace, error) {
	wfsNamespacesPath := w.namespacesBasePath

	namespaces := map[model.Identifier]*model.Namespace{}

	persitedNamespaceDirectories, err := ioutil.ReadDir(wfsNamespacesPath)
	if err != nil {
		return nil, fmt.Errorf("error reading namespace file: %s", err)
	}

	for _, persitedNamespace := range persitedNamespaceDirectories {
		if persitedNamespace.IsDir() {

			namespaceIdentifier := model.Identifier(persitedNamespace.Name())
			persitedNamespaceFilePath := w.getPersistedNamespaceFilePath(namespaceIdentifier.String())

			persitedNamespaceBytes, err := ioutil.ReadFile(persitedNamespaceFilePath)
			if err != nil {
				return nil, fmt.Errorf("error reading namespace file: %s", err)
			}

			var persitedNamespaceParsed model.Namespace

			err = json.Unmarshal(persitedNamespaceBytes, &persitedNamespaceParsed)
			if err != nil {
				return nil, fmt.Errorf("error reading namespace file: %s", err)
			}

			namespaces[namespaceIdentifier] = &persitedNamespaceParsed
		}
	}

	return namespaces, nil
}
