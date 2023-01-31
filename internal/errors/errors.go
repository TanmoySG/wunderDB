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
	UserDoesNotExistError = WdbError{
		ErrCode:        "userMissing",
		ErrMessage:     "user with id does not exist",
		HttpStatusCode: 404,
	}

	// auth error
	AuthenticatingUserDoesNotExist = WdbError{
		ErrCode:        "authenticatingUserNotFound",
		ErrMessage:     "auth failure, user trying to authenticate doesn't exist",
		HttpStatusCode: 404,
	}

	// Credential Error
	InvalidCredentialsError = WdbError{
		ErrCode:        "invalidCredentials",
		ErrMessage:     "username/password/token provided is not valid",
		HttpStatusCode: 401,
	}

	// Privilege Error
	PrivilegeUnauthorized = WdbError{
		ErrCode:        "permissionDenied",
		ErrMessage:     "permission denied",
		HttpStatusCode: 409,
	}

	// Other Errors
	SchemaValidationFailed = WdbError{
		ErrCode:        "schemaValidationError",
		ErrMessage:     "data failed schema validation",
		HttpStatusCode: 422,
	}

	EncodeDecodeError = WdbError{
		ErrCode:        encodeDecodeErrorCode,
		ErrMessage:     "error encoding/decoding data",
		HttpStatusCode: 406,
	}

	InvalidRoleError = WdbError{
		ErrCode:        "invalidRole",
		ErrMessage:     "role not valid",
		HttpStatusCode: 400,
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

	// Request validation Error
	ValidationError = WdbError{
		ErrCode:        "requestValidationFailed",
		ErrMessage:     "request validation failed",
		HttpStatusCode: 400,
	}
)
