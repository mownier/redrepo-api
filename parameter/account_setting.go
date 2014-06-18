package parameter

type AccountSetting struct {
	SignUp
	ConnectedToFacebook 	int `json:"connected_to_facebook"`
	ConnectedToTwitter 		int `json:"connected_to_twitter"`
}