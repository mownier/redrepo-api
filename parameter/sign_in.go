package parameter

type SignIn struct {
	BaseParam
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}