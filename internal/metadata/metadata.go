package metadata

import (
	"time"

	"github.com/TanmoySG/wunderDB/model"
)

type metadata model.Metadata

const (
	CreatedAt = "created_at"
	UpdatedAt = "updated_at"
)

// type changeTimeMetadata struct {
// 	CreatedAt string `json:"created_at"`
// 	UpdatedAt string `json:"updated_at"`
// }

func BasicCreateTimestampMetadata(optionalMetadata ...any) metadata {
	currentTime := time.Now()
	currentTimestamp := currentTime.Unix() // second precisioned timestamp

	return metadata{
		CreatedAt: currentTimestamp,
		UpdatedAt: currentTimestamp,
	}
}

func (m metadata) BasicUpdateTimestampMetadata(optionalMetadata ...any) metadata {
	currentTime := time.Now()
	currentTimestamp := currentTime.Unix() // second precisioned timestamp

	createdAtTimestamp, createdAtExists := m[CreatedAt]
	if !createdAtExists {
		createdAtTimestamp = currentTimestamp
	}

	updatedAtTimestamp := currentTimestamp

	return metadata{
		CreatedAt: createdAtTimestamp,
		UpdatedAt: updatedAtTimestamp,
	}
}

func (m metadata) Commit() model.Metadata {
	return model.Metadata(m)
}
