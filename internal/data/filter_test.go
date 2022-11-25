package data

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
	f, err := UseFilter(filter)

	assert.Equal(t, &Filter{
		Key:   "field",
		Value: "value",
	}, f)
	assert.NoError(t, err)

	testData := Data{
		"1": &model.Datum{
			Identifier: "1",
			Data: map[string]interface{}{
				"filed": "val",
				"num":   "1",
			},
		},
		"2": &model.Datum{
			Identifier: "1",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "2",
			},
		},
		"3": &model.Datum{
			Identifier: "1",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "3",
			},
		},
	}

	expectedData := Data{
		"2": &model.Datum{
			Identifier: "1",
			Data: map[string]interface{}{
				"field": "value",
				"num":   "2",
			},
		},
		"3": &model.Datum{
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
