package response

import (
	"encoding/json"

	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

func Format(action string, err *er.WdbError, data interface{}, notices ...string) ApiResponse {
	// default values
	var status string = StatusSuccess
	var errorObject Error = Error{}
	var httpStatusCode int = 200

	if err != nil {
		status = StatusFailure
		errorObject = Error{
			Code:  err.ErrCode,
			Stack: ErrorStack{err.ErrMessage},
		}
		httpStatusCode = err.HttpStatusCode
	}

	return ApiResponse{
		HttpStatusCode: httpStatusCode,
		Response: Response{
			Action: action,
			Status: status,
			Error:  &errorObject,
			// Data:     &data, // would be removed in future in favour of Response. See : https://github.com/TanmoySG/wunderDB/issues/121
			Response: &data,
			Notices:  notices,
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
