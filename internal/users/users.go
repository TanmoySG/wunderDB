package users

import (
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
)

const (
	passwordAuthentication = "password"
	tokenAuthentication    = "token"

	userExists       = true
	userDoesNotExist = false
)

type Users map[model.Identifier]*model.User

func Use(users Users) Users {
	return users
}

func (u Users) CheckIfExists(userID model.Identifier) (bool, *model.User) {
	user, dbExists := u[userID]
	if dbExists {
		return userExists, user
	} else {
		return userDoesNotExist, user
	}
}

func (u Users) CreateUser(userID model.Identifier, password string, hashingAlgorithm string, metadata model.Metadata) {
	u[userID] = &model.User{
		UserID:   userID,
		Metadata: metadata,
		Authentication: model.Authentication{
			HashingAlgorithm: hashingAlgorithm,
			HashedSecret:     authentication.Hash(password, hashingAlgorithm),
		},
		Permissions: []model.Permissions{},
	}
}

func (u Users) GrantRole(userID model.Identifier, permissions []model.Permissions) {
	// Permissions added latest have higher priority
	u[userID].Permissions = append(permissions, u[userID].Permissions...)
}

func (u Users) Permission(userID model.Identifier) []model.Permissions {
	return u[userID].Permissions
}
