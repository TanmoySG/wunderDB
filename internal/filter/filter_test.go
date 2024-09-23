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

	testData := map[model.Identifier]*model.Record{
		"1": {
			Identifier: "1",
			Data: map[string]interface{}{
				"field": "val",
				"num":   "1",
			},
		},
		"2": {
			Identifier: "2",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "2",
			},
		},
		"3": {
			Identifier: "3",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "3",
			},
		},
	}

	expectedData := map[model.Identifier]*model.Record{
		"2": {
			Identifier: "2",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "2",
			},
		},
		"3": {
			Identifier: "3",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "3",
			},
		},
	}

	data := f.Filter("num", testData)
	assert.Equal(t, &expectedData, &data)
}
