// account_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services

import ( 
	"code.google.com/p/gorest"
	"redrepo-api/models"
    "redrepo-api/dbase"
    "redrepo-api/errors"
	"encoding/json"
	"fmt"
)

type AccountService struct {
	BaseService
    createAccount		gorest.EndPoint `method:"POST"		path:"/signup"   					        postdata:"SignUpParam"`
    retrieveAccount 	gorest.EndPoint	`method:"GET"		path:"/account/{accountId:string}" 	        output:"AccountOutput"`
    retrieveSettings    gorest.EndPoint `method:"GET"       path:"/account/{accountId:string}/settings" output:"AccountSettingOutput"`
    updateAccount 		gorest.EndPoint	`method:"PUT"		path:"/account/" 					        postdata:"AccountSettingParam"`
    deleteAccount		gorest.EndPoint	`method:"DELETE"	path:"/account/{accountId:string}"`	
    verifyAccount       gorest.EndPoint `method:"POST"      path:"/account/verify"                      postdata:"VerificationParam"`
}

func (service AccountService) CreateAccount(param models.SignUpParam) {
    responseString := []byte("")
    responseCode := errors.INTERNAL_SERVER_ERROR
    dbmap := dbase.OpenDatabase()
    var accounts []dbase.Account
    _, err := dbmap.Select(&accounts, "select * from accounts")

    if err != nil {
        response := new(models.GeneralOutput)
        response.Message = "Request completed"
        jsonData, _ := json.Marshal(response)
        responseString = jsonData
        responseCode = 200
    } else {
        response := new(errors.ErrorOutput)
        response.Code = errors.INTERNAL_SERVER_ERROR
        response.Message = "Internal server error"
        jsonData, _ := json.Marshal(response)
        responseString = jsonData
        responseCode = response.Code
    }
   
    dbase.CloseDatabase(dbmap)
    
    service.ResponseBuilder().SetResponseCode(responseCode)
    service.ResponseBuilder().Write([]byte(responseString))
    return
}

func (service AccountService) RetrieveAccount(accountId string) (output models.AccountOutput) {
	// Pseudocode
	// 1. Find the accountId does exist
	// 1.1 If does exist, fire back a response with the info of the account
	// 1.2 Else, fire back a response telling that the account does not exist
	
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

func (service AccountService) RetrieveSettings(accountId string) (output models.AccountSettingOutput) {
    return
}

func (service AccountService) UpdateAccount(param models.AccountSettingParam) {
 
}

func (service AccountService) DeleteAccount(accountId string) {
   
}

func (service AccountService) VerifyAccount(param models.VerificationParam) {
    
}
