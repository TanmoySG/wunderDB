package databases

import (
	"testing"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/stretchr/testify/assert"
)

var (
	testDatabase1Name model.Identifier = "testDatabase1"
	testDatabase1     *model.Database  = &model.Database{
		Collections: map[model.Identifier]*model.Collection{},
		Metadata:    model.Metadata{},
	}

	testDatabases Databases = Databases{
		testDatabase1Name: testDatabase1,
	}
)

func Test_CheckIfExists(t *testing.T) {
	testCases := []struct {
		databaseName     string
		expectedDatabase *model.Database
		expectedIsExists bool
	}{
		{
			databaseName:     string(testDatabase1Name),
			expectedDatabase: testDatabase1,
			expectedIsExists: true,
		},
		{
			databaseName:     "db2",
			expectedDatabase: nil,
			expectedIsExists: false,
		},
	}

	dbs := From(testDatabases)

	for _, tc := range testCases {
		isExists, database := dbs.CheckIfExists(model.Identifier(tc.databaseName))
		assert.Equal(t, tc.expectedDatabase, database)
		assert.Equal(t, tc.expectedIsExists, isExists)
	}

}
