package model

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
