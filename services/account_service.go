// account_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services

import ( 
	"code.google.com/p/gorest"
	"redrepo-api/models"
    "redrepo-api/dbase"
	"encoding/json"
	"fmt"
)

type AccountService struct {
	BaseService
    createAccount		gorest.EndPoint `method:"POST"		path:"/signup"   					        postdata:"SignUpParameters"`
    retrieveAccount 	gorest.EndPoint	`method:"GET"		path:"/account/{accountId:string}" 	        output:"AccountOutput"`
    retrieveSettings    gorest.EndPoint `method:"GET"       path:"/account/{accountId:string}/settings" output:"AccountSettingOutput"`
    updateAccount 		gorest.EndPoint	`method:"PUT"		path:"/account/" 					        postdata:"AccountSettingParameters"`
    deleteAccount		gorest.EndPoint	`method:"DELETE"	path:"/account/{accountId:string}"`	
}

func (service AccountService) CreateAccount(param models.SignUpParameters) {
    dbmap := dbase.OpenDatabase()
    var accounts []dbase.Account
    _, err := dbmap.Select(&accounts, "select * from accounts")
    fmt.Printf("database error: %+v\n", err)
    responseString, err := json.Marshal(param)
    responseCode := 400
    if err == nil { responseCode = 200 }
    if responseCode == 200 {
        fmt.Printf("parameters: %+v\n", param)
        fmt.Printf("accounts: %+v\n", accounts)
        service.ResponseBuilder().SetResponseCode(responseCode)
        service.ResponseBuilder().Write([]byte(responseString))
    } else {
        fmt.Println(err)
    }
    dbase.CloseDatabase(dbmap)
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

func (service AccountService) UpdateAccount(param models.AccountSettingParameters) {
 
}

func (service AccountService) DeleteAccount(accountId string) {
   
}
