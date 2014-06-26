// base_serv.go
// @author Mounir Ybanez
// @date May 7, 2014

package services

import ( 
	"code.google.com/p/gorest"
	"errors"
	"reflect"
)

type BaseService struct {
	gorest.RestService `root:"/api/v1/" consumes:"application/json" produces:"application/json"`
}

func (s BaseService) ValidateAPIKey() error {
	apiKey := s.Context.Request().Header.Get("client_key")
	if !reflect.DeepEqual(apiKey, "redrepoapikey") {
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

func (s BaseService) FireResponse(respData []byte, respCode int) {
	s.ResponseBuilder().SetContentType("application/json")
    s.ResponseBuilder().SetResponseCode(respCode)
    s.ResponseBuilder().Write(respData).Overide(true)
}
