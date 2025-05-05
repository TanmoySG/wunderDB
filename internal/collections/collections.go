package collections

import (
	"github.com/TanmoySG/wunderDB/internal/metadata"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/model/redacted"
	s "github.com/TanmoySG/wunderDB/pkg/schema"
	wdbErrors "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

const (
	collectionExists       = true
	collectionDoesNotExist = false

	schemaFieldRequired   = "required"
	schemaFieldProperties = "properties"

	defaultPrimaryKeyField = "recordId"
)

type Collections map[model.Identifier]*model.Collection

func UseDatabase(database *model.Database) Collections {
	return Collections(database.Collections)
}

func (c Collections) CheckIfExists(collectionID model.Identifier) (bool, *model.Collection) {
	collection, collExists := c[collectionID]
	if collExists {
		return collectionExists, collection
	} else {
		return collectionDoesNotExist, collection
	}
}

func (c Collections) CreateCollection(collectionID model.Identifier, schema model.Schema, primaryKey *model.Identifier) *wdbErrors.WdbError {
	// get primary key field value
	primaryKeyField, err := c.getPrimaryKey(primaryKey, schema)
	if err != nil {
		return err
	}

	// standardize schema: add schema fields if not present
	schema = s.StandardizeSchema(schema)

	c[collectionID] = &model.Collection{
		Records:    map[model.Identifier]*model.Record{},
		Schema:     schema,
		Metadata:   metadata.New().BasicChangeMetadata(),
		PrimaryKey: &primaryKeyField,
	}

	return nil
}

func (c Collections) GetCollection(collectionID model.Identifier) *redacted.RedactedC {
	collection := c[collectionID]

	if collection == nil {
		return nil
	}

	redactedCollection := redacted.RedactedC{
		PrimaryKey: collection.PrimaryKey,
		Metadata:   collection.Metadata,
		Schema:     collection.Schema,
	}

	return &redactedCollection
}

func (c Collections) DeleteCollection(collectionID model.Identifier) {
	delete(c, collectionID)
}

func (c Collections) UpdateMetadata(collectionID model.Identifier) {
	c[collectionID].Metadata = metadata.Use(c[collectionID].Metadata).BasicChangeMetadata()
}

func (Collections) getPrimaryKey(primaryKey *model.Identifier, schema model.Schema) (model.Identifier, *wdbErrors.WdbError) {
	// default primary key field is "recordId"
	var primaryKeyField model.Identifier = model.Identifier(defaultPrimaryKeyField)

	if primaryKey != nil && schema[schemaFieldProperties] != nil && schema[schemaFieldRequired] != nil {
		// check if primary key is in "properties" field of JSON Schema
		_, pkeyExistsInProperties := schema[schemaFieldProperties].(map[string]interface{})[primaryKey.String()]

		// check if primary key is in "required" field of JSON Schema
		pkeyIsRequired := false
		for _, requiredField := range schema[schemaFieldRequired].([]interface{}) {
			if primaryKey.String() == requiredField.(string) {
				pkeyIsRequired = true
			}
		}

		// if primary key not in "properties" or "required"
		// then return primaryKey schema mismatch error
		if !pkeyExistsInProperties || !pkeyIsRequired {
			return "", &wdbErrors.PrimaryKeyNotInSchemaError
		}

		// set primaryKeyField as primary key mentioned
		primaryKeyField = model.Identifier(primaryKey.String())
	}

	return primaryKeyField, nil
}
