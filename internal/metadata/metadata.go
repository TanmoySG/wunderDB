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

func Use(m model.Metadata) metadata {
	return metadata(m)
}

func New() metadata {
	return metadata{}
}

func (m metadata) BasicChangeMetadata(optionalMetadata ...any) model.Metadata {
	currentTime := time.Now().UTC()
	updatedAtTimestamp := currentTime.Unix() // second precisioned timestamp

	createdAtTimestamp, createdAtExists := m[CreatedAt]
	if !createdAtExists {
		createdAtTimestamp = updatedAtTimestamp
	}

	return model.Metadata{
		CreatedAt: createdAtTimestamp,
		UpdatedAt: updatedAtTimestamp,
	}
}
