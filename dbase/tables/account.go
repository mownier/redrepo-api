package tables

type Account struct {
	Id			int		`db:"id"`
	FirstName 	string 	`db:"first_name"`
	LastName 	string 	`db:"last_name"`
	Email 		string 	`db:"email"`
	Password	string	`db:"password"`
	Username 	string 	`db:"username"`
	BloodType 	string 	`db:"blood_type"`
	Latitude 	float32	`db:"latitude"`
	Longitude 	float32	`db:"longitude"`
	DateJoined	string	`db:"date_joined"`
}