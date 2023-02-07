package filter

import (
	"testing"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {

	filter := map[string]interface{}{
		"key":   "field",
		"value": "value",
	}
	f, _ := UseFilter(filter)

	assert.Equal(t, &Filter{
		Key:   "field",
		Value: "value",
	}, f)

	testData := map[model.Identifier]*model.Datum{
		"1": {
			Identifier: "1",
			Data: map[string]interface{}{
				"filed": "val",
				"num":   "1",
			},
		},
		"2": {
			Identifier: "1",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "2",
			},
		},
		"3": {
			Identifier: "1",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "3",
			},
		},
	}

	expectedData := map[model.Identifier]*model.Datum{
		"2": {
			Identifier: "1",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "2",
			},
		},
		"3": {
			Identifier: "1",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "3",
			},
		},
	}

	data := f.Filter(testData)
	assert.Equal(t, &expectedData, &data)
}
