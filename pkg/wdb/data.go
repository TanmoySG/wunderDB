package wdbClient

import (
	"fmt"

	c "github.com/TanmoySG/wunderDB/internal/collections"
	d "github.com/TanmoySG/wunderDB/internal/data"
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/identities"

	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddData(databaseId, collectionId model.Identifier, inputData interface{}) error {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return fmt.Errorf("error deleting database %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	collections := c.UseDatabase(*database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return fmt.Errorf("error adding data: %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	dataId := identities.GenerateID()
	data := d.UseCollection(*collection)

	return data.Add(model.Identifier(dataId), inputData)
}

func (wdb wdbClient) GetData(databaseId, collectionId model.Identifier, filters interface{}) (map[model.Identifier]*model.Datum, error) {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return nil, fmt.Errorf("error adding data %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	collections := c.UseDatabase(*database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return nil, fmt.Errorf("error adding data: %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	data := d.UseCollection(*collection)

	fetchedData, err := data.Read(filters)
	if err != nil {
		return nil, fmt.Errorf("error adding data %s", err)
	}

	return fetchedData, nil
}

func (wdb wdbClient) UpdateData(databaseId, collectionId model.Identifier, updatedData, filters interface{}) error {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return fmt.Errorf("error adding data %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	collections := c.UseDatabase(*database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return fmt.Errorf("error adding data: %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	data := d.UseCollection(*collection)

	err := data.Update(updatedData, filters)
	if err != nil {
		return fmt.Errorf("error adding data %s", err)
	}

	return nil
}

func (wdb wdbClient) DeleteData(databaseId, collectionId model.Identifier, filters interface{}) error {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return fmt.Errorf("error adding data %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	collections := c.UseDatabase(*database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return fmt.Errorf("error adding data: %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	data := d.UseCollection(*collection)

	err := data.Delete(filters)
	if err != nil {
		return fmt.Errorf("error adding data %s", err)
	}

	return nil
}
