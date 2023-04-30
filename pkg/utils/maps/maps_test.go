package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	structTestData = struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John Doe",
		Age:  69,
	}

	structTestDataMap = map[string]interface{}{
		"name": "John Doe",
		"age":  float64(69),
	}
)

func Test_Marshal(t *testing.T) {

	testCases := []struct {
		inputData   interface{}
		expectedMap map[string]interface{}
	}{
		{
			inputData:   structTestData,
			expectedMap: structTestDataMap,
		},
		{
			inputData:   1, // should cause json.Unmarshal error
			expectedMap: nil,
		},
		{
			inputData:   make(chan int), // should cause json.Marshal error
			expectedMap: nil,
		},
	}

	for _, tc := range testCases {
		got := Marshal(tc.inputData)
		assert.Equal(t, tc.expectedMap, got)
	}
}

func Test_Merge(t *testing.T) {
	testMap1 := map[string]interface{}{
		"name": "John Doe",
		"age":  float64(69),
	}

	testMap2 := map[string]interface{}{
		"country": "india",
		"address": map[string]interface{}{
			"city":    "guwahati",
			"state":   "assam",
			"pincode": 781016,
		},
	}

	expectedMergedMap := map[string]interface{}{
		"name":    "John Doe",
		"age":     float64(69),
		"country": "india",
		"address": map[string]interface{}{
			"city":    "guwahati",
			"state":   "assam",
			"pincode": 781016,
		},
	}

	gotMergedMap, gotError := Merge(testMap1, testMap2)
	assert.Equal(t, nil, gotError)
	assert.Equal(t, expectedMergedMap, gotMergedMap)
}
