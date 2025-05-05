package redacted

import "github.com/TanmoySG/wunderDB/model"

type RedactedD struct {
	Collections []model.Identifier                 `json:"collections"` // TODO: maybe this can be removed to return just the metadata and access
	Metadata    model.Metadata                     `json:"metadata"`
	Access      map[model.Identifier]*model.Access `json:"access,omitempty"`
}

type RedactedC struct {
	Metadata   model.Metadata    `json:"metadata"`
	Schema     model.Schema      `json:"schema"`
	PrimaryKey *model.Identifier `json:"primaryKey,omitempty"`
}
