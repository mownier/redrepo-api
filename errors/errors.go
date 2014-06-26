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
	ACCOUNT_NOT_FOUND = 702
	CLIENT_NOT_ALLOWED = 703
	NOT_AUTHORIZED = 704
	VERIFICATION_CODE_EXPIRED = 705
	ACCOUNT_ALREADY_VERIFIED = 706
)

type ErrorResponse struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

var errorMessage = map[int]string {
	ACCOUNT_ALREADY_EXIST: "Account already exist.",
	INVALID_PARAMETER_VALUE: "Missing parameter or invalid value.",
	ACCOUNT_NOT_FOUND: "Account not found.",
	CLIENT_NOT_ALLOWED: "Client not allowed.",
	NOT_AUTHORIZED: "Not authorized.",
	VERIFICATION_CODE_EXPIRED: "Verification code expired.",
	ACCOUNT_ALREADY_VERIFIED: "Account already verified.",
}

func ErrorMessage(errorCode int) string {
	return errorMessage[errorCode]
}

func ErrorResponseData(errorCode int) ([]byte, int) {
	errorResponse := new(ErrorResponse)
    errorResponse.Code = errorCode
    errorResponse.Message = ErrorMessage(errorCode)
    jsonData, _ := json.Marshal(errorResponse)
	return jsonData, http.StatusBadRequest
}

func ThrowInternalServerErrorResponse() ([]byte, int) {
	resp := new(ErrorResponse)
	resp.Code = http.StatusInternalServerError
	resp.Message = http.StatusText(http.StatusInternalServerError)
	jsonData, _ := json.Marshal(resp)
	return jsonData, http.StatusInternalServerError
}
