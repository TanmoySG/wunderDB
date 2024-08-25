package wdbClient

import (
	c "github.com/TanmoySG/wunderDB/internal/collections"
	"github.com/TanmoySG/wunderDB/internal/identities"
	r "github.com/TanmoySG/wunderDB/internal/records"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"

	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddRecords(databaseId, collectionId model.Identifier, inputData interface{}) *er.WdbError {
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

	data := r.UseCollection(collection)

	recordId := identities.GenerateID()
	err := data.Add(model.Identifier(recordId), inputData)
	if err != nil {
		return err
	}

	wdb.updateParentMetadata(&databaseId, &collectionId)
	return nil
}

func (wdb wdbClient) GetRecords(databaseId, collectionId model.Identifier, filters interface{}) (map[model.Identifier]*model.Record, *er.WdbError) {
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

	data := r.UseCollection(collection)

	fetchedData, err := data.Read(filters)
	if err != nil {
		return nil, err
	}

	return fetchedData, nil
}

func (wdb wdbClient) QueryRecords(databaseId, collectionId model.Identifier, query string, mode r.QueryType) (interface{}, *er.WdbError) {
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

	records := r.UseCollection(collection)

	return records.Query(query, mode)
}

func (wdb wdbClient) UpdateRecords(databaseId, collectionId model.Identifier, updatedData, filters interface{}) *er.WdbError {
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

	data := r.UseCollection(collection)
	err := data.Update(updatedData, filters)
	if err != nil {
		return err
	}

	wdb.updateParentMetadata(&databaseId, &collectionId)
	return nil
}

func (wdb wdbClient) DeleteRecords(databaseId, collectionId model.Identifier, filters interface{}) *er.WdbError {
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

	data := r.UseCollection(collection)

	err := data.Delete(filters)
	if err != nil {
		return err
	}

	wdb.updateParentMetadata(&databaseId, &collectionId)
	return nil
}
