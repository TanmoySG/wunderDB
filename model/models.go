package model

type Identifier string

type DataMap map[string]interface{}
type Metadata map[string]interface{}
type ExtraFields map[string]interface{}
type Privileges map[string]bool
type Schema map[string]interface{}

type WDB struct {
	Configurations Configurations            `json:"configurations"`
	Namespaces     map[Identifier]*Namespace `json:"namespaces"`
	Databases      map[Identifier]*Database  `json:"databases"`
	Users          map[Identifier]*User      `json:"users"`
	Roles          map[Identifier]*Role      `json:"roles"`
}

type Configurations struct {
	Admin *Identifier `json:"admin"`
}

type Namespace struct {
	Databases []Identifier           `json:"databases"`
	Metadata  Metadata               `json:"metadata"`
	Access    map[Identifier]*Access `json:"access,omitempty"`
}

type Database struct {
	Collections map[Identifier]*Collection `json:"collections"`
	Metadata    Metadata                   `json:"metadata"`
	Access      map[Identifier]*Access     `json:"access,omitempty"`
}

type Collection struct {
	Data     map[Identifier]*Datum  `json:"data"`
	Metadata Metadata               `json:"metadata"`
	Schema   Schema                 `json:"schema"`
	Access   map[Identifier]*Access `json:"access,omitempty"`
}

type Datum struct {
	Data       interface{} `json:"data"`
	Metadata   Metadata    `json:"metadata"`
	Identifier Identifier  `json:"id"`
}

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

type Permissions struct {
	Role Identifier `json:"role"`
	On   *Entities  `json:"on,omitempty"`
}

type Role struct {
	RoleID Identifier `json:"roleId"`
	Grants Grants     `json:"grants"`
}

type Grants struct {
	GlobalPrivileges     *Privileges `json:"globalPrivileges,omitempty"`
	DatabasePrivileges   *Privileges `json:"databasePrivileges,omitempty"`
	CollectionPrivileges *Privileges `json:"collectionPrivileges,omitempty"`
}

type Entities struct {
	Databases   *string `json:"databases,omitempty"`
	Collections *string `json:"collections,omitempty"`
}

type Authentication struct {
	HashedSecret     string `json:"hashedSecret"`
	HashingAlgorithm string `json:"hashingAlgorithm"` // md5, sha1 or sha256
}
