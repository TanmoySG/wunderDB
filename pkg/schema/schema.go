package schema

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wunderDB/model"
	jsonschema "github.com/xeipuuv/gojsonschema"
)

type Schema struct {
	loadedSchema jsonschema.JSONLoader
}

func UseSchema(schema model.Schema) (*Schema, error) {
	marshaledSchemaJSON, err := json.Marshal(schema)
	if err != nil {
		return nil, fmt.Errorf("error with schema: %s", err)
	}
	loadedSchema := jsonschema.NewStringLoader(string(marshaledSchemaJSON))

	return &Schema{loadedSchema: loadedSchema}, nil
}

func (s Schema) Validate(data interface{}) (bool, error) {
	marshaledDataJSON, err := json.Marshal(data)
	if err != nil {
		return false, fmt.Errorf("error with data: %s", err)
	}
	loadedData := jsonschema.NewStringLoader(string(marshaledDataJSON))

	validity, _ := jsonschema.Validate(s.loadedSchema, loadedData)
	return validity.Valid(), nil
}
