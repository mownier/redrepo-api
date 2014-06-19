package handlers

import (
    "code.google.com/p/go.crypto/bcrypt"
	"redrepo-api/dbase/tables"
    "redrepo-api/dbase/joins"
	"redrepo-api/parameter"
    "redrepo-api/response"
    "strconv"
	"time"
	)

func BindAccountEntryWithSignUpParameter(account *tables.Account, param parameter.SignUp) {
	account.FirstName = param.FirstName
    account.LastName = param.LastName
    account.Email = param.Email

    pass, _ := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
    account.Password = string(pass)
    
    account.Username = param.Username
    account.BloodType = param.BloodType
    account.Latitude = param.Latitude
    account.Longitude = param.Longitude
    
    now := time.Now()
    account.DateJoined = now.Format(time.RFC3339)
}

func BindAccountResponseWithResult(response *response.SignUp, result *joins.AccountSettingJoinResult) {
    response.Id = strconv.Itoa(result.Id)
    response.FirstName = result.FirstName
    response.LastName = result.LastName
    response.Email = result.Email
    response.Username = result.Username
    response.JoinedDate = result.DateJoined
    response.BloodType = result.BloodType
    response.Latitude = result.Latitude
    response.Longitude = result.Longitude
    response.VerifiedAccount = result.VerifiedAccount
}

func BindAccountSettingResponseWithResult(response *response.AccountSetting, result *joins.AccountSettingJoinResult) {
    response.Id = strconv.Itoa(result.Id)
    response.FirstName = result.FirstName
    response.LastName = result.LastName
    response.Email = result.Email
    response.Username = result.Username
    response.JoinedDate = result.DateJoined
    response.BloodType = result.BloodType
    response.Latitude = result.Latitude
    response.Longitude = result.Longitude
    response.ConnectedToTwitter = result.ConnectedToTwitter
    response.ConnectedToFacebook = result.ConnectedToFacebook
}

func BindAccountResponse(response *response.Account, result *tables.Account) {
    response.Id = strconv.Itoa(result.Id)
    response.FirstName = result.FirstName
    response.LastName = result.LastName
    response.Email = result.Email
    response.Username = result.Username
    response.JoinedDate = result.DateJoined
    response.BloodType = result.BloodType
    response.Latitude = result.Latitude
    response.Longitude = result.Longitude
}

func BindAccountSettingResponse(response *response.AccountSetting, result *joins.AccountSettingJoinResult) {
    response.Id = strconv.Itoa(result.Id)
    response.FirstName = result.FirstName
    response.LastName = result.LastName
    response.Email = result.Email
    response.Username = result.Username
    response.JoinedDate = result.DateJoined
    response.BloodType = result.BloodType
    response.Latitude = result.Latitude
    response.Longitude = result.Longitude
    response.ConnectedToFacebook = result.ConnectedToFacebook
    response.ConnectedToTwitter = result.ConnectedToTwitter
}
