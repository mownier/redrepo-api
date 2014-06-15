// error_handler.go
// @author Mounir Ybanez
// @date May 24, 2014

package errors

const(
	AUTHENTICATION_ERROR = 401
	BAD_REQUEST = 400
	INTERNAL_SERVER_ERROR = 500
)

type ErrorOutput struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}
