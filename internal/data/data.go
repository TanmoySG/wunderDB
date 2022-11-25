package data

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/identities"
	"github.com/TanmoySG/wunderDB/model"
)

type Data map[model.Identifier]*model.Datum

func UseCollection(collection model.Collection) Data {
	return Data(collection.Data)
}

func (d Data) Add(data interface{}) error {
	// s, err := schema.UseSchema(c.Schema)
	// if err != nil {
	// 	return fmt.Errorf("error adding data: %s", err)
	// }

	// isDataValid, err := s.Validate(data)
	// if err != nil {
	// 	return fmt.Errorf("error adding data: %s", err)
	// }

	// if !isDataValid {
	// 	return fmt.Errorf("error adding data: data failed schema validation")
	// }

	dataKey := identities.GenerateID()

	d[model.Identifier(dataKey)] = &model.Datum{
		Identifier: model.Identifier(dataKey),
		Data:       data,
		Metadata:   model.Metadata{},
	}
	return nil
}

func (d Data) Get(filter interface{}) (Data, error) {
	if filter != nil {
		f, err := UseFilter(filter)
		if err != nil {
			return nil, fmt.Errorf("error reading data : %s", err)
		}

		filteredDate := f.Filter(d)
		return filteredDate, nil
	}
	return d, nil
}
