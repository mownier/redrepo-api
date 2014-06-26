package joins

type AccountVerificationJoinResult struct {
	Id			int		`db:"id"`
	Code		string	`db:"code"`
	Status 		string	`db:"status"`
	Email 		string	`db:"email"`
	Username	string	`db:"username"`
}