package users

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wunderDB/model"
)

type Users map[model.Identifier]*model.User

func GetAdminUser() model.User {
	return model.User{
		UserID:         "admin",
		Authentication: model.Authentication{},
		Metadata: model.Metadata{
			"email": "",
		},
	}
}

func MarshalAdmin(admin model.User) {
	namespacesJson, err := json.Marshal(admin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(namespacesJson))
}
