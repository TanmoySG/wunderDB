package roles

import (
	"encoding/json"

	p "github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/utils/maps"
)

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
