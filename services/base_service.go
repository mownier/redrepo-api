// base_serv.go
// @author Mounir Ybanez
// @date May 7, 2014

package services

import ( 
	"code.google.com/p/gorest"
	"errors"
	"reflect"
)

const (
	EMAIL_SENDER_USERNAME = "redrepo.mail@gmail.com"
	EMAIL_SENDER_PASS = "iamredrepo"
	EMAIL_SENDER_SEREVER = "smtp.gmai.com"
	EMAIL_SENDER_PORT = 3232
)

type BaseService struct {
	gorest.RestService `root:"/api/v1/" consumes:"application/json" produces:"application/json"`
}

func (s BaseService) ValidateAPIKey() error {
	apiKey := s.Context.Request().Header.Get("client_key")
	if !reflect.DeepEqual(apiKey, "12345") {
		return errors.New("Invalid api key.")
	}
	return nil
}

func (s BaseService) GetAccessToken() (string, error) {
	accessToken := s.Context.Request().Header.Get("Authorization")
	if len(accessToken) == 0 {
		return accessToken, errors.New("Not authorized.")
	}
	return accessToken, nil
}
