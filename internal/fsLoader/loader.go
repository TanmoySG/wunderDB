package fsLoader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/fs"
)

func (w WFileSystem) LoadNamespaces() (map[model.Identifier]*model.Namespace, error) {

	var namespaces map[model.Identifier]*model.Namespace

	if fs.CheckFileExists(w.namespacesBasePath) {
		persitedNamespacesBytes, err := ioutil.ReadFile(w.namespacesBasePath)
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
		persitedDatabasesBytes, err := ioutil.ReadFile(w.databasesBasePath)
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
