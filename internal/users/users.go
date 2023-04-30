package users

import (
	"github.com/TanmoySG/wunderDB/model"
)

const (
	passwordAuthentication = "password"
	tokenAuthentication    = "token"

	userExists       = true
	userDoesNotExist = false
)

type Users map[model.Identifier]*model.User

func From(users Users) Users {
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

func (u Users) CreateUser(userID model.Identifier, hashedPassword string, hashingAlgorithm string, metadata model.Metadata) {
	u[userID] = &model.User{
		UserID:   userID,
		Metadata: metadata,
		Authentication: model.Authentication{
			HashingAlgorithm: hashingAlgorithm,
			HashedSecret:     hashedPassword,
		},
		Permissions: []model.Permissions{},
	}
}

func (u Users) GetUser(userID model.Identifier) *model.User {
	user := u[userID]
	redactedUser := model.User{
		UserID:   user.UserID,
		Metadata: user.Metadata,
		Authentication: model.Authentication{
			HashingAlgorithm: user.Authentication.HashingAlgorithm,
		},
	}
	return &redactedUser
}

func (u Users) GrantRole(userID model.Identifier, permissions model.Permissions) {
	// Permissions added latest have higher priority
	u[userID].Permissions = append([]model.Permissions{permissions}, u[userID].Permissions...)
}

func (u Users) Permission(userID model.Identifier) []model.Permissions {
	return u[userID].Permissions
}

func (u Users) Authenticate(userID model.Identifier, hashedPassword string) bool {
	return u[userID].Authentication.HashedSecret == hashedPassword
}
