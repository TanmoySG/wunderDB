package data

import (
	"encoding/json"
	"fmt"
)

type Filter struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	// Filter operator - < / greater , > / lesser , = / equal , != / unequal
	// Operator string      `json:"operator"` // Future scope
}

func fieldExists(fieldKey string, dataMap map[string]interface{}) bool {
	if _, exists := dataMap[fieldKey]; exists {
		return fieldExist
	}

	return fieldDoesnotExist
}

func UseFilter(filter interface{}) (*Filter, error) {

	var dataFilter Filter

	filterJson, err := json.Marshal(filter)
	if err != nil {
		return nil, fmt.Errorf("error marshaling namespace file: %s", err)
	}

	err = json.Unmarshal(filterJson, &dataFilter)
	if err != nil {
		return nil, fmt.Errorf("error marshaling namespace file: %s", err)
	}

	return &dataFilter, nil
}

func (f Filter) Filter(data Data) Data {
	filteredData := make(Data)

	for identifier, dataRow := range data {
		dataMap := dataRow.DataMap()
		// TODO: move key check out of loop, or maybe not
		if fieldExists(f.Key, dataMap) {
			if dataMap[f.Key] == f.Value {
				filteredData[identifier] = dataRow
			}
		}
	}
	return filteredData
}
