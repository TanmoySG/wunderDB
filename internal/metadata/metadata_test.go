package metadata

import (
	"testing"
	"time"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/stretchr/testify/assert"
)

func Test_metadata_BasicChangeMetadata(t *testing.T) {
	createTimestamp := time.Now().UTC().Unix()
	// empty Metadata
	wantM := model.Metadata{
		CreatedAt: createTimestamp,
		UpdatedAt: createTimestamp,
	}

	// initializing new metadata, for new metadata
	// - createdAt = updatedAt
	// - createdAt = updatedAt = time.Now / createdTimestamp
	m := New().BasicChangeMetadata()
	assert.Equal(t, m[CreatedAt], createTimestamp)
	assert.Equal(t, m[UpdatedAt], createTimestamp)
	assert.Equal(t, m[CreatedAt], m[UpdatedAt])
	assert.Equal(t, wantM, m)

	// adding sleep so that create and update timestamps are different
	time.Sleep(1 * time.Second)
	updateTimestamp := createTimestamp + 1

	// update want m
	wantM[UpdatedAt] = updateTimestamp

	// Update metadata, using Use()
	// createdAt = createTimestamp
	// updatedAt = updateTimestamp
	m = Use(m).BasicChangeMetadata()
	assert.Equal(t, m[CreatedAt], createTimestamp)
	assert.Equal(t, m[UpdatedAt], updateTimestamp)
	assert.NotEqual(t, m[CreatedAt], m[UpdatedAt])
	assert.Equal(t, wantM, m)

}

func Test_Use_and_New(t *testing.T) {

	// test New()
	want := metadata{}
	got := New()
	assert.Equal(t, want, got)

	// tests for Use()
	tests := []struct {
		metadata model.Metadata
		want     metadata
	}{
		{
			metadata: model.Metadata{},
			want:     metadata{},
		},
		{
			metadata: model.Metadata{CreatedAt: time.Now().UTC().Unix(), UpdatedAt: time.Now().UTC().Unix()},
			want:     metadata{CreatedAt: time.Now().UTC().Unix(), UpdatedAt: time.Now().UTC().Unix()},
		},
		{
			metadata: model.Metadata{"randomMetadataKey": "randomMetadataValue"},
			want:     metadata{"randomMetadataKey": "randomMetadataValue"},
		},
	}

	for _, tc := range tests {
		got := Use(tc.metadata)
		assert.Equal(t, tc.want, got)
	}
}
