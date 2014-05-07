// init.go
// @author Mounir Ybanez
// @date May 6, 2014

package services

import (
	"code.google.com/p/gorest"
	"net/http"
)

func Start() {
	// Register an instance of AccountService
	gorest.RegisterService(new(AccountService))
	gorest.RegisterService(new(AuthenticationService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":2121", nil)
}
