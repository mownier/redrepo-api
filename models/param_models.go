// param_models.go
// @author Mounir Ybanez
// @date May 24, 2014

package models

type SignInParam struct {
	BaseParamModel
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type SignUpParam struct {
	BaseParamModel
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
	Username 	string	`json:"username"`
	BloodType	string	`json:"blood_type"`
	Latitude	float32	`json:"latitude"`
	Longitude	float32	`json:"longitude"`
}

type AccountSettingParam struct {
	SignUpParam
	ConnectedToFacebook 	int `json:"connected_to_facebook"`
	ConnectedToTwitter 		int `json:"connected_to_twitter"`
}

type VerificationParam struct {
	BaseParamModel
	Username 	string `json:"username"`
	Code 		string `json:"code"`
}
