package data

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/filter"
	"github.com/TanmoySG/wunderDB/internal/identities"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/schema"
	"github.com/TanmoySG/wunderDB/pkg/utils/maps"
)

type Data struct {
	Data   map[model.Identifier]*model.Datum
	Schema model.Schema
}

func UseCollection(collection model.Collection) Data {
	return Data{
		Data:   collection.Data,
		Schema: collection.Schema,
	}
}

func (d Data) Add(data interface{}) error {
	s, err := schema.UseSchema(d.Schema)
	if err != nil {
		return fmt.Errorf("error adding data: %s", err)
	}

	isValid, err := s.Validate(data)
	if err != nil {
		return fmt.Errorf("error adding data: %s", err)
	}

	if isValid {
		dataKey := identities.GenerateID()
		d.Data[model.Identifier(dataKey)] = &model.Datum{
			Identifier: model.Identifier(dataKey),
			Data:       data,
			Metadata:   model.Metadata{},
		}
		return nil
	} else {
		return fmt.Errorf("error adding data: data failed schema validation")
	}
}

func (d Data) Read(filters interface{}) (map[model.Identifier]*model.Datum, error) {
	if filters != nil {
		f, err := filter.UseFilter(filters)
		if err != nil {
			return nil, fmt.Errorf("error reading data : %s", err)
		}

		filteredDate := f.Filter(d.Data)
		return filteredDate, nil
	}
	return d.Data, nil
}

func (d Data) Update(updatedData interface{}, filters interface{}) error {
	if filters != nil {
		f, err := filter.UseFilter(filters)
		if err != nil {
			return fmt.Errorf("error updating data : %s", err)
		}

		var iterError error

		f.Iterate(d.Data, func(identifier model.Identifier, dataRow model.Datum) {

			mergableDataMaps := []map[string]interface{}{
				maps.Marshal(updatedData),
				dataRow.DataMap(),
			}

			data, err := maps.Merge(mergableDataMaps...)
			if err != nil {
				iterError = err
			} else {
				schema, err := schema.UseSchema(d.Schema)

				if err != nil {
					iterError = err
				} else {
					isValid, err := schema.Validate(data)
					if err == nil && isValid {
						d.Data[identifier] = &model.Datum{
							Data: data,
						}
					}
					iterError = err
				}
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

		f.Iterate(d.Data, func(identifier model.Identifier, dataRow model.Datum) {
			delete(d.Data, identifier)
		})
		return nil
	}
	return fmt.Errorf("error updating data : filters missing")
}
