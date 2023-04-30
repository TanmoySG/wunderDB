package identities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateID(t *testing.T) {
	got := GenerateID()
	assert.NotNil(t, got)
	assert.NotEqual(t, "", got)
}
