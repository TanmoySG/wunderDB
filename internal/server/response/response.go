package response

import (
	"encoding/json"

	er "github.com/TanmoySG/wunderDB/internal/errors"
)

func Format(action string, err *er.WdbError, data interface{}) ApiResponse {
	var status string
	var errorObject Error
	var httpStatusCode int

	if err != nil {
		status = StatusFailure
		errorObject = Error{
			Code:  err.ErrCode,
			Stack: ErrorStack{err.ErrMessage},
		}
		httpStatusCode = err.HttpStatusCode
	} else {
		status = StatusSuccess
		errorObject = Error{}
		httpStatusCode = 200

	}

	return ApiResponse{
		HttpStatusCode: httpStatusCode,
		Response: Response{
			Action: action,
			Status: status,
			Error:  &errorObject,
			Data:   data,
		},
	}
}

func (r ApiResponse) Marshal() []byte {
	responseJson, err := json.Marshal(r.Response)
	if err != nil {
		return nil
	}
	return responseJson
}
