package model

import (
	"github.com/TanmoySG/wunderDB/pkg/utils/maps"
)

func (i Identifier) String() string {
	return string(i)
}

func NewWDBInstance(namespaces map[Identifier]*Namespace, databases map[Identifier]*Database, users map[Identifier]*User) WDB {
	return WDB{
		Namespaces: namespaces,
		Databases:  databases,
		Users:      users,
	}
}

func (d Datum) DataMap() map[string]interface{} {
	return maps.Marshal(d.Data)
}
