// init.go
// @author Mounir Ybanez
// @date May 6, 2014

package services

import (
	"strconv"
	"code.google.com/p/gorest"
	"net/http"
)

func Start(port int) {
	gorest.RegisterService(new(AccountService))
	gorest.RegisterService(new(AuthenticationService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}
