package wdbClient

import (
	c "github.com/TanmoySG/wunderDB/internal/collections"
	d "github.com/TanmoySG/wunderDB/internal/data"
	"github.com/TanmoySG/wunderDB/internal/identities"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"

	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddData(databaseId, collectionId model.Identifier, inputData interface{}) *er.WdbError {
	if !wdb.safeName.Check(databaseId.String()) {
		return &er.DatabaseNameFormatError
	}

	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	if !wdb.safeName.Check(collectionId.String()) {
		return &er.CollectionNameFormatError
	}

	collections := c.UseDatabase(database)
	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return &er.CollectionDoesNotExistsError
	}

	collection.Lock()
	defer collection.Unlock()

	data := d.UseCollection(collection)

	recordId := identities.GenerateID()
	err := data.Add(model.Identifier(recordId), inputData)
	if err != nil {
		return err
	}

	wdb.updateParentMetadata(&databaseId, &collectionId)
	return nil
}

func (wdb wdbClient) GetData(databaseId, collectionId model.Identifier, filters interface{}) (map[model.Identifier]*model.Record, *er.WdbError) {
	if !wdb.safeName.Check(databaseId.String()) {
		return nil, &er.DatabaseNameFormatError
	}

	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return nil, &er.DatabaseDoesNotExistsError
	}

	if !wdb.safeName.Check(collectionId.String()) {
		return nil, &er.CollectionNameFormatError
	}

	collections := c.UseDatabase(database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return nil, &er.CollectionDoesNotExistsError
	}

	data := d.UseCollection(collection)

	fetchedData, err := data.Read(filters)
	if err != nil {
		return nil, err
	}

	return fetchedData, nil
}

func (wdb wdbClient) QueryData(databaseId, collectionId model.Identifier, query string, mode d.QueryType) (interface{}, *er.WdbError) {
	if !wdb.safeName.Check(databaseId.String()) {
		return nil, &er.DatabaseNameFormatError
	}

	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return nil, &er.DatabaseDoesNotExistsError
	}

	if !wdb.safeName.Check(collectionId.String()) {
		return nil, &er.CollectionNameFormatError
	}

	collections := c.UseDatabase(database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return nil, &er.CollectionDoesNotExistsError
	}

	data := d.UseCollection(collection)

	return data.Query(query, mode)
}

func (wdb wdbClient) UpdateData(databaseId, collectionId model.Identifier, updatedData, filters interface{}) *er.WdbError {
	if !wdb.safeName.Check(databaseId.String()) {
		return &er.DatabaseNameFormatError
	}

	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	if !wdb.safeName.Check(collectionId.String()) {
		return &er.CollectionNameFormatError
	}

	collections := c.UseDatabase(database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return &er.CollectionDoesNotExistsError
	}

	collection.Lock()
	defer collection.Unlock()

	data := d.UseCollection(collection)
	err := data.Update(updatedData, filters)
	if err != nil {
		return err
	}

	wdb.updateParentMetadata(&databaseId, &collectionId)
	return nil
}

func (wdb wdbClient) DeleteData(databaseId, collectionId model.Identifier, filters interface{}) *er.WdbError {
	if !wdb.safeName.Check(databaseId.String()) {
		return &er.DatabaseNameFormatError
	}

	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	if !wdb.safeName.Check(collectionId.String()) {
		return &er.CollectionNameFormatError
	}

	collections := c.UseDatabase(database)

	collectionExists, collection := collections.CheckIfExists(collectionId)
	if !collectionExists {
		return &er.CollectionDoesNotExistsError
	}

	collection.Lock()
	defer collection.Unlock()

	data := d.UseCollection(collection)

	err := data.Delete(filters)
	if err != nil {
		return err
	}

	wdb.updateParentMetadata(&databaseId, &collectionId)
	return nil
}
