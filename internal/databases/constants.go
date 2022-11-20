package databases

import ()

const (
	DatabaseErrorFormat = "[%s] %s : %s"

	databaseExists       = true
	databaseDoesNotExist = false
)

// var (
// 	NamespaceAlreadyExistsError = fmt.Errorf(NamespaceErrorFormat, er.NamespaceAlreadyExistsError.ErrCode, "error creating namespace", er.NamespaceAlreadyExistsError.ErrMessage)
// 	NamespaceDoesNotExistsError = fmt.Errorf(NamespaceErrorFormat, er.NamespaceDoesNotExistsError.ErrCode, "error deleting namespace", er.NamespaceDoesNotExistsError.ErrMessage)
// )
