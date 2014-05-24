// output_models.go
// @author Mounir Ybanez
// @date May 24, 2014

package models

type UserOutputModel struct {
	BaseOutputModel
	Id 					string	`json:"id"`
	FirstName 			string	`json:"first_name"`
	LastName 			string	`json:"last_name"`
	Email				string	`json:"email"`
	Username 			string	`json:"username"`
	JoinedDate			string	`json:"date_joined"`
	Latitude			float32	`json:"latitude"`
	Longitude			float32	`json:"longitude"`
	ConnectedToFacebook int 	`json:"connected_to_facebook"`
	ConnectedToTwitter	int 	`json:"connected_to_twitter"`
}

type AuthenticatedUserOutputModel struct {
	BaseOutputModel
	User 		*UserOutputModel 	`json:"user"`
	AccessToken	string				`json:"access_token"`		
}