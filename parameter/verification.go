package parameter

type Verification struct {
	BaseParam
	Username 	string `json:"username"`
	Code 		string `json:"code"`
}