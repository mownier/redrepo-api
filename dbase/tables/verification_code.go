package tables

type VerificationCode struct {
	Id			int		`db:"id"`
	Username 	string 	`db:"username"`
	Code 		string 	`db:"code"`
	Status		string	`db:"status"`
}