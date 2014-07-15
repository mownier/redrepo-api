// authentication_service.go
// @author Mounir Ybanez
// @date May 6, 2014

package services


import ( 
	"code.google.com/p/gorest"
	"code.google.com/p/go.crypto/bcrypt"
	"redrepo-api/dbase/tables"
	"redrepo-api/response"
	"redrepo-api/parameter"
	"redrepo-api/dbase"
	"redrepo-api/errors"
	"redrepo-api/handlers"
	"github.com/coopernurse/gorp"
)

type AuthenticationService struct {
	BaseService
	signOut				gorest.EndPoint	`method:"DELETE"	path:"/signout"`
	signIn				gorest.EndPoint	`method:"POST"		path:"/signin"			postdata:"SignIn"`
	signInWithFacebook	gorest.EndPoint `method:"POST"		path:"/signin/facebook" postdata:"string"`
	signInWithTwitter	gorest.EndPoint `method:"POST"		path:"/signin/twitter" 	postdata:"string"`
	requestNewPassword	gorest.EndPoint	`method:"POST"		path:"/forgotpassword"	postdata:"string"`
}

func (service AuthenticationService) SignInWithFacebook(params string) {
	var err error
	var dbmap *gorp.DbMap

	defer errors.CatchPanic(&err, "SignInWithFacebook")
	
	dbmap, err = dbase.OpenDatabase()
    
    dbase.CloseDatabase(dbmap)
}

func (service AuthenticationService) SignInWithTwitter(params string) {

}

func (service AuthenticationService) SignIn(param parameter.SignIn) {
	var respData []byte
    var respCode int

    clientKey, clientError := service.GetClientKey()
    if clientError == nil {
    	dbmap, connectionError := dbase.OpenDatabase()
        if connectionError == nil {
        	account := new(tables.Account)
    		selectError := dbmap.SelectOne(account, "select a.* from accounts a, account_settings s where a.username=s.username and (a.username=? or a.email=?) and s.verified=?", param.Username, param.Username,1)
    		if selectError == nil {
    			mismatchPasswordError := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(param.Password))
    			if mismatchPasswordError == nil {
    				authenticatedAccount := new(response.AuthenticatedAccount)
	    			handlers.BindAuthenticatedAccountResponse(authenticatedAccount, account)

    				authSession := new(tables.AuthSession)
    				authSession.AccountId = account.Id
    				authSession.AccessToken = authenticatedAccount.AccessToken
    				authSession.ClientKey = clientKey
    				authSession.AuthProvider = "NONE"

    				insertError := dbmap.Insert(authSession)
    				if insertError == nil {
						respData, respCode = authenticatedAccount.GetJSONResponseData()
					} else {
						errors.Log(insertError)
			            respData, respCode = errors.ThrowInternalServerErrorResponse()
					}
				} else {
					errors.LogErrorMessage(errors.ErrorMessage(errors.MISMATCH_USERNAME_PASSWORD))
					respData, respCode = errors.ErrorResponseData(errors.MISMATCH_USERNAME_PASSWORD)
				}  			
			} else {
				errors.Log(selectError)
                respData, respCode = errors.ErrorResponseData(errors.ACCOUNT_NOT_VERIFIED)
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

func (service AuthenticationService) SignOut() {
	var respData []byte
    var respCode int

    accessToken, authError := service.GetAccessToken()
    if authError == nil {
    	dbmap, connectionError := dbase.OpenDatabase()
    	if connectionError == nil {
    		_, deleteError := dbmap.Exec("delete from auth_sessions where access_token=?", accessToken)
    		if deleteError == nil {
    			respCode = 200
    			respData = []byte(`{ "message" : "Successfully signed out." }`)
			} else {
				errors.Log(deleteError)
	            respData, respCode = errors.ThrowInternalServerErrorResponse()
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

func (service AuthenticationService) RequestNewPassword(params string) {

}




