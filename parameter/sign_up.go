package parameter

import (
	"regexp"
	"fmt"
	)

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
			param.isValidBloodType() &&
			param.isValidCoordinate())
}

func (param SignUp) isValidFirstName() bool {
	valid := !param.BaseParam.IsEmpty(param.FirstName) && 
				param.BaseParam.IsAlpha(param.FirstName)
	
	if valid == false {
		fmt.Println("Invalid first name.")
	}
	return valid
}

func (param SignUp) isValidLastName() bool {
	valid := !param.BaseParam.IsEmpty(param.LastName) && 
				param.BaseParam.IsAlpha(param.LastName)
	if valid == false {
		fmt.Println("Invalid last name.")
	}
	return valid
}

func (param SignUp) isValidEmail() bool {
	re := regexp.MustCompile(".+@.+\\..+")
	valid := re.Match([]byte(param.Email))
	if valid == false {
		fmt.Println("Invalid email.")
	}
	return valid
}

func (param SignUp) isValidUserName() bool {
	valid := !param.BaseParam.IsEmpty(param.FirstName) && 
				param.BaseParam.IsAlphaNumeric(param.FirstName)
	if valid == false {
		fmt.Println("Invalid username.")
	}
	return valid
}

func (param SignUp) isValidBloodType() bool {
	return true
}

func (param SignUp) isValidCoordinate() bool {
	valid := (param.Latitude >= -90.0 && param.Latitude <= 90.0) &&
				(param.Longitude >= -180.0 && param.Longitude <= 180.0)
	if valid == false {
		fmt.Println("Invalid coordinate.")
	}
	return valid
}



