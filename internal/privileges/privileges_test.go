package privileges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	privilege             string
	expectedIsAvailable   bool
	expectedCategory      string
	expectedPrivilegeType PrivilegeActionType
}{
	{
		privilege:             CreateCollection,
		expectedIsAvailable:   true,
		expectedCategory:      PrivilegeScope[CreateCollection],
		expectedPrivilegeType: PrivilegeType[CreateCollection],
	},
	{
		privilege:             "random",
		expectedIsAvailable:   false,
		expectedCategory:      "",
		expectedPrivilegeType: WildcardPrivilege,
	},
}

func Test_IsAvailable(t *testing.T) {
	for _, tc := range testCases {
		isAvailable := IsAvailable(tc.privilege)
		assert.Equal(t, tc.expectedIsAvailable, isAvailable)
	}
}

func Test_Category(t *testing.T) {
	for _, tc := range testCases {
		category := Category(tc.privilege)
		assert.Equal(t, tc.expectedCategory, category)
	}
}

func Test_GetPrivilegeType(t *testing.T) {
	for _, tc := range testCases {
		privilegeTypee := GetPrivilegeType(tc.privilege)
		assert.Equal(t, tc.expectedPrivilegeType, privilegeTypee)
	}
}
