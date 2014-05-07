// models.go
// @author Mounir Ybanez
// @date May 7, 2014

package services

import ( 
	"code.google.com/p/gorest"
)

type BaseService struct {
	gorest.RestService `root:"/api/v1/" consumes:"application/json" produces:"application/json"`
}

type User struct {
	Id 					string	`json:"id"`
	FirstName 			string	`json:"first_name"`
	LastName 			string	`json:"last_name"`
	Email				string	`json:"email"`
	Password 			string	`json:"password"`
	Username 			string	`json:"username"`
	JoinedDate			string	`json:"date_joined"`
	Latitude			float32	`json:"latitude"`
	Longitude			float32	`json:"longitude"`
	ConnectedToFacebook bool 	`json:"connected_to_facebook"`
	ConnectedToTwitter	bool 	`json:"connected_to_twitter"`
}