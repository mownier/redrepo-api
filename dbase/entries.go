// entries.go
// @author Mounir Ybanez
// @date June 8, 2014

package dbase

type BaseEntry struct {
	Id int64 `db:"id"`
}

type Account struct {
	BaseEntry
	FirstName 	string 	`db:"first_name"`
	LastName 	string 	`db:"last_name"`
	Email 		string 	`db:"email"`
	Password	string
	Username 	string 	`db:"username"`
	BloodType 	string 	`db:"blood_type"`
	Latitude 	float64	`db:"latitude"`
	Longitude 	float64	`db:"longitude"`
	DateJoined	string	`db:"date_joined"`
}