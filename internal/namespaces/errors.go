package namespaces

type errors struct {
	errCode    string
	errMessage string
}

const (
	NamespaceErrorFormat = "[%s] %s : %s"
)

var (
	NamespaceAlreadyExistsError = errors{
		errCode:    "namespaceExists",
		errMessage: "namespace with ID already exists",
	}
	NamespaceDoesnotExistsError = errors{
		errCode:    "namespaceMissing",
		errMessage: "namespace with ID doesn't exist",
	}
)
