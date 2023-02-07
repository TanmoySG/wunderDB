package namespaces

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

type Namespaces map[model.Identifier]*model.Namespace

func WithWDB(namespace Namespaces) Namespaces {
	return namespace
}

func (ns Namespaces) CheckIfNamespaceExists(namespaceID model.Identifier) bool {
	_, nsExists := ns[namespaceID]
	if nsExists {
		return namespaceExists
	} else {
		return namespaceDoesNotExist
	}
}

func (ns Namespaces) CreateNamespace(namespaceID model.Identifier, metadata model.Metadata, access model.Access) error {
	if ns.CheckIfNamespaceExists(namespaceID) {
		return fmt.Errorf(NamespaceErrorFormat, er.NamespaceAlreadyExistsError.ErrCode, "error creating namespace", er.NamespaceAlreadyExistsError.ErrMessage)
	}
	ns[namespaceID] = &model.Namespace{
		Databases: []model.Identifier{},
		Metadata:  metadata,
		Access:    map[model.Identifier]*model.Access{},
	}
	return nil
}

func (ns Namespaces) GetNamespace(namespaceID model.Identifier) (*model.Namespace, error) {
	if !ns.CheckIfNamespaceExists(namespaceID) {
		return nil, fmt.Errorf(NamespaceErrorFormat, er.NamespaceAlreadyExistsError.ErrCode, "error creating namespace", er.NamespaceAlreadyExistsError.ErrMessage)
	}
	return ns[namespaceID], nil
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
