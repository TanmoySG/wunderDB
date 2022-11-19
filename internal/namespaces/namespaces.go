package namespaces

import (
	"fmt"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

type Namespaces map[model.Identifier]*model.Namespace

func (ns Namespaces) CheckIfNamespaceExists(namespaceID model.Identifier) bool {
	_, nsExists := ns[namespaceID]
	if nsExists {
		return nsExistsInNamespaces
	} else {
		return nsDoesNotExistInNamespaces
	}
}

func (ns Namespaces) CreateNewNamespace(namespaceID model.Identifier, metadata model.Metadata, access model.Access) error {
	if ns.CheckIfNamespaceExists(namespaceID) {
		return fmt.Errorf(NamespaceErrorFormat, er.NamespaceAlreadyExistsError.ErrCode, "error creating namespace", er.NamespaceAlreadyExistsError.ErrMessage)
	}
	ns[namespaceID] = &model.Namespace{
		Databases: map[model.Identifier]*model.Database{},
		Metadata:  metadata,
		Access:    map[model.Identifier]*model.Access{},
	}
	return nil
}

func (ns Namespaces) DeleteNamespace(namespaceID model.Identifier) error {
	if ns.CheckIfNamespaceExists(namespaceID) {
		delete(ns, namespaceID)
		return nil
	}
	return fmt.Errorf(NamespaceErrorFormat, er.NamespaceDoesNotExistsError.ErrCode, "error deleting namespace", er.NamespaceDoesNotExistsError.ErrMessage)
}

func (ns Namespaces) ModifyNamespaceMetadata(namespaceID model.Identifier) {
	// future scoped
}
