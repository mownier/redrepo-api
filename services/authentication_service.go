// authentication_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services


import ( 
	"code.google.com/p/gorest"
)

type AuthenticationService struct {
	BaseService
	signOut				gorest.EndPoint	`method:"POST"	path:"/signout" 		postdata:"string"`
	signIn				gorest.EndPoint	`method:"POST"	path:"/signin"			postdata:"string"`
	signInWithFacebook	gorest.EndPoint `method:"POST"	path:"/signin/facebook" postdata:"string"`
	signInWithTwitter	gorest.EndPoint `method:"POST"	path:"/signin/twitter" 	postdata:"string"`
	requestNewPassword	gorest.EndPoint	`method:"POST"	path:"/forgotpassword"	postdata:"string"`
}

func (service AuthenticationService) SignInWithFacebook(params string) {
	responseString := "Sign in with Facebook"
	service.ResponseBuilder().Write([]byte(responseString))
}

func (service AuthenticationService) SignInWithTwitter(params string) {

}

func (service AuthenticationService) SignIn(params string) {

}

func (service AuthenticationService) SignOut(params string) {
	
}

func (service AuthenticationService) RequestNewPassword(params string) {

}
