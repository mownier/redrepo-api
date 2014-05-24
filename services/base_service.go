// base_serv.go
// @author Mounir Ybanez
// @date May 7, 2014

package services

import ( 
	"code.google.com/p/gorest"
)

type BaseService struct {
	gorest.RestService `root:"/api/v1/" consumes:"application/json" produces:"application/json"`
}
