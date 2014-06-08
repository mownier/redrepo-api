// error_handler.go
// @author Mounir Ybanez
// @date May 24, 2014

package handlers

const(
	SUCCESS = 200
	AUTHENTICATION_ERROR = 401
	BAD_REQUEST = 400
	INTERNAL_SERVER_ERROR = 500
	DATABASE_SERVER_ERROR = 501
)

type ErrorHandler struct {

}
