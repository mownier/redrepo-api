package response 

import (
	"encoding/json"
	"net/http"
	)

type SignUp struct {
	BaseResponse
	Id 					string	`json:"id"`
	FirstName 			string	`json:"first_name"`
	LastName 			string	`json:"last_name"`
	Email				string	`json:"email"`
	Username 			string	`json:"username"`
	JoinedDate			string	`json:"date_joined"`
	BloodType			string	`json:"blood_type"`
	Latitude			float32	`json:"latitude"`
	Longitude			float32	`json:"longitude"`
	VerifiedAccount		int		`json:"verified_account"`
}

func (resp SignUp) GetJSONResponseData() ([]byte, int) {
	jsonData, _ := json.Marshal(resp)
	return jsonData, http.StatusOK
}