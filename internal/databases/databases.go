package databases

import (
	"fmt"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

type Namespace model.Namespace

func UseNamespace(namespace model.Namespace) Namespace {
	return Namespace(namespace)
}

func (n Namespace) CheckIfDatabaseExists(databaseID model.Identifier) bool {
	_, dbExists := n.Databases[databaseID]
	if dbExists {
		return dbExists
	} else {
		return dbDoesNotExist
	}
}

func (n Namespace) CreateNewDatabase(databaseID model.Identifier, metadata model.Metadata, access model.Access) error {
	if n.CheckIfDatabaseExists(databaseID) {
		return fmt.Errorf(NamespaceErrorFormat, er.NamespaceAlreadyExistsError.ErrCode, "error creating namespace", er.NamespaceAlreadyExistsError.ErrMessage)
	}
	n.Databases[databaseID] = &model.Database{
		Collections: map[model.Identifier]*model.Collection{},
		Metadata:    metadata,
		Access:      map[model.Identifier]*model.Access{},
	}
	return nil
}

func (n Namespace) DeleteDatabase(databaseID model.Identifier) error {
	if n.CheckIfDatabaseExists(databaseID) {
		delete(n.Databases, databaseID)
		return nil
	}
	return fmt.Errorf(NamespaceErrorFormat, er.NamespaceDoesNotExistsError.ErrCode, "error deleting namespace", er.NamespaceDoesNotExistsError.ErrMessage)
}
