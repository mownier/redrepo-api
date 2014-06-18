package handlers

import (
    "code.google.com/p/go.crypto/bcrypt"
	"redrepo-api/dbase/entries"
	"redrepo-api/parameter"
	"time"
	)

func Bind(entry *entries.Account, param parameter.SignUp) {
	entry.FirstName = param.FirstName
    entry.LastName = param.LastName
    entry.Email = param.Email

    pass, _ := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
    entry.Password = string(pass)
    
    entry.Username = param.Username
    entry.BloodType = param.BloodType
    entry.Latitude = param.Latitude
    entry.Longitude = param.Longitude
    
    now := time.Now()
    entry.DateJoined = now.Format(time.RFC3339)
}
