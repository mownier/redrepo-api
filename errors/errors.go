// error_handler.go
// @author Mounir Ybanez
// @date May 24, 2014

package errors

import (
	"encoding/json" 
	"net/http"
)

const (
	ACCOUNT_ALREADY_EXIST = 700
	INVALID_PARAMETER_VALUE = 701
)

type ErrorResponse struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

var errorMessage = map[int]string {
	700: "Account already exist.",
	701: "Missing parameter or invalid value.",
}

func ErrorMessage(errorCode int) string {
	return errorMessage[errorCode]
}

func ErrorResponseData(errorCode int) ([]byte, int) {
	errorResponse := new(ErrorResponse)
    errorResponse.Code = errorCode
    errorResponse.Message = ErrorMessage(errorCode)
    jsonData, _ := json.Marshal(errorResponse)
	return jsonData, errorCode
}

func ThrowInternalServerErrorResponse() ([]byte, int) {
	resp := new(ErrorResponse)
	resp.Code = http.StatusInternalServerError
	resp.Message = http.StatusText(http.StatusInternalServerError)
	jsonData, _ := json.Marshal(resp)
	return jsonData, http.StatusInternalServerError
}
