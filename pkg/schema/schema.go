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

	validity, _ := jsonschema.Validate(s.loadedSchema, loadedData)
	return validity.Valid(), nil
}
