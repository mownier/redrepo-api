// authentication_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services


import ( 
	"fmt"
	"code.google.com/p/gorest"
	"encoding/json"
)

type AuthenticationService struct {
	BaseService
	signOut				gorest.EndPoint	`method:"POST"	path:"/signout" 		postdata:"string"`
	signIn				gorest.EndPoint	`method:"POST"	path:"/signin"			postdata:"string"`
	requestNewPassword	gorest.EndPoint	`method:"POST"	path:"/forgotpassword"	postdata:"string"`
}

func (service AuthenticationService) SignIn(params string) {
	fmt.Println("Signing in...")
}

func (service AuthenticationService) SignOut(params string) {
	
}

func (service AuthenticationService) RequestNewPassword(params string) {

}
