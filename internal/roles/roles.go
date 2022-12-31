package roles

import (
	"encoding/json"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	p "github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/utils/maps"
)

const (
	roleExists       = true
	roleDoesNotExist = false

	allowed = true
	denied  = false
)

type Roles map[model.Identifier]*model.Role

func Use(roles Roles) Roles {
	return roles
}

func (r Roles) CheckIfExists(roleID model.Identifier) (bool, *model.Role) {
	role, dbExists := r[roleID]
	if dbExists {
		return roleExists, role
	} else {
		return roleDoesNotExist, role
	}
}

func (r Roles) CreateRole(roleID model.Identifier, allowed []string, denied []string) *er.WdbError {

	grants, err := getPrivileges(allowed, denied)
	if err != nil {
		return &er.WdbError{
			ErrCode:        er.EncodeDecodeError.ErrCode,
			ErrMessage:     err.Error(),
			HttpStatusCode: er.EncodeDecodeError.HttpStatusCode,
		}
	}

	r[roleID] = &model.Role{
		RoleID: roleID,
		Grants: *grants,
	}
	return nil

}

func (r Roles) ListRoles() Roles {
	return r
}

func (r Roles) Check(permissions []model.Permissions, privilege string, on *model.Entities) bool {
	for _, permission := range permissions {
		roleID := permission.Role

		privilegeCategory := p.Category(privilege)
		if privilegeCategory == p.GlobalPrivileges {
			globalPrivileges := r[roleID].Grants.GlobalPrivileges
			return checkPermission(privilege, *globalPrivileges)
		} else if privilegeCategory == p.DatabasePrivileges {
			if *on.Databases == *permission.On.Databases {
				databasePrivileges := r[roleID].Grants.DatabasePrivileges
				return checkPermission(privilege, *databasePrivileges)
			}
			return denied
		} else if privilegeCategory == p.CollectionPrivileges {
			if *on.Databases == *permission.On.Databases {
				if on.Collections == permission.On.Collections {
					collectionPrivileges := r[roleID].Grants.CollectionPrivileges
					return checkPermission(privilege, *collectionPrivileges)
				}
				return denied
			}
			return denied
		}
	}
	return denied
}

func checkPermission(privilege string, rolePrivileges model.Privileges) bool {
	privilegeAllowed, privilegeExists := rolePrivileges[privilege]
	if privilegeExists {
		if privilegeAllowed == p.Allowed {
			return p.Allowed
		}
	}
	return p.Denied
}

func getPrivileges(allowedActions, deniedActions []string) (*model.Grants, error) {

	var grants model.Grants

	allowedGrantsMap := sortPrivileges(allowedActions, p.Allowed)
	deniedGrantsMap := sortPrivileges(deniedActions, p.Denied)

	globalPrivileges, err := mergeGrantMaps(maps.Marshal(allowedGrantsMap.GlobalPrivileges), maps.Marshal(deniedGrantsMap.GlobalPrivileges))
	if err != nil {
		return nil, err
	}

	databasePrivileges, err := mergeGrantMaps(maps.Marshal(allowedGrantsMap.DatabasePrivileges), maps.Marshal(deniedGrantsMap.DatabasePrivileges))
	if err != nil {
		return nil, err
	}

	collectionPrivileges, err := mergeGrantMaps(maps.Marshal(allowedGrantsMap.CollectionPrivileges), maps.Marshal(deniedGrantsMap.CollectionPrivileges))
	if err != nil {
		return nil, err
	}

	grants.GlobalPrivileges = globalPrivileges
	grants.DatabasePrivileges = databasePrivileges
	grants.CollectionPrivileges = collectionPrivileges

	return &grants, nil

}

func sortPrivileges(actions []string, assignedPermission bool) model.Grants {
	var privilegeGrants model.Grants

	globalPrivileges := model.Privileges{}
	databasePrivileges := model.Privileges{}
	collectionPrivileges := model.Privileges{}

	for _, action := range actions {
		if p.IsAvailable(action) {
			actionCategory := p.Category(action)
			switch actionCategory {
			case p.DatabasePrivileges:
				databasePrivileges[action] = assignedPermission
			case p.CollectionPrivileges:
				collectionPrivileges[action] = assignedPermission
			case p.GlobalPrivileges:
				globalPrivileges[action] = assignedPermission
			}
		}
	}

	privilegeGrants.GlobalPrivileges = &globalPrivileges
	privilegeGrants.DatabasePrivileges = &databasePrivileges
	privilegeGrants.CollectionPrivileges = &collectionPrivileges

	return privilegeGrants
}

func mergeGrantMaps(allowedPrivilegesMap, deniedPrivilegesMap map[string]interface{}) (*model.Privileges, error) {
	var privileges model.Privileges

	mergeableGrantMaps := []map[string]interface{}{
		allowedPrivilegesMap,
		deniedPrivilegesMap,
	}

	mergedGrantsMap, err := maps.Merge(mergeableGrantMaps...)
	if err != nil {
		return nil, err
	}

	mergedGrantBytes, err := json.Marshal(mergedGrantsMap)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(mergedGrantBytes, &privileges)

	return &privileges, nil
}
