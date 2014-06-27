package tables

type AuthSession struct {
	Id 					int		`db:"id"`
	AccountId			int		`db:"account_id"`
	AccessToken			string	`db:"access_token"`
	ClientKey			string	`db:"client_key"`
	AuthProvider 		string	`db:"auth_provider"`
	AuthProviderToken	string	`db:"auth_provider_token"`
}
