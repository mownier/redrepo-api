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
    "redrepo-api/mail"
    "fmt"
    )

type AccountService struct {
    BaseService
    createAccount       gorest.EndPoint `method:"POST"      path:"/signup"                              postdata:"SignUp"`
    retrieveAccount     gorest.EndPoint `method:"GET"       path:"/account/{accountId:string}"          output:"Account"`
    retrieveSettings    gorest.EndPoint `method:"GET"       path:"/account/{accountId:string}/settings" output:"AccountSetting"`
    updateAccount       gorest.EndPoint `method:"PUT"       path:"/account/"                            postdata:"AccountSetting"`
    deleteAccount       gorest.EndPoint `method:"DELETE"    path:"/account/{accountId:string}"`  
    verifyAccount       gorest.EndPoint `method:"POST"      path:"/account/verification"                postdata:"Verification"`
    requestNewCode      gorest.EndPoint `method:"POST"      path:"/account/verification/newcode"        postdata:"NewCode"`
}

func (service AccountService) CreateAccount(param parameter.SignUp) { 
    var respData []byte
    var respCode int
    
    clientError := service.BaseService.ValidateAPIKey()
    if clientError == nil {
        paramError := param.ValidateValues()
        if paramError == nil {
            dbmap, connectionError := dbase.OpenDatabase()
            if connectionError == nil {
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
                                    // Sends an email verification asynchronously 
                                    go func() {
                                        accountVerificationJoinResult := new(joins.AccountVerificationJoinResult)
                                        dbmap.SelectOne(accountVerificationJoinResult, "select v.code, v.status, a.email from accounts a, verification_codes v where a.username=v.username and v.code=?", verificationCode.Code)
                                        mailService := mail.DefaultMailService()
                                        sendMailError := mailService.Send(mail.GetEmailVerificationTemplate(accountVerificationJoinResult.Code, accountVerificationJoinResult.Email), 
                                                            []string{accountVerificationJoinResult.Email})
                                        if sendMailError == nil {
                                           fmt.Println("Successfully sent verification code to email: " + account.Email)
                                        } else {
                                            errors.LogErrorMessage("Unable to send verification code to email: " + account.Email)
                                        }
                                    }()
                                    accountSettingJoinResult := new(joins.AccountSettingJoinResult)
                                    selectError := dbmap.SelectOne(accountSettingJoinResult, "select a.*, s.connected_to_facebook, s.connected_to_twitter, s.verified_account from accounts a, account_settings s WHERE a.username=?", account.Username) 
                                    if selectError == nil {
                                        response := new(response.SignUp)
                                        handlers.BindAccountResponseWithResult(response, accountSettingJoinResult)
                                        respData, respCode = response.GetJSONResponseData()
                                        fmt.Printf("success: Responded with json: %s\n", string(respData))
                                    } else {
                                        errors.Log(selectError)
                                        respData, respCode = errors.ThrowInternalServerErrorResponse() 
                                    }
                                } else {
                                    _, deleteError := dbmap.Exec("delete from accounts where username=?", param.Username)
                                    if deleteError == nil {
                                        errors.Log(inserError)
                                    } else {
                                        errors.Log(deleteError)
                                    }
                                    respData, respCode = errors.ThrowInternalServerErrorResponse()
                                }
                            } else {
                                _, deleteError := dbmap.Exec("delete from accounts where username=?", param.Username)
                                if deleteError == nil {
                                    errors.Log(inserError)
                                } else {
                                    errors.Log(deleteError)
                                }
                                respData, respCode = errors.ThrowInternalServerErrorResponse()
                            }
                        } else {
                            errors.Log(inserError)
                            respData, respCode = errors.ThrowInternalServerErrorResponse()
                        }
                    } else {
                        errors.LogErrorMessage(errors.ErrorMessage(errors.ACCOUNT_ALREADY_EXIST))
                        respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_ALREADY_EXIST)
                    }
                } else {
                    errors.Log(selectError)
                    respData, respCode = errors.ThrowInternalServerErrorResponse()
                }
            } else {
                errors.Log(connectionError)
                respData, respCode = errors.ThrowInternalServerErrorResponse()
            }
            dbase.CloseDatabase(dbmap)
        } else {
            errors.Log(paramError)
            respData, respCode = errors.ErrorResponseData(errors.INVALID_PARAMETER_VALUE)
        }
    } else {
        errors.Log(clientError)
        respData, respCode = errors.ErrorResponseData(errors.CLIENT_NOT_ALLOWED)
    }
    
    service.BaseService.FireResponse(respData, respCode)
    return
}

func (service AccountService) RetrieveAccount(accountId string) (resp response.Account) {
    var respData []byte
    var respCode int

    _, authError := service.BaseService.GetAccessToken()
    if authError == nil {
        dbmap, connectionError := dbase.OpenDatabase()
        if connectionError == nil {
            account := new(tables.Account)
            selectError := dbmap.SelectOne(account, "select * from accounts where id=? or username=?", accountId, accountId) 
            if selectError == nil {
                response := new(response.Account)
                handlers.BindAccountResponse(response, account)
                respData, respCode = response.GetJSONResponseData()
                fmt.Printf("success: Responded with json: %s", string(respData))
            } else {
                errors.Log(selectError)
                respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_NOT_FOUND) 
            }
        } else {
            errors.Log(connectionError)
            respData, respCode = errors.ThrowInternalServerErrorResponse()
        }
        dbase.CloseDatabase(dbmap)
    } else {
        errors.Log(authError)
        respData, respCode = errors.ErrorResponseData(errors.NOT_AUTHORIZED) 
    }
    service.BaseService.FireResponse(respData, respCode)
    return
}

func (service AccountService) RetrieveSettings(accountId string) (resp response.AccountSetting) {
    var respData []byte
    var respCode int

    _, authError := service.BaseService.GetAccessToken()
    if authError == nil {
        dbmap, connectionError := dbase.OpenDatabase()
        if connectionError == nil {
            accountSettingJoinResult := new(joins.AccountSettingJoinResult)
            selectError := dbmap.SelectOne(accountSettingJoinResult, "select a.*, s.connected_to_facebook, s.connected_to_twitter, s.verified_account from accounts a, account_settings s where a.id=? or a.username=?", accountId, accountId) 
            if selectError == nil {
                response := new(response.AccountSetting)
                handlers.BindAccountSettingResponse(response, accountSettingJoinResult)
                respData, respCode = response.GetJSONResponseData()
                fmt.Printf("success: Responded with json: %s", string(respData))
            } else {
                errors.Log(selectError)
                respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_NOT_FOUND) 
            }
        } else {
            errors.Log(connectionError)
            respData, respCode = errors.ThrowInternalServerErrorResponse()
        }
        dbase.CloseDatabase(dbmap)
    } else {
        errors.Log(authError)
        respData, respCode = errors.ErrorResponseData(errors.NOT_AUTHORIZED) 
    }
    service.BaseService.FireResponse(respData, respCode)
    return
}

func (service AccountService) UpdateAccount(param parameter.AccountSetting) {
 
}

func (service AccountService) DeleteAccount(accountId string) {
   
}

func (service AccountService) VerifyAccount(param parameter.Verification) {
    var respData []byte
    var respCode int
    clientError := service.BaseService.ValidateAPIKey()
    if clientError == nil {
        dbmap, connectionError := dbase.OpenDatabase()
        if connectionError == nil {
            accountVerificationJoinResult := new(joins.AccountVerificationJoinResult)
            selectError := dbmap.SelectOne(accountVerificationJoinResult, "select v.id, v.code, v.status, a.email, a.username from accounts a, verification_codes v where a.username=v.username and v.code=? and v.status=?", param.Code, "ACTIVE")
            if selectError == nil {
                _, updateError := dbmap.Exec("update account_settings a set a.verified_account=? where a.username=?", 1, accountVerificationJoinResult.Username)
                if updateError == nil {
                    _, updateError = dbmap.Exec("update verification_codes v SET v.status=? where v.code=?", "EXPIRED", accountVerificationJoinResult.Code)
                    if updateError == nil {
                        respCode = 200
                        respData = []byte(`{ "message" : "Verification successful"}`)
                    } else {
                        errors.Log(updateError)
                        respData, respCode = errors.ThrowInternalServerErrorResponse()
                    }
                } else {
                    errors.Log(updateError)
                    respData, respCode = errors.ThrowInternalServerErrorResponse()
                }
            } else {
                errors.Log(selectError)
                respData, respCode = errors.ErrorResponseData(errors.VERIFICATION_CODE_EXPIRED)
            }
        } else {
            errors.Log(connectionError)
            respData, respCode = errors.ThrowInternalServerErrorResponse()
        }
        dbase.CloseDatabase(dbmap)
    } else {
        errors.Log(clientError)
        respData, respCode = errors.ErrorResponseData(errors.CLIENT_NOT_ALLOWED)
    }
    service.BaseService.FireResponse(respData, respCode)
}

func (service AccountService) RequestNewCode(param parameter.NewCode) {
    var respData []byte
    var respCode int

    clientError := service.BaseService.ValidateAPIKey()
    if clientError == nil {
        dbmap, connectionError := dbase.OpenDatabase()
        if connectionError == nil {
            accountSettingJoinResult := new(joins.AccountSettingJoinResult)
            selectError := dbmap.SelectOne(accountSettingJoinResult, "select s.verified_account, a.username from accounts a, account_settings s WHERE a.email=?", param.Email)
            if selectError == nil {
                if accountSettingJoinResult.VerifiedAccount == 0 {
                    // TODO: Check if naay active the verification code para sa user
                    // If naay active, update table: 'verification_codes'
                    // else insert new verification_code
                    success := false
                    accountVerificationJoinResult := new(joins.AccountVerificationJoinResult)
                    selectError = dbmap.SelectOne(accountVerificationJoinResult, "select v.code from accounts a, verification_codes v where v.username=? and v.status=?", accountSettingJoinResult.Username, "ACTIVE")
                    var newCode string
                    if selectError == nil {
                        // Update table: 'verification_codes'
                        newCode = uniuri.NewLen(6)
                        _, updateError := dbmap.Exec("update verification_codes v set v.code=? where v.username=? and v.code=?", newCode, accountSettingJoinResult.Username, accountVerificationJoinResult.Code)
                        if updateError == nil {
                            success = true
                        } else {
                            errors.Log(updateError)
                            respData, respCode = errors.ThrowInternalServerErrorResponse()
                        }
                    } else {
                        // Insert new verification_code
                        verificationCode := new(tables.VerificationCode)
                        verificationCode.Username = accountSettingJoinResult.Username
                        verificationCode.Status = "ACTIVE"
                        verificationCode.Code = uniuri.NewLen(6)
                        newCode = verificationCode.Code
                        inserError := dbmap.Insert(verificationCode)
                        if inserError == nil {
                            success = true
                        } else {
                            errors.Log(inserError)
                            respData, respCode = errors.ThrowInternalServerErrorResponse()
                        }
                    }
                    if success == true {
                        // Send Email
                        mailService := mail.DefaultMailService()
                        sendMailError := mailService.Send(mail.GetEmailVerificationTemplate(newCode, param.Email), []string{param.Email})
                        if sendMailError == nil {
                            respCode = 200
                            respData = []byte(`{ "message" : "Successfully sent new verification code."}`)
                        } else {
                            errors.Log(sendMailError)
                            respData, respCode = errors.ThrowInternalServerErrorResponse()
                        }
                    }     
                } else {
                    errors.LogErrorMessage(errors.ErrorMessage(errors.ACCOUNT_ALREADY_VERIFIED))
                    respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_ALREADY_VERIFIED)
                }
            } else {
                errors.Log(selectError)
                respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_NOT_FOUND)
            }
        } else {
            errors.Log(connectionError)
            respData, respCode = errors.ThrowInternalServerErrorResponse()
        }
        dbase.CloseDatabase(dbmap)
    } else {
        errors.Log(clientError)
        respData, respCode = errors.ErrorResponseData(errors.CLIENT_NOT_ALLOWED)
    }
    service.BaseService.FireResponse(respData, respCode)
    return
}

