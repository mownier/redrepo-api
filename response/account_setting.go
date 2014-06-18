package response

type AccountSetting struct {
	Account
	ConnectedToFacebook int `json:"connected_to_facebook"`
	ConnectedToTwitter	int `json:"connected_to_twitter"`
}