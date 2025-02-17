package databases

import (
	"testing"
	"time"

	"github.com/TanmoySG/wunderDB/internal/metadata"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/model/redacted"
	"github.com/stretchr/testify/assert"
)

var (
	fixedTimestamp = 1682707391 //Friday, April 28, 2023 6:43:11 PM (GMT)

	testDatabase1Name model.Identifier = "testDatabase1"
	testDatabase1     *model.Database  = &model.Database{
		Collections: map[model.Identifier]*model.Collection{},
		Metadata: model.Metadata{
			metadata.CreatedAt: fixedTimestamp,
			metadata.UpdatedAt: fixedTimestamp,
		},
	}

	redactedTestDatabase1 *redacted.RedactedD = &redacted.RedactedD{
		Collections: []model.Identifier{},
		Metadata: model.Metadata{
			metadata.CreatedAt: fixedTimestamp,
			metadata.UpdatedAt: fixedTimestamp,
		},
	}

	testDatabase2Name model.Identifier = "testDatabase2"
	testDatabase2     *model.Database  = &model.Database{
		Collections: map[model.Identifier]*model.Collection{},
		Metadata:    model.Metadata{},
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

	testDatabases := Databases{testDatabase1Name: testDatabase1}
	dbs := From(testDatabases)

	for _, tc := range testCases {
		isExists, database := dbs.CheckIfExists(model.Identifier(tc.databaseName))
		assert.Equal(t, tc.expectedDatabase, database)
		assert.Equal(t, tc.expectedIsExists, isExists)
	}

}

func Test_CreateDatabase(t *testing.T) {
	testDatabases := Databases{testDatabase1Name: testDatabase1}
	dbs := From(testDatabases)
	dbs.CreateDatabase(testDatabase2Name, model.Access{})

	expectedDatabasesMap := testDatabases
	expectedDatabasesMap[testDatabase2Name] = testDatabase2
	expectedDatabasesMap[testDatabase2Name].Metadata = metadata.New().BasicChangeMetadata()

	database, exists := dbs[testDatabase2Name]

	assert.Equal(t, true, exists)
	assert.Equal(t, expectedDatabasesMap[testDatabase2Name], database)
	assert.Equal(t, expectedDatabasesMap, dbs)
	assert.Equal(t, expectedDatabasesMap[testDatabase2Name].Metadata, database.Metadata)
}

func Test_GetDatabase(t *testing.T) {
	testCases := []struct {
		databaseName     string
		expectedDatabase *redacted.RedactedD
	}{
		{
			databaseName:     string(testDatabase1Name),
			expectedDatabase: redactedTestDatabase1,
		},
		{
			databaseName:     string(testDatabase2Name),
			expectedDatabase: nil,
		},
	}
	testDatabases := Databases{testDatabase1Name: testDatabase1}

	dbs := From(testDatabases)
	for _, tc := range testCases {
		fetchedDatabase := dbs.GetDatabase(model.Identifier(tc.databaseName))
		assert.Equal(t, tc.expectedDatabase, fetchedDatabase)
	}
}

func Test_DeleteDatabase(t *testing.T) {
	testDatabases := Databases{testDatabase1Name: testDatabase1}

	dbs := From(testDatabases)
	dbs.DeleteDatabase(testDatabase1Name)
	assert.Equal(t, Databases{}, dbs)
}

func Test_UpdateMetadata(t *testing.T) {
	expectedDatabasesChange := Databases{
		testDatabase1Name: &model.Database{
			Collections: map[model.Identifier]*model.Collection{},
			Metadata: model.Metadata{
				metadata.CreatedAt: fixedTimestamp,
				metadata.UpdatedAt: time.Now().UTC().Unix(),
			},
		},
	}
	testDatabases := Databases{testDatabase1Name: testDatabase1}

	dbs := From(testDatabases)
	dbs.UpdateMetadata(testDatabase1Name)

	assert.Equal(t, expectedDatabasesChange[testDatabase1Name].Metadata[metadata.CreatedAt], dbs[testDatabase1Name].Metadata[metadata.CreatedAt])
	assert.Equal(t, expectedDatabasesChange[testDatabase1Name].Metadata[metadata.UpdatedAt], dbs[testDatabase1Name].Metadata[metadata.UpdatedAt])
}
