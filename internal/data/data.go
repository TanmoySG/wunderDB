package data

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/filter"
	"github.com/TanmoySG/wunderDB/internal/metadata"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/schema"
	"github.com/TanmoySG/wunderDB/pkg/utils/maps"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	wdbErrors "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/spyzhov/ajson"
)

const (
	defaultPrimaryKeyField = "recordId"
)

var (
	JsonPathQuery QueryType = "jsonpath"
	EvaluateQuery QueryType = "evaluate"
)

type QueryType string

type Data struct {
	Data       map[model.Identifier]*model.Record
	Schema     model.Schema
	PrimaryKey *model.Identifier
}

func UseCollection(collection *model.Collection) Data {
	return Data{
		Data:       collection.Data,
		Schema:     collection.Schema,
		PrimaryKey: collection.PrimaryKey,
	}
}

func (d Data) Add(recordId model.Identifier, data interface{}) *er.WdbError {
	s, err := schema.UseSchema(d.Schema)
	if err != nil {
		return err
	}

	isValid, err := s.Validate(data)
	if err != nil {
		return err
	}

	if !isValid {
		return &er.SchemaValidationFailed
	}

	primaryKeyId := d.getPrimaryKey(recordId, &data)
	if _, ok := d.Data[model.Identifier(primaryKeyId)]; ok {
		return &er.RecordWithPrimaryKeyValueAlreadyExists
	}

	d.Data[primaryKeyId] = &model.Record{
		Identifier: model.Identifier(primaryKeyId),
		RecordId:   model.Identifier(recordId),
		Data:       data,
		Metadata:   metadata.New().BasicChangeMetadata(),
	}

	return nil
}

func (d Data) Read(filters interface{}) (map[model.Identifier]*model.Record, *er.WdbError) {
	if filters != nil {
		f, err := filter.UseFilter(filters)
		if err != nil {
			return nil, err
		}

		filteredData := f.Filter(*d.PrimaryKey, d.Data)
		return filteredData, nil
	}

	return d.Data, nil
}

func (d Data) Update(updatedData interface{}, filters interface{}) *er.WdbError {
	if filters == nil {
		return &er.FilterMissingError

	}

	f, err := filter.UseFilter(filters)
	if err != nil {
		return err
	}

	var iterError *er.WdbError

	f.Iterate(*d.PrimaryKey, d.Data, func(identifier model.Identifier, dataRow model.Record) {

		data, err := maps.Merge(maps.Marshal(updatedData), dataRow.DataMap())
		if err != nil {
			iterError = &er.DataEncodeDecodeError
		} else {
			schema, err := schema.UseSchema(d.Schema)

			if err != nil {
				iterError = err
			} else {
				isValid, err := schema.Validate(data)
				if err == nil && isValid {
					d.Data[identifier].Data = &data
					d.Data[identifier].Metadata = metadata.Use(d.Data[identifier].Metadata).BasicChangeMetadata()
				}
				iterError = err
			}
		}
	})

	if iterError != nil {
		return iterError
	}
	return nil
}

func (d Data) Delete(filters interface{}) *er.WdbError {
	if filters != nil {
		f, err := filter.UseFilter(filters)
		if err != nil {
			return err
		}

		f.Iterate(*d.PrimaryKey, d.Data, func(identifier model.Identifier, dataRow model.Record) {
			delete(d.Data, identifier)
		})
		return nil
	}
	return &er.FilterMissingError
}

func (d Data) Query(query string, mode QueryType) (interface{}, *er.WdbError) {

	jsonData, err := json.Marshal(d.Data)
	if err != nil {
		return nil, nil
	}

	var queryResultNodes []*ajson.Node
	var queryResults []interface{}

	root, err := ajson.Unmarshal(jsonData)
	if err != nil {
		return nil, nil
	}

	switch mode {
	case JsonPathQuery:
		jpqResult, err := root.JSONPath(query)
		if err != nil {
			return nil, er.JSONPathQueryError.SetMessage(err.Error())
		}
		queryResultNodes = jpqResult
	case EvaluateQuery:
		evqResult, err := ajson.Eval(root, query)
		if err != nil {
			return nil, er.QueryExecutionFailed.SetMessage(err.Error())
		}

		queryResultNodes = []*ajson.Node{evqResult}
	}

	for _, node := range queryResultNodes {
		marshaledNode, err := ajson.Marshal(node)
		if err != nil {
			return nil, &wdbErrors.QueryResultProcessingError
		}

		var result interface{}
		err = json.Unmarshal(marshaledNode, &result)
		if err != nil {
			return nil, &wdbErrors.QueryResultProcessingError
		}

		queryResults = append(queryResults, result)
	}

	return queryResults, nil
}

func (d Data) getPrimaryKey(recordId model.Identifier, data interface{}) model.Identifier {
	primaryKeyValue := recordId.String()

	if d.PrimaryKey.String() != defaultPrimaryKeyField {
		dataMap := maps.Marshal(data)

		// all primary key values are converted to string
		primaryKeyValue = fmt.Sprint(dataMap[d.PrimaryKey.String()])
	}

	return model.Identifier(primaryKeyValue)
}
