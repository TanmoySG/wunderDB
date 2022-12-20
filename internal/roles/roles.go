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

func (r Roles) CreateRole(roleID model.Identifier, allowedActions []string, deniedActions []string) {
	privileges := model.Privileges{}

	for _, action := range allowedActions {
		if p.IsAvailable(action) {
			privileges[action] = p.Allowed
		}
	}

	for _, action := range deniedActions {
		if p.IsAvailable(action) {
			privileges[action] = p.Denied
		}
	}

	r[roleID] = &model.Role{
		RoleID:     roleID,
		Privileges: privileges,
	}
}

func (r Roles) ListRoles() Roles {
	return r
}
