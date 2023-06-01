package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/collections"
	"github.com/TanmoySG/wunderDB/internal/roles/sysroles"
	"github.com/TanmoySG/wunderDB/model"
	wdbErrors "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

// updates parent entity metadata
func (wdb wdbClient) updateParentMetadata(databaseId, collectionId *model.Identifier) {
	if databaseId != nil {
		wdb.Databases.UpdateMetadata(*databaseId)
		if collectionId != nil {
			_, database := wdb.Databases.CheckIfExists(*databaseId)
			collections := collections.UseDatabase(database)
			collections.UpdateMetadata(*collectionId)
		}
	}
}

func (wdb wdbClient) grantSystemDefaultRole(userId model.Identifier, role sysroles.SystemDefaultRole, args ...string) *wdbErrors.WdbError {
	if exists, _ := wdb.Roles.CheckIfExists(model.Identifier(role.RoleID)); !exists {
		err := wdb.Roles.CreateRole(model.Identifier(role.RoleID), role.Privileges, []string{}, role.Hidden)
		if err != nil {
			return err
		}
	}

	var onEntities model.Entities
	switch len(args) {
	case 1:
		onEntities.Databases = &args[0]
	case 2:
		onEntities.Databases = &args[0]
		onEntities.Collections = &args[1]
	case 3:
		onEntities.Databases = &args[0]
		onEntities.Collections = &args[1]
		onEntities.Users = &args[2]
	default:
		// move to wdb-errors, after finalizing the error code, http code
		return &wdbErrors.WdbError{
			ErrCode:        "irregularEntitiesList",
			ErrMessage:     "error with number of entities in list",
			HttpStatusCode: 409,
		}
	}

	permission := model.Permissions{
		Role: model.Identifier(role.RoleID),
		On:   &onEntities,
	}

	wdb.Users.GrantRole(userId, permission)
	return nil
}
