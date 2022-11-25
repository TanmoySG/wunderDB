package model

type Identifier string

type DataMap map[string]interface{}
type Metadata map[string]interface{}
type ExtraFields map[string]interface{}
type Authentication map[string]interface{}

type WDB struct {
	Namespaces map[Identifier]*Namespace `json:"namespaces"`
	Databases  map[Identifier]*Database  `json:"databases"`
	Users      map[Identifier]*User      `json:"users"`
}

type Namespace struct {
	Databases []Identifier           `json:"databases"`
	Metadata  Metadata               `json:"metadata"`
	Access    map[Identifier]*Access `json:"access,omitempty"` // Use only namespace level access control for Initial build v2
}

type Database struct {
	Collections map[Identifier]*Collection `json:"collections"`
	Metadata    Metadata                   `json:"metadata"`
	Access      map[Identifier]*Access     `json:"access,omitempty"` // not in scope for Initial Version of 2.0
}

type Collection struct {
	Data     map[Identifier]*Datum  `json:"data"`
	Metadata Metadata               `json:"metadata"`
	Schema   Schema                 `json:"schema"`
	Access   map[Identifier]*Access `json:"access,omitempty"` // not in scope for Initial Version of 2.0
}

type Datum struct {
	Data       interface{} `json:"data"`
	Metadata   Metadata    `json:"metadata"`
	Identifier Identifier  `json:"id"`
}

type Schema map[string]interface{}

// Need to Decide exact requirements for Access
// Access Control List - currently only implemented at Namespace Level
type Access struct {
	UserID         string   `json:"userId"`
	AllowedActions []string `json:"allowedActions"`
}

type User struct {
	UserID         Identifier
	Authentication Authentication
	Metadata       Metadata
}
