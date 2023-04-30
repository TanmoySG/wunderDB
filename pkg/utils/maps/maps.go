package maps

import (
	"encoding/json"
)

func Marshal(data interface{}) map[string]interface{} {
	var dataMap map[string]interface{}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(dataBytes, &dataMap)
	if err != nil {
		return nil
	}

	return dataMap
}

func Merge(mapA, mapB map[string]interface{}) (map[string]interface{}, error) {
	for k, v := range mapA {
		mapB[k] = v
	}
	return mapB, nil
}
