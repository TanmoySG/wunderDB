package schema

import (
	"encoding/json"

	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	jsonschema "github.com/xeipuuv/gojsonschema"
)

type Schema struct {
	loadedSchema jsonschema.JSONLoader
}

var (
	// update default values for required schema fields
	// TODO: make this configurable at startup
	requiredSchemaFields = map[string]interface{}{
		"additionalProperties": false,
	}
)

func UseSchema(schema model.Schema) (*Schema, *er.WdbError) {
	marshaledSchemaJSON, err := json.Marshal(schema)
	if err != nil {
		return nil, &er.SchemaEncodeDecodeError
	}
	loadedSchema := jsonschema.NewStringLoader(string(marshaledSchemaJSON))

	return &Schema{loadedSchema: loadedSchema}, nil
}

func (s Schema) Validate(data interface{}) (bool, *er.WdbError) {
	marshaledDataJSON, err := json.Marshal(data)
	if err != nil {
		return false, &er.DataEncodeDecodeError
	}
	loadedData := jsonschema.NewStringLoader(string(marshaledDataJSON))

	validity, err := jsonschema.Validate(s.loadedSchema, loadedData)
	if err != nil {
		return false, er.SchemaValidationFailed.SetMessage(err.Error())
	}

	// TODO: return schema validation errors in response
	// &er.SchemaValidationFailed.SetMessage(validity.Errors())
	return validity.Valid(), nil
}

// adds default schema fields if not present, eg: additionalProperties [default: false]
func StandardizeSchema(schema model.Schema) model.Schema {

	// return schema if schema is empty
	if len(schema) == 0 {
		return schema
	}

	for field, fieldDefaultValue := range requiredSchemaFields {
		if _, ok := schema[field]; !ok {
			schema[field] = fieldDefaultValue
		}
	}

	return schema
}
