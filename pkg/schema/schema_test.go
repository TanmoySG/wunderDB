package schema

import (
	"testing"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/stretchr/testify/assert"
)

func Test_Validate(t *testing.T) {

	testJSONSchema := model.Schema{
		"type": "object",
	}
	js, err := UseSchema(testJSONSchema)

	assert.NoError(t, err)

	testData := model.Data{
		"001": "ia",
	}

	isValid, err := js.Validate(testData)
	assert.NoError(t, err)
	assert.Equal(t, true, isValid)
}
