package records

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/filter"
	"github.com/TanmoySG/wunderDB/internal/metadata"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/schema"
	"github.com/TanmoySG/wunderDB/pkg/utils/maps"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
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

type Records struct {
	Data       map[model.Identifier]*model.Record
	Schema     model.Schema
	PrimaryKey *model.Identifier
}

func UseCollection(collection *model.Collection) Records {
	return Records{
		Data:       collection.Records,
		Schema:     collection.Schema,
		PrimaryKey: collection.PrimaryKey,
	}
}

func (r Records) Add(recordId model.Identifier, data interface{}) *er.WdbError {
	s, err := schema.UseSchema(r.Schema)
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

	primaryKeyId := r.getPrimaryKey(recordId, &data)
	if _, ok := r.Data[model.Identifier(primaryKeyId)]; ok {
		return &er.RecordWithPrimaryKeyValueAlreadyExists
	}

	r.Data[primaryKeyId] = &model.Record{
		Identifier: model.Identifier(primaryKeyId),
		RecordId:   model.Identifier(recordId),
		Data:       data,
		Metadata:   metadata.New().BasicChangeMetadata(),
	}

	return nil
}

func (r Records) Read(filters interface{}) (map[model.Identifier]*model.Record, *er.WdbError) {
	if filters != nil {
		f, err := filter.UseFilter(filters)
		if err != nil {
			return nil, err
		}

		filteredData := f.Filter(*r.PrimaryKey, r.Data)
		if len(filteredData) == 0 {
			return nil, &er.RecordDoesNotExistsError
		}

		return filteredData, nil
	}

	return r.Data, nil
}

func (r Records) Update(updatedData interface{}, filters interface{}) *er.WdbError {
	var recordsCount int = 0
	var f *filter.Filter
	var err *er.WdbError

	if filters == nil {
		return &er.FilterMissingError

	}

	f, err = filter.UseFilter(filters)
	if err != nil {
		return err
	}

	f.Iterate(*r.PrimaryKey, r.Data, func(identifier *model.Identifier, dataRow *model.Record) {
		if identifier != nil && dataRow != nil {
			data, mergeErr := maps.Merge(maps.Marshal(updatedData), dataRow.DataMap())
			if mergeErr != nil {
				err = &er.DataEncodeDecodeError
			} else {
				schema, schemaErr := schema.UseSchema(r.Schema)
				if schemaErr != nil {
					err = schemaErr
				} else {
					isValid, schemaErr := schema.Validate(data)
					if schemaErr == nil && isValid {
						r.Data[*identifier].Data = &data
						r.Data[*identifier].Metadata = metadata.Use(r.Data[*identifier].Metadata).BasicChangeMetadata()
					}
					err = schemaErr
				}
			}
			recordsCount++
		}
	})

	if recordsCount == 0 && err == nil {
		return &er.RecordDoesNotExistsError
	}

	return err
}

func (r Records) Delete(filters interface{}) *er.WdbError {
	var recordsCount int = 0
	var f *filter.Filter
	var err *er.WdbError

	if filters != nil {
		f, err = filter.UseFilter(filters)
		if err != nil {
			return err
		}

		f.Iterate(*r.PrimaryKey, r.Data, func(identifier *model.Identifier, dataRow *model.Record) {
			if identifier != nil {
				delete(r.Data, *identifier)
				recordsCount++
			}
		})

		if recordsCount == 0 {
			return &er.RecordDoesNotExistsError
		}

		return nil
	}

	return &er.FilterMissingError
}

func (r Records) Query(query string, mode QueryType) (interface{}, *er.WdbError) {

	jsonData, err := json.Marshal(r.Data)
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
			return nil, &er.QueryResultProcessingError
		}

		var result interface{}
		err = json.Unmarshal(marshaledNode, &result)
		if err != nil {
			return nil, &er.QueryResultProcessingError
		}

		queryResults = append(queryResults, result)
	}

	return queryResults, nil
}

func (r Records) getPrimaryKey(recordId model.Identifier, data interface{}) model.Identifier {
	primaryKeyValue := recordId.String()

	if r.PrimaryKey.String() != defaultPrimaryKeyField {
		dataMap := maps.Marshal(data)

		// all primary key values are converted to string
		primaryKeyValue = fmt.Sprint(dataMap[r.PrimaryKey.String()])
	}

	return model.Identifier(primaryKeyValue)
}
