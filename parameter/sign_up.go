package parameter

type SignUp struct {
	BaseParam
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
	Username 	string	`json:"username"`
	BloodType	string	`json:"blood_type"`
	Latitude	float32	`json:"latitude"`
	Longitude	float32	`json:"longitude"`
}

func (param SignUp) HasErrors() bool {
	return !(param.isValidFirstName() && 
			param.isValidLastName() && 
			param.isValidEmail() && 
			param.isValidUserName() && 
			param.isValidBloodType())
}

func (param SignUp) isValidFirstName() bool {
	return !param.BaseParam.IsEmpty(param.FirstName) && 
			param.BaseParam.IsAlphaNumeric(param.FirstName)
}

func (param SignUp) isValidLastName() bool {
	return true
}

func (param SignUp) isValidEmail() bool {
	return true
}

func (param SignUp) isValidUserName() bool {
	return true
}

func (param SignUp) isValidBloodType() bool {
	return true
}



