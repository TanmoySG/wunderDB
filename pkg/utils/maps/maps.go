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

func Merge(maps ...map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result, nil
}
