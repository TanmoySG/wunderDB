package roles

import (
	p "github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/model"
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

func (r Roles) Create(roleID model.Identifier, allowed []string, denied []string, hidden bool) *er.WdbError {

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
		Hidden: hidden,
	}

	return nil

}

func (r Roles) List(forceListAllRoles bool) Roles {
	var filteredRoles = make(Roles)

	if forceListAllRoles {
		return r
	}

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

func (r Roles) Update(roleID model.Identifier, allowed []string, denied []string, hidden bool) *er.WdbError {
	grants, err := getPrivileges(allowed, denied)
	if err != nil {
		return &er.WdbError{
			ErrCode:        er.EncodeDecodeError.ErrCode,
			ErrMessage:     err.Error(),
			HttpStatusCode: er.EncodeDecodeError.HttpStatusCode,
		}
	}

	r[roleID].Grants = *grants
	r[roleID].Hidden = hidden

	return nil
}
