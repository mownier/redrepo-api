package entries

type AccountSetting struct {
	Id 						int32 	`db:"id"`
	AccountId 				int32 	`db:"account_id"`
	ConnectedToFacebook 	bool 	`db:"connected_to_facebook"`
	ConnectedToTwitter 		bool 	`db:"connected_to_twitter"`
	VerifiedAccount 		bool 	`db:"verified_account"`
}
