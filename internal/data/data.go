package data

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/filter"
	"github.com/TanmoySG/wunderDB/internal/identities"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/utils/maps"
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

func (d Data) Get(filters interface{}) (map[model.Identifier]*model.Datum, error) {
	if filters != nil {
		f, err := filter.UseFilter(filters)
		if err != nil {
			return nil, fmt.Errorf("error reading data : %s", err)
		}

		filteredDate := f.Filter(d)
		return filteredDate, nil
	}
	return d, nil
}

func (d Data) Update(updatedData interface{}, filters interface{}) error {
	if filters != nil {
		f, err := filter.UseFilter(filters)
		if err != nil {
			return fmt.Errorf("error updating data : %s", err)
		}

		var iterError error

		f.Iterate(d, func(identifier model.Identifier, dataRow model.Datum) {

			mergableDataMaps := []map[string]interface{}{
				maps.Marshal(updatedData),
				dataRow.DataMap(),
			}

			data, err := maps.Merge(mergableDataMaps...)
			if err != nil {
				iterError = err
			}

			d[identifier] = &model.Datum{
				Data: data,
			}
		})

		if iterError != nil {
			return fmt.Errorf("error updating data : %s", iterError)
		}
		return nil
	}
	return fmt.Errorf("error updating data : filters missing")
}

func (d Data) Delete(updatedData interface{}, filters interface{}) error {
	if filters != nil {
		f, err := filter.UseFilter(filters)
		if err != nil {
			return fmt.Errorf("error reading data : %s", err)
		}

		f.Iterate(d, func(identifier model.Identifier, dataRow model.Datum) {
			delete(d, identifier)
		})
		return nil
	}
	return fmt.Errorf("error updating data : filters missing")
}
