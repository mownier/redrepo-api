package response

type AuthenticatedAccountOutput struct {
	BaseResponse
	Account 	*Account `json:"account"`
	AccessToken	string	 `json:"access_token"`		
}