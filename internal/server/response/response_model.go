package response

const (
	StatusSuccess = "success"
	StatusFailure = "failure"
)

type ApiResponse struct {
	Response       Response
	HttpStatusCode int
}

type Response struct {
	Action string      `json:"action"`
	Status string      `json:"status"`
	Error  *Error      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type Error struct {
	Code  string     `json:"code,omitempty"`
	Stack ErrorStack `json:"stack,omitempty"`
}

type ErrorStack []string
