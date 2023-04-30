package data

import (
	"testing"
	"time"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/stretchr/testify/assert"
)

var (
	testSchema = model.Schema{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
			"age": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{
			"name",
			"age",
		},
	}

	testFilter = map[string]interface{}{"key": "name", "value": "john"}

	collection = model.Collection{
		Schema: testSchema,
		Data:   map[model.Identifier]*model.Datum{},
	}
)

func Test_HappyDataFlow(t *testing.T) {

	dc := UseCollection(collection)

	datarow1Id := model.Identifier("1")
	datarow1 := map[string]interface{}{
		"name": "john",
		"age":  "29",
	}

	datarow2Id := model.Identifier("2")
	datarow2 := map[string]interface{}{
		"name": "jane",
		"age":  "28",
	}

	currentTimestamp := time.Now().UTC().Unix()
	metadata1 := model.Metadata{
		"created_at": currentTimestamp,
		"updated_at": currentTimestamp,
	}

	want := map[model.Identifier]*model.Datum{
		datarow1Id: {
			Data:       datarow1,
			Identifier: datarow1Id,
			Metadata:   metadata1,
		},
		datarow2Id: {
			Data:       datarow2,
			Identifier: datarow2Id,
			Metadata:   metadata1,
		},
	}

	err := dc.Add(datarow1Id, datarow1)
	assert.Nil(t, err)

	err = dc.Add(datarow2Id, datarow2)
	assert.Nil(t, err)

	assert.Equal(t, want, dc.Data)

	// Read Data
	fetchedData, err := dc.Read(nil)
	assert.Nil(t, err)
	assert.Equal(t, want, fetchedData)

	// Read Data - filtered
	filteredData := map[model.Identifier]*model.Datum{
		datarow1Id: {
			Data:       datarow1,
			Identifier: datarow1Id,
			Metadata:   metadata1,
		},
	}

	fetchedData, err = dc.Read(testFilter)
	assert.Nil(t, err)
	assert.Equal(t, filteredData, fetchedData)

	// Update Data
	updateData := map[string]interface{}{
		"age": "30",
	}

	updatedTimestamp := time.Now().UTC().Unix()
	updatedMetadata := model.Metadata{
		"created_at": currentTimestamp,
		"updated_at": updatedTimestamp,
	}
	updatedDatarow1 := map[string]interface{}{
		"name": "john",
		"age":  "30",
	}
	wantChange := &model.Datum{
		Data:       &updatedDatarow1,
		Identifier: datarow1Id,
		Metadata:   updatedMetadata,
	}

	err = dc.Update(updateData, testFilter)
	assert.Nil(t, err)
	assert.Equal(t, wantChange, dc.Data[datarow1Id])

	// delete
	want = map[model.Identifier]*model.Datum{
		datarow2Id: {
			Data:       datarow2,
			Identifier: datarow2Id,
			Metadata:   metadata1,
		},
	}

	err = dc.Delete(testFilter)
	assert.Nil(t, err)
	assert.Equal(t, want, dc.Data)
}
