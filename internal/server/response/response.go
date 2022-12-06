package response

import (
	"encoding/json"
)

func Format(action string, err error, data interface{}) Response {
	var status string
	var errorObject Error

	if err != nil {
		status = StatusFailure
		errorObject = Error{
			Code:  err.Error(),
			Stack: ErrorStack{err.Error()},
		}
	} else {
		status = StatusSuccess
		errorObject = Error{}
	}

	return Response{
		Action: action,
		Status: status,
		Error:  &errorObject,
		Data:   data,
	}
}

func (r Response) Marshal() []byte {
	responseJson, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return responseJson
}
