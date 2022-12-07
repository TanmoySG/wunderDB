package wdbClient

import (
	c "github.com/TanmoySG/wunderDB/internal/collections"
	d "github.com/TanmoySG/wunderDB/internal/data"
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/identities"

	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddData(databaseId, collectionId model.Identifier, inputData interface{}) *er.WdbError {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return &er.DatabaseDoesNotExistsError
	}

	dataId := identities.GenerateID()
	data := d.UseCollection(*collection)

	return data.Add(model.Identifier(dataId), inputData)
}

func (wdb wdbClient) GetData(databaseId, collectionId model.Identifier, filters interface{}) (map[model.Identifier]*model.Datum, *er.WdbError) {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return nil, &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return nil, &er.DatabaseDoesNotExistsError
	}

	data := d.UseCollection(*collection)

	fetchedData, err := data.Read(filters)
	if err != nil {
		return nil, err
	}

	return fetchedData, nil
}

func (wdb wdbClient) UpdateData(databaseId, collectionId model.Identifier, updatedData, filters interface{}) *er.WdbError {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return &er.DatabaseDoesNotExistsError
	}

	data := d.UseCollection(*collection)

	err := data.Update(updatedData, filters)
	if err != nil {
		return err
	}

	return nil
}

func (wdb wdbClient) DeleteData(databaseId, collectionId model.Identifier, filters interface{}) *er.WdbError {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return &er.DatabaseDoesNotExistsError
	}

	data := d.UseCollection(*collection)

	err := data.Delete(filters)
	if err != nil {
		return err
	}

	return nil
}
