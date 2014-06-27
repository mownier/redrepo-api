package tables

type AuthClient struct {
	Id			int		`db:"id"`
	Key			string	`db:"key"`	
	Secret		string	`db:"secret"`
	Name		string	`db:"name"`
	Description	string	`db:"description"`
}
