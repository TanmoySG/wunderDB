package data

import (
	"github.com/TanmoySG/wunderDB/internal/identities"
	"github.com/TanmoySG/wunderDB/model"
)

type Data map[model.Identifier]*model.Datum

func UseCollection(collection model.Collection) Data {
	return Data(collection.Data)
}

func (d Data) AddData(data interface{}) error {
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

func (d Data) GetData(markers *interface{}) (Data, error) {
	return d, nil
}
