package admin

import (
	"testing"

	p "github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/stretchr/testify/assert"
)

func Test_getAllowedRole(t *testing.T) {
	got := getAllowedRole()
	assert.Equal(t, len(p.PrivilegeScope), len(got))
}

func Test_getPermission(t *testing.T) {
	want := model.Permissions{
		Role: "wdb_super_admin_role",
		On: &model.Entities{
			Users:       &WILDCARD,
			Databases:   &WILDCARD,
			Collections: &WILDCARD,
		},
	}

	got := getPermission()

	assert.Equal(t, want, got)
}

func Test_decodeDefaultPassword(t *testing.T) {
	testCases := []struct {
		base64Encoded   string
		expectedDecoded string
	}{
		{
			base64Encoded:   BASE64_DEFAULT_ADMIN_PASSWORD,
			expectedDecoded: "admin",
		},
		{
			base64Encoded:   "YWRtaW4",
			expectedDecoded: "",
		},
	}

	for _, tc := range testCases {
		decoded := decodeDefaultPassword(tc.base64Encoded)
		assert.Equal(t, tc.expectedDecoded, decoded)
	}
}
