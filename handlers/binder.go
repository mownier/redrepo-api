package handlers

import (
    "code.google.com/p/go.crypto/bcrypt"
	"redrepo-api/dbase/entries"
	"redrepo-api/parameter"
    "redrepo-api/response"
    "strconv"
	"time"
	)

func BindAccountEntryWithSignUpParameter(account *entries.Account, param parameter.SignUp) {
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

func BindAccountResponseWithAccountEntry(response *response.Account, account entries.Account) {
    response.Id = strconv.Itoa(account.Id)
    response.FirstName = account.FirstName
    response.LastName = account.LastName
    response.Email = account.Email
    response.Username = account.Username
    response.JoinedDate = account.DateJoined
    response.BloodType = account.BloodType
    response.Latitude = account.Latitude
    response.Longitude = account.Longitude
}
