package collections

import (
	"testing"
	"time"

	"github.com/TanmoySG/wunderDB/internal/metadata"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/stretchr/testify/assert"
)

var (
	fixedTimestamp = 1682707391 //Friday, April 28, 2023 6:43:11 PM (GMT)

	testCollection1Name model.Identifier  = "testCollection1"
	testCollection1     *model.Collection = &model.Collection{
		Schema: model.Schema{},
		Metadata: model.Metadata{
			metadata.CreatedAt: fixedTimestamp,
			metadata.UpdatedAt: fixedTimestamp,
		},
		Data: map[model.Identifier]*model.Datum{},
	}

	testCollection2Name model.Identifier  = "testCollection2"
	testCollection2     *model.Collection = &model.Collection{
		Schema:   model.Schema{},
		Metadata: model.Metadata{},
		Data:     map[model.Identifier]*model.Datum{},
	}
)

func Test_CheckIfExists(t *testing.T) {
	testCases := []struct {
		collectionName     string
		expectedCollection *model.Collection
		expectedIsExists   bool
	}{
		{
			collectionName:     string(testCollection1Name),
			expectedCollection: testCollection1,
			expectedIsExists:   true,
		},
		{
			collectionName:     "db2",
			expectedCollection: nil,
			expectedIsExists:   false,
		},
	}

	database := model.Database{Collections: Collections{testCollection1Name: testCollection1}}
	collection := UseDatabase(&database)

	for _, tc := range testCases {
		isExists, collection := collection.CheckIfExists(model.Identifier(tc.collectionName))
		assert.Equal(t, tc.expectedCollection, collection)
		assert.Equal(t, tc.expectedIsExists, isExists)
	}
}

func Test_CreateCollection(t *testing.T) {
	testCollection := Collections{testCollection1Name: testCollection1}
	database := model.Database{Collections: testCollection}

	collections := UseDatabase(&database)
	collections.CreateCollection(testCollection2Name, model.Schema{})

	expectedCollectionsMap := testCollection
	expectedCollectionsMap[testCollection2Name] = testCollection2
	expectedCollectionsMap[testCollection2Name].Metadata = metadata.New().BasicChangeMetadata()

	retrievedCollection, exists := collections[testCollection2Name]

	assert.Equal(t, true, exists)
	assert.Equal(t, expectedCollectionsMap, collections)
	assert.Equal(t, expectedCollectionsMap[testCollection2Name], retrievedCollection)
	assert.Equal(t, expectedCollectionsMap[testCollection2Name].Metadata, retrievedCollection.Metadata)
}

func Test_GetCollection(t *testing.T) {
	testCases := []struct {
		collectionName     string
		expectedCollection *model.Collection
	}{
		{
			collectionName:     string(testCollection1Name),
			expectedCollection: testCollection1,
		},
		{
			collectionName:     string(testCollection2Name),
			expectedCollection: nil,
		},
	}

	database := model.Database{Collections: Collections{testCollection1Name: testCollection1}}
	collection := UseDatabase(&database)

	for _, tc := range testCases {
		fetchedDatabase := collection.GetCollection(model.Identifier(tc.collectionName))
		assert.Equal(t, tc.expectedCollection, fetchedDatabase)
	}
}

func Test_DeleteDatabase(t *testing.T) {
	database := model.Database{Collections: Collections{testCollection1Name: testCollection1}}
	collection := UseDatabase(&database)
	collection.DeleteCollection(testCollection1Name)

	assert.Equal(t, Collections{}, collection)
}

func Test_UpdateMetadata(t *testing.T) {
	expectedCollectionsChange := Collections{
		testCollection1Name: &model.Collection{
			Data: map[model.Identifier]*model.Datum{},
			Metadata: model.Metadata{
				metadata.CreatedAt: fixedTimestamp,
				metadata.UpdatedAt: time.Now().UTC().Unix(),
			},
			Schema: model.Schema{},
		},
	}

	database := model.Database{Collections: Collections{testCollection1Name: testCollection1}}
	collection := UseDatabase(&database)
	collection.UpdateMetadata(testCollection1Name)

	assert.Equal(t, expectedCollectionsChange[testCollection1Name].Metadata[metadata.CreatedAt], collection[testCollection1Name].Metadata[metadata.CreatedAt])
	assert.Equal(t, expectedCollectionsChange[testCollection1Name].Metadata[metadata.UpdatedAt], collection[testCollection1Name].Metadata[metadata.UpdatedAt])
}
