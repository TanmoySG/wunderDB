package namespaces

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/model"
)

type Namespaces map[model.Identifier]*model.Namespace

func (ns Namespaces) CreateNewNamespace(namespaceID model.Identifier, metadata model.Metadata, access model.Access) error {
	if ns.isNamespaceExists(namespaceID) {
		return fmt.Errorf(NamespaceErrorFormat, NamespaceAlreadyExistsError.errCode, "error creating namespace", NamespaceAlreadyExistsError.errMessage)
	}
	ns[namespaceID] = &model.Namespace{
		Databases: map[model.Identifier]*model.Database{},
		Metadata:  metadata,
		Access:    map[model.Identifier]*model.Access{},
	}
	return nil
}

func (ns Namespaces) DeleteNamespace(namespaceID model.Identifier) error {
	if ns.isNamespaceExists(namespaceID) {
		delete(ns, namespaceID)
		return nil
	}
	return fmt.Errorf(NamespaceErrorFormat, NamespaceDoesnotExistsError.errCode, "error deleting namespace", NamespaceDoesnotExistsError.errMessage)
}

func (ns Namespaces) ModifyNamespaceMetadata(namespaceID model.Identifier) {
	// future scoped
}
