package roles

import (
	"encoding/json"

	p "github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/utils/maps"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

const (
	roleExists       = true
	roleDoesNotExist = false

	allowed = true
	denied  = false
)

type Roles map[model.Identifier]*model.Role

func From(roles Roles) Roles {
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
	var filteredRoles = make(Roles)
	for roleId, role := range r {
		if !role.Hidden {
			filteredRoles[roleId] = role
		}
	}
	return filteredRoles
}

func (r Roles) Check(permissions []model.Permissions, privilege string, on *model.Entities) bool {
	for _, userPermission := range permissions {
		roleID := userPermission.Role

		validRole, role := r.CheckIfExists(roleID)
		if !validRole {
			continue
		}

		privilegeCategory := p.Category(privilege)
		if privilegeCategory == p.GlobalPrivileges {
			globalPrivileges := role.Grants.GlobalPrivileges
			return checkPermission(privilege, *globalPrivileges)
		} else if privilegeCategory == p.UserPrivileges {
			if userPermission.On.Users != nil {
				// add condition *on.Users == *userPermission.On.Users
				if *userPermission.On.Users == p.Wildcard {
					if on.Databases != nil {
						if on.Collections != nil {
							if *on.Collections == *userPermission.On.Collections || *userPermission.On.Collections == p.Wildcard {
								collectionPrivileges := role.Grants.CollectionPrivileges
								return checkPermission(privilege, *collectionPrivileges)
							}
							return denied
						}

						if *on.Databases == *userPermission.On.Databases || *userPermission.On.Databases == p.Wildcard {
							databasePrivileges := role.Grants.DatabasePrivileges
							return checkPermission(privilege, *databasePrivileges)
						}
						return denied
					}
					userPrivileges := role.Grants.UserPrivileges
					return checkPermission(privilege, *userPrivileges)
				}
				return denied
			}
			return denied
		} else if privilegeCategory == p.DatabasePrivileges {
			if on.Databases != nil && userPermission.On.Databases != nil {
				if *on.Databases == *userPermission.On.Databases || *userPermission.On.Databases == p.Wildcard {
					databasePrivileges := role.Grants.DatabasePrivileges
					return checkPermission(privilege, *databasePrivileges)
				}
			}
			return denied
		} else if privilegeCategory == p.CollectionPrivileges {
			if *on.Databases == *userPermission.On.Databases || *userPermission.On.Databases == p.Wildcard {
				if *on.Collections == *userPermission.On.Collections || *userPermission.On.Collections == p.Wildcard {
					collectionPrivileges := role.Grants.CollectionPrivileges
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

	userPrivileges, err := mergeGrantMaps(maps.Marshal(allowedGrantsMap.UserPrivileges), maps.Marshal(deniedGrantsMap.UserPrivileges))
	if err != nil {
		return nil, err
	}

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
	grants.UserPrivileges = userPrivileges

	return &grants, nil

}

func sortPrivileges(actions []string, assignedPermission bool) model.Grants {
	var privilegeGrants model.Grants

	userPrivileges := model.Privileges{}
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
			case p.UserPrivileges:
				userPrivileges[action] = assignedPermission
			}
		}
	}

	privilegeGrants.UserPrivileges = &userPrivileges
	privilegeGrants.GlobalPrivileges = &globalPrivileges
	privilegeGrants.DatabasePrivileges = &databasePrivileges
	privilegeGrants.CollectionPrivileges = &collectionPrivileges

	return privilegeGrants
}

func mergeGrantMaps(allowedPrivilegesMap, deniedPrivilegesMap map[string]interface{}) (*model.Privileges, error) {
	var privileges model.Privileges

	mergedGrantsMap, err := maps.Merge(allowedPrivilegesMap, deniedPrivilegesMap)
	if err != nil {
		return nil, err
	}

	mergedGrantBytes, err := json.Marshal(mergedGrantsMap)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(mergedGrantBytes, &privileges)
	if err != nil {
		return nil, err
	}
	return &privileges, nil
}
