package wdbErrors

// getters for errors
func (e *WdbError) Error() string {
	return e.ErrMessage
}

func (e *WdbError) StatusCode() int {
	return e.HttpStatusCode
}

func (e *WdbError) Code() string {

	return e.ErrCode
}

// setters for errors
func (e *WdbError) SetCode(code string) *WdbError {
	e.ErrCode = code
	return e
}

func (e *WdbError) SetMessage(message string) *WdbError {
	e.ErrMessage = message
	return e
}

func (e *WdbError) SetStatusCode(code int) *WdbError {
	e.HttpStatusCode = code
	return e
}
