// account_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services

import ( 
	"code.google.com/p/gorest"
)

type AccountService struct {
	BaseService
    createAccount		gorest.EndPoint `method:"POST"		path:"/account/"				postdata:"string"`
    retrieveAccount 	gorest.EndPoint	`method:"GET"		path:"/account/{userId:string}" output:"string"`
    updateAccount 		gorest.EndPoint	`method:"PUT"		path:"/account/" 				postdata:"string"`
    deleteAccount		gorest.EndPoint	`method:"DELETE"	path:"/account/{userId:string}"`
}

func (service AccountService) CreateAccount(account string) {
	response := "{\"first_name\":\"Mounir\",\"last_name\":\"Ybanez\",\"email\":\"rinuom91@gmail.com\",\"username\":\"mownier\",\"latitude\":\"123\",\"longitude\":\"10\"}"
	service.ResponseBuilder().SetResponseCode(201)
	service.ResponseBuilder().Write([]byte(response))
}

func (service AccountService) RetrieveAccount(userId string) string {
	response := "{\"first_name\":\"Mounir\",\"last_name\":\"Ybanez\",\"email\":\"rinuom91@gmail.com\",\"username\":\"mownier\",\"latitude\":\"123\",\"longitude\":\"10\"}"
	service.ResponseBuilder().SetResponseCode(200)
	service.ResponseBuilder().Write([]byte(response))
	return ""
}

func (service AccountService) UpdateAccount(userId string) {
  	response := "{\"first_name\":\"Mounir\",\"last_name\":\"Ybanez\",\"email\":\"rinuom91@gmail.com\",\"username\":\"mownier\",\"latitude\":\"123\",\"longitude\":\"10\"}"
	service.ResponseBuilder().SetResponseCode(200)
	service.ResponseBuilder().Write([]byte(response))
}

func (service AccountService) DeleteAccount(userId string) {
   	response := "{\"first_name\":\"Mounir\",\"last_name\":\"Ybanez\",\"email\":\"rinuom91@gmail.com\",\"username\":\"mownier\",\"latitude\":\"123\",\"longitude\":\"10\"}"
	service.ResponseBuilder().SetResponseCode(200)
	service.ResponseBuilder().Write([]byte(response))
}
