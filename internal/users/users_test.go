package users

import (
	"fmt"
	"testing"

	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/stretchr/testify/assert"
)

var (
	hashedSecret string = "hashedSecret"

	user1Name model.Identifier = "user1"
	user1     *model.User      = &model.User{
		UserID:      user1Name,
		Metadata:    model.Metadata{},
		Permissions: []model.Permissions{},
		Authentication: model.Authentication{
			HashedSecret:     hashedSecret,
			HashingAlgorithm: authentication.MD5,
		},
	}

	user2Name model.Identifier = "user2"
	user2     *model.User      = &model.User{
		UserID:      user2Name,
		Metadata:    model.Metadata{},
		Permissions: []model.Permissions{},
		Authentication: model.Authentication{
			HashedSecret:     hashedSecret,
			HashingAlgorithm: authentication.SHA256,
		},
	}

	testRoleName   model.Identifier  = "testRole"
	testEntity     string            = "test"
	testPermission model.Permissions = model.Permissions{
		Role: testRoleName,
		On: &model.Entities{
			Databases:   &testEntity,
			Collections: &testEntity,
			Users:       &testEntity,
		},
	}
)

func initTestUserObject() Users {
	usersList := Users{user1Name: user1}
	return From(usersList)
}

func Test_CheckIfExists(t *testing.T) {
	testCases := []struct {
		userId           model.Identifier
		expectedUser     *model.User
		expectedIsExists bool
	}{
		{
			userId:           user1Name,
			expectedUser:     user1,
			expectedIsExists: true,
		},
		{
			userId:           user2Name,
			expectedUser:     nil,
			expectedIsExists: false,
		},
	}

	u := initTestUserObject()

	for _, tc := range testCases {
		isExists, user := u.CheckIfExists(model.Identifier(tc.userId))
		assert.Equal(t, tc.expectedUser, user)
		assert.Equal(t, tc.expectedIsExists, isExists)
	}
}

func Test_CreateUser(t *testing.T) {
	want := Users{
		user1Name: user1,
		user2Name: user2,
	}

	u := initTestUserObject()
	u.CreateUser(user2Name, hashedSecret, authentication.SHA256, model.Metadata{})

	gotUser, isExists := u[user2Name]
	assert.Equal(t, user2, gotUser)
	assert.Equal(t, userExists, isExists)
	assert.Equal(t, want, u)
}

func Test_GetUser(t *testing.T) {
	wantUser := &model.User{
		UserID:   user1Name,
		Metadata: user1.Metadata,
		Authentication: model.Authentication{
			HashingAlgorithm: user1.Authentication.HashingAlgorithm,
		},
	}

	u := initTestUserObject()
	gotUser := u.GetUser(user1Name)

	assert.Equal(t, wantUser, gotUser)
}

func Test_GrantRole(t *testing.T) {
	u := initTestUserObject()

	gotUserBeforeGrant, isExists := u[user1Name]
	assert.Equal(t, userExists, isExists)
	assert.Equal(t, []model.Permissions{}, gotUserBeforeGrant.Permissions)

	wantPermissions := []model.Permissions{testPermission}

	u.GrantRole(user1Name, testPermission)
	gotUserAfterGrant, isExists := u[user1Name]
	assert.Equal(t, userExists, isExists)
	assert.Equal(t, wantPermissions, gotUserAfterGrant.Permissions)
}

func Test_RevokeRole(t *testing.T) {
	u := initTestUserObject()
	u[user1Name].Permissions = []model.Permissions{testPermission}

	u.RevokeRole(user1Name, testPermission)
	gotUserAfterGrant, isExists := u[user1Name]
	assert.Equal(t, userExists, isExists)
	assert.Equal(t, []model.Permissions{}, gotUserAfterGrant.Permissions)
}

func Test_Permission(t *testing.T) {
	u := initTestUserObject()

	gotPermissions := u.Permission(user1Name)
	assert.Equal(t, user1.Permissions, gotPermissions)
}

func Test_Authenticate(t *testing.T) {
	testCases := []struct {
		userId       model.Identifier
		hashedSecret string
		expected     bool
	}{
		{
			userId:       user1Name,
			hashedSecret: hashedSecret,
			expected:     true,
		},
		{
			userId:       user1Name,
			hashedSecret: fmt.Sprintf("invalid-%s", hashedSecret),
			expected:     false,
		},
	}

	u := initTestUserObject()

	for _, tc := range testCases {
		isAuthentic := u.Authenticate(tc.userId, tc.hashedSecret)
		assert.Equal(t, tc.expected, isAuthentic)
	}
}
