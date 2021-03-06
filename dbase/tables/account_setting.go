package tables

type AccountSetting struct {
	Id 						int 	`db:"id"`
	Username 				string 	`db:"username"`
	ConnectedToFacebook 	bool 	`db:"connected_to_facebook"`
	ConnectedToTwitter 		bool 	`db:"connected_to_twitter"`
	Verified 				bool 	`db:"verified"`
}
