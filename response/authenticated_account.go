package response

import ( 
	"net/http"
	"encoding/json"
	)

type AuthenticatedAccount struct {
	Account
	AccessToken	string	 `json:"access_token"`		
}

func (resp AuthenticatedAccount) GetJSONResponseData() ([]byte, int) {
	jsonData, _ := json.Marshal(resp)
	return jsonData, http.StatusOK
}