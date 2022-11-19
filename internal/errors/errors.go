package wdbErrors

type WdbError struct {
	ErrCode    string
	ErrMessage string

}

var (
	NamespaceAlreadyExistsError = WdbError{
		ErrCode:    "namespaceExists",
		ErrMessage: "namespace with ID already exists",
	}
	NamespaceDoesNotExistsError = WdbError{
		ErrCode:    "namespaceMissing",
		ErrMessage: "namespace with ID doesn't exist",
	}
)
