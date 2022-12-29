package wdbErrors

type WdbError struct {
	ErrCode        string
	ErrMessage     string
	HttpStatusCode int
}

const (
	encodeDecodeErrorCode = "encodeDecodeError"
)

var (
	NamespaceAlreadyExistsError = WdbError{
		ErrCode:        "namespaceExists",
		ErrMessage:     "namespace with ID already exists",
		HttpStatusCode: 409,
	}
	NamespaceDoesNotExistsError = WdbError{
		ErrCode:        "namespaceMissing",
		ErrMessage:     "namespace with ID doesn't exist",
		HttpStatusCode: 404,
	}

	// Database Errors
	DatabaseAlreadyExistsError = WdbError{
		ErrCode:        "databaseExists",
		ErrMessage:     "database with ID already exists",
		HttpStatusCode: 409,
	}
	DatabaseDoesNotExistsError = WdbError{
		ErrCode:        "databaseMissing",
		ErrMessage:     "database with ID doesn't exist",
		HttpStatusCode: 404,
	}

	// Collection Errors
	CollectionAlreadyExistsError = WdbError{
		ErrCode:        "collectionExists",
		ErrMessage:     "collection with ID already exists",
		HttpStatusCode: 409,
	}
	CollectionDoesNotExistsError = WdbError{
		ErrCode:        "collectionMissing",
		ErrMessage:     "collection with ID doesn't exist",
		HttpStatusCode: 404,
	}

	// Missing Error
	FilterMissingError = WdbError{
		ErrCode:        "filterMissing",
		ErrMessage:     "filters missing",
		HttpStatusCode: 404,
	}

	// Role Errors
	RoleAlreadyExistsError = WdbError{
		ErrCode:        "roleExists",
		ErrMessage:     "role with name already exists",
		HttpStatusCode: 409,
	}

	// User Errors
	UserAlreadyExistsError = WdbError{
		ErrCode:        "userExists",
		ErrMessage:     "user with id already exists",
		HttpStatusCode: 409,
	}
	UserAlreadyDoesNotExistError = WdbError{
		ErrCode:        "userMissing",
		ErrMessage:     "user with id already exists",
		HttpStatusCode: 404,
	}

	// Other Errors
	SchemaValidationFailed = WdbError{
		ErrCode:        "schemaValidationError",
		ErrMessage:     "data failed schema validation",
		HttpStatusCode: 422,
	}

	// Encode/Decode Error
	DataEncodeDecodeError = WdbError{
		ErrCode:        encodeDecodeErrorCode,
		ErrMessage:     "error encoding/decoding data",
		HttpStatusCode: 406,
	}
	SchemaEncodeDecodeError = WdbError{
		ErrCode:        encodeDecodeErrorCode,
		ErrMessage:     "error encoding/decoding schema",
		HttpStatusCode: 406,
	}
	FilterEncodeDecodeError = WdbError{
		ErrCode:        encodeDecodeErrorCode,
		ErrMessage:     "error encoding/decoding filter",
		HttpStatusCode: 406,
	}
)
