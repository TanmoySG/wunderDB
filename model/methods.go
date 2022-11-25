package model

import (
	"encoding/json"
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
	var dataMap DataMap

	dataBytes, err := json.Marshal(d.Data)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(dataBytes, &dataMap)
	if err != nil {
		return nil
	}

	return dataMap
}
