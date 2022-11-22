package data

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/schema"
)

type Collection model.Collection

func UseCollection(collection model.Collection) Collection {
	return Collection(collection)
}

func (c Collection) AddData(data interface{}) error {
	s, err := schema.UseSchema(c.Schema)
	if err != nil {
		return fmt.Errorf("error adding data: %s", err)
	}

	isDataValid, err := s.Validate(data)
	if err != nil {
		return fmt.Errorf("error adding data: %s", err)
	}

	if !isDataValid {
		return fmt.Errorf("error adding data: data failed schema validation")
	}

	c.Data[""] = data
	return nil
}
