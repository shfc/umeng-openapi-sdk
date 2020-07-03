package errors

import (
	"encoding/json"
)

type Err struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
	RequestId    string `json:"request_id"`
}

func (e *Err) Error() string {
	r, _ := json.Marshal(e)
	return string(r)
}