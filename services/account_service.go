// account_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services

import (
    "code.google.com/p/gorest"
    "redrepo-api/parameter"
    "redrepo-api/response"
    "redrepo-api/dbase"
    "redrepo-api/dbase/entries"
    "redrepo-api/handlers"
    "redrepo-api/errors"
    "fmt"
    )

type AccountService struct {
    BaseService
    createAccount       gorest.EndPoint `method:"POST"      path:"/signup"                              postdata:"SignUp"`
    retrieveAccount     gorest.EndPoint `method:"GET"       path:"/account/{accountId:string}"          output:"Account"`
    retrieveSettings    gorest.EndPoint `method:"GET"       path:"/account/{accountId:string}/settings" output:"AccountSetting"`
    updateAccount       gorest.EndPoint `method:"PUT"       path:"/account/"                            postdata:"AccountSetting"`
    deleteAccount       gorest.EndPoint `method:"DELETE"    path:"/account/{accountId:string}"`  
    verifyAccount       gorest.EndPoint `method:"POST"      path:"/account/verify"                      postdata:"Verification"`
}

func (service AccountService) CreateAccount(param parameter.SignUp) {    
    var respData []byte
    var respCode int
    
    // Check if the given parameters have errors
    if param.HasErrors() == true {
        respData, respCode = errors.ErrorResponseData(errors.INVALID_PARAMETER_VALUE)
    } else {
        dbmap, connectionError := dbase.OpenDatabase()
        if (!connectionError) {
             var accounts []entries.Account

            _, selectError := dbmap.Select(&accounts, "select id from accounts where username = :username or email = :email limit 1", map[string]interface{}{
                "username": param.Username,
                "email": param.Email,
                }) 

            if selectError == nil {
                if len(accounts) == 0 {
                    account := new(entries.Account)
                    handlers.Bind(account, param)

                    inserError := dbmap.Insert(account)
                    if (inserError == nil) {
                        accounts = nil
                        _, selectError := dbmap.Select(&accounts, "select * from accounts where username = :username or email = :email limit 1", map[string]interface{}{
                            "username": param.Username,
                            "email": param.Email,
                            }) 
                        if selectError == nil {
                            fmt.Printf("accounts: %+v", accounts)
                            resp := new(response.Account)
                            respData, respCode = resp.GetJSONResponseData()
                        } else {
                            fmt.Printf("select error: %+v\n", selectError)
                            respData, respCode = errors.ThrowInternalServerErrorResponse() 
                        }
                    } else {
                        fmt.Printf("insert error: %+v\n", inserError)
                        respData, respCode = errors.ThrowInternalServerErrorResponse()
                    }
                } else {
                    fmt.Printf("error: Account already exist.\n")
                    respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_ALREADY_EXIST)
                }
            } else {
                fmt.Printf("select error: %+v\n", selectError)
                respData, respCode = errors.ThrowInternalServerErrorResponse()
            }

            dbase.CloseDatabase(dbmap)

        } else {
            respData, respCode = errors.ThrowInternalServerErrorResponse()
        }
       
    }

    service.ResponseBuilder().SetResponseCode(respCode)
    service.ResponseBuilder().Write(respData)

    return
}

func (service AccountService) RetrieveAccount(accountId string) (resp response.Account) {
    service.ResponseBuilder().SetResponseCode(200)
    return
}

func (service AccountService) RetrieveSettings(accountId string) (resp response.AccountSetting) {
    return
}

func (service AccountService) UpdateAccount(param parameter.AccountSetting) {
 
}

func (service AccountService) DeleteAccount(accountId string) {
   
}

func (service AccountService) VerifyAccount(param parameter.Verification) {
    
}
