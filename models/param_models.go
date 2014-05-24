// param_models.go
// @author Mounir Ybanez
// @date May 24, 2014

package models

type SignInParamModel struct {
	BaseParamModel
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type UserParamModel struct {
	BaseParamModel
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
	Username 	string	`json:"username"`
	Latitude	float32	`json:"latitude"`
	Longitude	float32	`json:"longitude"`
}
