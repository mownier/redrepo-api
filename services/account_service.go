// account_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services

import ( 
	"code.google.com/p/gorest"
	"redrepo-api/models"
	"encoding/json"
	"fmt"
)

type AccountService struct {
	BaseService
    createAccount		gorest.EndPoint `method:"POST"		path:"/account/"					postdata:"UserParamModel"`
    retrieveAccount 	gorest.EndPoint	`method:"GET"		path:"/account/{userId:string}" 	output:"UserOutputModel"`
    updateAccount 		gorest.EndPoint	`method:"PUT"		path:"/account/" 					postdata:"UserParamModel"`
    deleteAccount		gorest.EndPoint	`method:"DELETE"	path:"/account/{userId:string}"`	
}

func (service AccountService) CreateAccount(param models.UserParamModel) {
	// Pseudocode
	// 1. Insert new user
	// 2. After successfully insert, fire back the response

	authenticatedUser 					:= new(models.AuthenticatedUserOutputModel)
	authenticatedUser.User 				= new(models.UserOutputModel)
	authenticatedUser.User.FirstName 	= param.FirstName
	authenticatedUser.User.LastName 	= param.LastName
	authenticatedUser.User.Email 		= param.Email
	authenticatedUser.User.Username 	= param.Username
	authenticatedUser.AccessToken 		= "12321321hkjhj213"

    responseString, err := json.Marshal(authenticatedUser)
    responseCode := 400
    if err == nil {
    	fmt.Printf("%+v\n", param)
    	service.ResponseBuilder().Write([]byte(responseString))
    	responseCode = 200 
    } else {
    	fmt.Println(err)
    }

    service.ResponseBuilder().SetResponseCode(responseCode)
}

func (service AccountService) RetrieveAccount(userId string) (output models.UserOutputModel) {
	// Pseudocode
	// 1. Find the userId does exist
	// 1.1 If does exist, fire back a response with the info of the user
	// 1.2 Else, fire back a response telling that the user does not exist
	
	jsonString := `{"id":"1","first_name":"Mounir","last_name":"Ybanez","email":"rinuom91@gmail.com","username":"mownier","date_joined":"May 24, 2014","latitude":123,"longitude":10,"connected_to_facebook":0,"connected_to_twitter":0}`
    err := json.Unmarshal([]byte(jsonString), &output)
    responseCode := 400
    if err == nil {
        fmt.Printf("%+v\n", output)
        responseCode = 200
    } else {
        fmt.Println(err)
    }

    service.ResponseBuilder().SetResponseCode(responseCode)
	return
}

func (service AccountService) UpdateAccount(param models.UserParamModel) {
 
}

func (service AccountService) DeleteAccount(userId string) {
   
}
