package namespaces

import "github.com/TanmoySG/wunderDB/model"

const (
	nsExistsInNamespaces       = true
	nsDoesNotExistInNamespaces = false
)

func (ns Namespaces) isNamespaceExists(namespaceID model.Identifier) bool {
	_, nsExists := ns[namespaceID]
	if nsExists {
		return nsExistsInNamespaces
	} else {
		return nsDoesNotExistInNamespaces
	}
}
