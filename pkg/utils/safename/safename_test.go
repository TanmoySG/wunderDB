package safename

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UsePattern(t *testing.T) {

	testData := []struct {
		inputPattern  string
		expectedError error
	}{
		{
			inputPattern:  "[a-zA-Z0-9-_.]+.*[^-_.]$",
			expectedError: nil,
		},
		{
			inputPattern:  "!![]",
			expectedError: fmt.Errorf("failed to compile pattern: error parsing regexp: missing closing ]: `[]`"),
		},
		{
			inputPattern:  "^[a-zA-Z0-9-_.]+$",
			expectedError: nil,
		},
	}

	for _, tc := range testData {
		_, err := UsePattern(tc.inputPattern)
		assert.Equal(t, tc.expectedError, err)
	}

}

func Test_Check(t *testing.T) {

	usePattern := "(?misU)^[a-zA-Z0-9-_.@]+$"
	patternClient, err := UsePattern(usePattern)
	assert.Nil(t, err)

	testData := []struct {
		inputString     string
		expectedIsValid bool
	}{
		{
			inputString:     "abc_2",
			expectedIsValid: true,
		},
		{
			inputString:     "abc1.3",
			expectedIsValid: true,
		},
		{
			inputString:     "abc_",
			expectedIsValid: true,
		},
		{
			inputString:     "john@doe.com",
			expectedIsValid: true,
		},
		{
			inputString:     "a b",
			expectedIsValid: false,
		},
		{
			inputString:     "a-b",
			expectedIsValid: true,
		},
		{
			inputString:     "abc/",
			expectedIsValid: false,
		},
	}

	for _, tc := range testData {
		isValid := patternClient.Check(tc.inputString)
		if !assert.Equal(t, tc.expectedIsValid, isValid) {
			log.Printf("test failed for input string [%s], expected [%v], got [%v]", tc.inputString, isValid, tc.expectedIsValid)
		}
	}

}
