package model

type Identifier string

type DataMap map[string]interface{}
type Metadata map[string]interface{}
type ExtraFields map[string]interface{}
type Privileges map[string]bool
type Schema map[string]interface{}

type WDB struct {
	Namespaces map[Identifier]*Namespace `json:"namespaces"`
	Databases  map[Identifier]*Database  `json:"databases"`
	Users      map[Identifier]*User      `json:"users"`
	Roles      map[Identifier]*Role      `json:"roles"`
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

// Need to Decide exact requirements for Access
// Access Control List - currently only implemented at Namespace Level
type Access struct {
	UserID         string   `json:"userId"`
	AllowedActions []string `json:"allowedActions"`
}

type User struct {
	UserID         Identifier     `json:"userId"`
	Authentication Authentication `json:"authentication"`
	Metadata       Metadata       `json:"metadata"`
	Permissions    []Permissions  `json:"permissions"`
}

type Role struct {
	RoleID     Identifier `json:"roleID"`
	Privileges Privileges `json:"privileges"`
}

type Permissions struct {
	Role string           `json:"roles"`
	On   *AllowedEntities `json:"on,omitempty"`
}

type AllowedEntities struct {
	Databases   *[]string `json:"databases"`
	Collections *[]string `json:"collections"`
}

type Authentication struct {
	HashedSecret     string `json:"hashedSecret"`
	HashingAlgorithm string `json:"hashingAlgorithm"` // md5, sha1 or sha256
}
