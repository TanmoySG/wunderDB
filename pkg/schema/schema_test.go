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
	js, _ := UseSchema(testJSONSchema)

	testData := map[string]interface{}{
		"001": "ia",
	}

	isValid, _ := js.Validate(testData)
	assert.Equal(t, true, isValid)
}
