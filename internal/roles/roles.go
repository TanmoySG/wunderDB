package roles

import (
	p "github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/model"
)

const (
	roleExists       = true
	roleDoesNotExist = false
)

type Roles map[model.Identifier]*model.Role

type Permissions struct {
	DatabasePermissions   *Privileges `json:"databasePermissions,omitempty"`
	CollectionPermissions *Privileges `json:"collectionPermissions,omitempty"`
	DataPermissions       *Privileges `json:"dataPermissions,omitempty"`
	UserPermissions       *Privileges `json:"userPermissions,omitempty"`
	RolePermissions       *Privileges `json:"rolePermissions,omitempty"`
}

type Privileges struct {
	Allowed []string `json:"allowed,omitempty"`
	Denied  []string `json:"denied,omitempty"`
}

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

func (r Roles) CreateRole(roleID model.Identifier, permissions Permissions) {
	grants := model.Grants{}

	if permissions.DatabasePermissions != nil {
		grants.DatabasePrivileges = getPrivileges(permissions.DatabasePermissions.Allowed, permissions.DatabasePermissions.Denied)
	}

	if permissions.CollectionPermissions != nil {
		grants.CollectionPrivileges = getPrivileges(permissions.CollectionPermissions.Allowed, permissions.CollectionPermissions.Denied)
	}

	if permissions.DataPermissions != nil {
		grants.DataPrivileges = getPrivileges(permissions.DataPermissions.Allowed, permissions.DataPermissions.Denied)
	}

	if permissions.UserPermissions != nil {
		grants.UserPrivileges = getPrivileges(permissions.UserPermissions.Allowed, permissions.UserPermissions.Denied)
	}

	if permissions.UserPermissions != nil {
		grants.RolePrivileges = getPrivileges(permissions.RolePermissions.Allowed, permissions.RolePermissions.Denied)
	}

	r[roleID] = &model.Role{
		RoleID: roleID,
		Grants: grants,
	}
}

func (r Roles) ListRoles() Roles {
	return r
}

func (r Roles) Check(permissions []model.Permissions, privilege string, on *model.Entities) bool {
	for _, permission := range permissions {

		roleID := permission.Role
		if on.Databases != nil && on.Databases == permission.On.Databases {
			databasePrivileges := r[roleID].Grants.DatabasePrivileges
			return checkPermission(privilege, *databasePrivileges)
		} else if on.Collections != nil && on.Collections == permission.On.Collections {
			collectionPrivileges := r[roleID].Grants.CollectionPrivileges
			return checkPermission(privilege, *collectionPrivileges)
		} else if on.Roles != nil && (*on.Roles) == p.Allowed {
			rolePrivileges := r[roleID].Grants.CollectionPrivileges
			return checkPermission(privilege, *rolePrivileges)
		}
	}
	return false
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

func getPrivileges(allowedActions []string, deniedActions []string) *model.Privileges {
	privileges := model.Privileges{}

	for _, action := range allowedActions {
		if p.IsAvailable(action) {
			privileges[action] = p.Allowed
		}
	}

	for _, action := range deniedActions {
		if p.IsAvailable(action) {
			privileges[action] = p.Allowed
		}
	}

	return &privileges
}
