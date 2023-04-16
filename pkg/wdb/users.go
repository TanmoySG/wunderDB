package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

var wildcard = privileges.Wildcard

func (wdb wdbClient) CreateUser(userID model.Identifier, password string) *er.WdbError {
	if ! wdb.safeName.Check(userID.String()) {
		return &er.EntityNameFormatError
	}

	if exists, _ := wdb.Users.CheckIfExists(userID); exists {
		return &er.UserAlreadyExistsError
	}
	hashedPassword := authentication.Hash(password, wdb.HashingAlgorithm)
	wdb.Users.CreateUser(userID, hashedPassword, wdb.HashingAlgorithm, model.Metadata{})
	return nil
}

func (wdb wdbClient) GrantRoles(userID model.Identifier, permission model.Permissions) *er.WdbError {
	if ! wdb.safeName.Check(userID.String()) {
		return &er.EntityNameFormatError
	}

	if exists, _ := wdb.Users.CheckIfExists(userID); !exists {
		return &er.UserDoesNotExistError
	}

	// initialize permission entities if not present
	if permission.On == nil {
		permission.On = &model.Entities{}
	}

	// add wildcard (all) users if none is specified
	if permission.On.Users == nil {
		permission.On.Users = &wildcard
	}

	validRole, _ := wdb.Roles.CheckIfExists(permission.Role)
	if !validRole {
		return &er.InvalidRoleError
	}

	wdb.Users.GrantRole(userID, permission)

	return nil
}

func (wdb wdbClient) AuthenticateUser(userID model.Identifier, password string) (bool, *er.WdbError) {
	if ! wdb.safeName.Check(userID.String()) {
		return authentication.InvalidUser, &er.EntityNameFormatError
	}

	if exists, _ := wdb.Users.CheckIfExists(userID); !exists {
		return authentication.InvalidUser, &er.AuthenticatingUserDoesNotExist
	}
	user := wdb.Users.GetUser(userID)
	hashedPassword := authentication.Hash(password, user.Authentication.HashingAlgorithm)
	isAuthentic := wdb.Users.Authenticate(userID, hashedPassword)
	if isAuthentic {
		return authentication.ValidUser, nil
	}
	return authentication.InvalidUser, &er.InvalidCredentialsError
}
