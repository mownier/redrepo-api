// base_serv.go
// @author Mounir Ybanez
// @date May 7, 2014

package services

import ( 
	"code.google.com/p/gorest"
	"errors"
	"redrepo-api/dbase"
	"redrepo-api/dbase/tables"
	RR_errors "redrepo-api/errors"
)

type BaseService struct {
	gorest.RestService `root:"/api/v1/" consumes:"application/json" produces:"application/json"`
}

// TODO: Refactor to 'CheckClientKey'
func (s BaseService) GetClientKey() (string, error) {
	var returnError error
	var returnString string

	clientKey := s.Context.Request().Header.Get("Client-Key")
	dbmap, connectionError := dbase.OpenDatabase()
	if connectionError == nil {
		authClient := new(tables.AuthClient)
		selectError := dbmap.SelectOne(authClient, "select a.name from auth_clients a where a.key=?", clientKey)
		if selectError == nil {
			returnString = clientKey
			returnError = nil
		} else {
			returnString = ""
			returnError = errors.New(RR_errors.ErrorMessage(RR_errors.CLIENT_NOT_ALLOWED))
		}
	} else {
		returnString = ""
		returnError = connectionError
	}
	
	dbase.CloseDatabase(dbmap)
	return returnString, returnError
}

// TODO: Refactor to 'CheckAccessToken'
func (s BaseService) GetAccessToken() (string, error) {
	var returnError error
	var returnString string

	accessToken := s.Context.Request().Header.Get("Authorization")
	dbmap, connectionError := dbase.OpenDatabase()
	if connectionError == nil {
		authSession := new(tables.AuthSession)
		selectError := dbmap.SelectOne(authSession, "select a.account_id from auth_sessions a where a.access_token=?", accessToken)
		if selectError == nil {
			returnString = accessToken
			returnError = nil
		} else {
			returnString = ""
			returnError = errors.New(RR_errors.ErrorMessage(RR_errors.NOT_AUTHORIZED))
		}
	} else {
		returnString = ""
		returnError = connectionError
	}

	dbase.CloseDatabase(dbmap)
	return returnString, returnError
}

func (s BaseService) FireResponse(respData []byte, respCode int) {
	s.ResponseBuilder().SetContentType("application/json")
    s.ResponseBuilder().SetResponseCode(respCode)
    s.ResponseBuilder().Write(respData).Overide(true)
}
