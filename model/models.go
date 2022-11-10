package model

type FilePath string
type Identifier string

type WDBFileSystem struct {
	Namespaces map[Identifier]Namespace
	Users      map[Identifier]User
}

type Namespace struct {
	NamespaceAlias string
	NamespaceID    Identifier
	NamespacePath  FilePath
}

type User map[string]interface{}
