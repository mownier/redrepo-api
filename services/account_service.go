// account_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services

import (
    "code.google.com/p/gorest"
    "github.com/dchest/uniuri"
    "redrepo-api/parameter"
    "redrepo-api/response"
    "redrepo-api/dbase"
    "redrepo-api/dbase/tables"
    "redrepo-api/dbase/joins"
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
             var accounts []tables.Account

            _, selectError := dbmap.Select(&accounts, "select id from accounts where username = :username or email = :email limit 1", map[string]interface{}{
                "username": param.Username,
                "email": param.Email,
                }) 

            if selectError == nil {
                if len(accounts) == 0 {
                    account := new(tables.Account)
                    handlers.BindAccountEntryWithSignUpParameter(account, param)

                    inserError := dbmap.Insert(account)
                    if inserError == nil {
                        accountSetting := new(tables.AccountSetting)
                        accountSetting.Username = account.Username
                        inserError = dbmap.Insert(accountSetting)
                        if inserError == nil {
                            verificationCode := new(tables.VerificationCode)
                            verificationCode.Username = account.Username
                            verificationCode.Status = "ACTIVE"
                            verificationCode.Code = uniuri.NewLen(6)
                            inserError = dbmap.Insert(verificationCode)
                            if inserError == nil {
                                accountSettingJoinResult := new(joins.AccountSettingJoinResult)
                                selectError := dbmap.SelectOne(accountSettingJoinResult, "select a.*, s.connected_to_facebook, s.connected_to_twitter, s.verified_account from accounts a, account_settings s WHERE a.username=?", account.Username) 
                                if selectError == nil {
                                    response := new(response.SignUp)
                                    handlers.BindAccountResponseWithResult(response, accountSettingJoinResult)
                                    respData, respCode = response.GetJSONResponseData()
                                    fmt.Printf("success: Responded with json: %s", string(respData))
                                } else {
                                    fmt.Printf("select error: %+v\n", selectError)
                                    respData, respCode = errors.ThrowInternalServerErrorResponse() 
                                }
                            } else {
                                _, deleteError := dbmap.Exec("delete from accounts where username=?", param.Username)
                                if deleteError == nil {
                                    fmt.Printf("insert error: %+v\n", inserError)
                                } else {
                                    fmt.Printf("delete error: %+v\n", deleteError)
                                }
                                respData, respCode = errors.ThrowInternalServerErrorResponse()
                            }
                        } else {
                            _, deleteError := dbmap.Exec("delete from accounts where username=?", param.Username)
                            if deleteError == nil {
                                fmt.Printf("insert error: %+v\n", inserError)
                            } else {
                                fmt.Printf("delete error: %+v\n", deleteError)
                            }
                            respData, respCode = errors.ThrowInternalServerErrorResponse()
                        }
                    } else {
                        fmt.Printf("insert error: %+v\n", inserError)
                        respData, respCode = errors.ThrowInternalServerErrorResponse()
                    }
                } else {
                    fmt.Println("error: Account already exist.")
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
    var respData []byte
    var respCode int

    dbmap, connectionError := dbase.OpenDatabase()
    if (!connectionError) {
        account := new(tables.Account)
        selectError := dbmap.SelectOne(account, "select * from accounts where id=?", accountId) 
        if selectError == nil {
            response := new(response.Account)
            handlers.BindAccountResponse(response, account)
            respData, respCode = response.GetJSONResponseData()
            fmt.Printf("success: Responded with json: %s", string(respData))
        } else {
            fmt.Printf("select error: %+v\n", selectError)
            respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_NOT_FOUND) 
        }
    } else {
        respData, respCode = errors.ThrowInternalServerErrorResponse()
    }
    service.ResponseBuilder().SetResponseCode(respCode)
    service.ResponseBuilder().Write(respData).Overide(true)
    return
}

func (service AccountService) RetrieveSettings(accountId string) (resp response.AccountSetting) {
    var respData []byte
    var respCode int
    dbmap, connectionError := dbase.OpenDatabase()
    if (!connectionError) {
        accountSettingJoinResult := new(joins.AccountSettingJoinResult)
        selectError := dbmap.SelectOne(accountSettingJoinResult, "select a.*, s.connected_to_facebook, s.connected_to_twitter, s.verified_account from accounts a, account_settings s where a.id=?", accountId) 
        if selectError == nil {
            response := new(response.AccountSetting)
            handlers.BindAccountSettingResponse(response, accountSettingJoinResult)
            respData, respCode = response.GetJSONResponseData()
            fmt.Printf("success: Responded with json: %s", string(respData))
        } else {
            fmt.Printf("select error: %+v\n", selectError)
            respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_NOT_FOUND) 
        }
    } else {
        respData, respCode = errors.ThrowInternalServerErrorResponse()
    }
    service.ResponseBuilder().SetResponseCode(respCode)
    service.ResponseBuilder().Write(respData).Overide(true)
    return
}

func (service AccountService) UpdateAccount(param parameter.AccountSetting) {
 
}

func (service AccountService) DeleteAccount(accountId string) {
   
}

func (service AccountService) VerifyAccount(param parameter.Verification) {
    
}
