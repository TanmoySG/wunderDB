package model

type FilePath string
type Identifier string

type Paths map[string]FilePath

// type Metadata struct {
// 	Alias string     `json:"alias"`
// 	ID    Identifier `json:"id"`
// }

type Metadata map[string]interface{}
type User map[string]interface{}

type WDBFileSystem struct {
	Namespaces map[Identifier]Namespace `json:"namespaces"`
	Users      map[Identifier]User      `json:"users"`
}

type Namespace struct {
	Databases map[Identifier]Database `json:"databases"`
	Metadata  Metadata                `json:"metadata"`
	Access    *Access                 `json:"access,omitempty"` // Use only namespace level access control for Initial build v2
	Path      Paths                   `json:"path"`
}

type Database struct {
	Collections map[Identifier]Collection `json:"collections"`
	Metadata    Metadata                  `json:"metadata"`
	Access      *Access                   `json:"access,omitempty"` // not in scope for Initial Version of 2.0
	Path        Paths                     `json:"path"`
}

type Collection struct {
	Data     Data     `json:"data"`
	Metadata Metadata `json:"metadata"`
	Schema   Schema   `json:"schema"`
	Access   *Access  `json:"access,omitempty"` // not in scope for Initial Version of 2.0
	Path     Paths    `json:"path"`
}

type Data map[Identifier]interface{}
type Schema map[string]interface{}

// Need to Decide exact requirements for Access
// Access Control List
type Access struct {
	UserID         string
	AllowedActions []string
}
