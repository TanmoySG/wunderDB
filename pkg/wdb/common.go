package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/collections"
	"github.com/TanmoySG/wunderDB/model"
)

// updates parent entity metadata
func (wdb wdbClient) updateParentMetadata(databaseId, collectionId *model.Identifier) {
	if databaseId != nil {
		wdb.Databases.UpdateMetadata(*databaseId)
		if collectionId != nil {
			_, database := wdb.Databases.CheckIfExists(*databaseId)
			collections := collections.UseDatabase(*database)
			collections.UpdateMetadata(*collectionId)
		}
	}
}
