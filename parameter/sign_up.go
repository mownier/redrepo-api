package parameter

import (
	"regexp"
	"errors"
	// "reflect"
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

var bloodTypes = []string {
	"A", "A+", "A-",
	"O", "O+", "O-",
	"AB", "AB+", "A-",
}


func (param SignUp) ValidateValues() error {
	error := param.validateFirstName()
	if error == nil {
		error = param.validateLastName()
		if error == nil {
			error = param.validateEmail()
			if error == nil {
				error = param.validateUsername()
				if error == nil {
					error = param.validateBloodType()
					if error == nil {
						error = param.validateCoordinate()
					}
				}
			}
		}
	}
	return error
}

func (param SignUp) validateFirstName() error {
	valid := !param.BaseParam.IsEmpty(param.FirstName) && 
				param.BaseParam.IsAlpha(param.FirstName)
	if valid == true {
		return nil
	} else {
		return errors.New("Invalid first name.")
	}
}

func (param SignUp) validateLastName() error {
	valid := !param.BaseParam.IsEmpty(param.LastName) && 
				param.BaseParam.IsAlpha(param.LastName)
	if valid == true {
		return nil
	} else {
		return errors.New("Invalid last name.")
	}
}

func (param SignUp) validateEmail() error {
	re := regexp.MustCompile(".+@.+\\..+")
	valid := re.Match([]byte(param.Email))
	if valid == true {
		return nil
	} else {
		return errors.New("Invalid email.")
	}
}

func (param SignUp) validateUsername() error {
	valid := !param.BaseParam.IsEmpty(param.FirstName) && 
				param.BaseParam.IsAlphaNumeric(param.FirstName)
	if valid == true {
		return nil
	} else {
		return errors.New("Invalid username.")
	}
}

func (param SignUp) validateBloodType() error {
	valid := false
	for i := 0; i < len(bloodTypes); i++ {
		bloodType := bloodTypes[i]
		if bloodType == param.BloodType {
			valid = true
			break
		}
	}

	if valid == true {
		return nil
	} else {
		return errors.New("Invalid blood type.")
	}
}

func (param SignUp) validateCoordinate() error {
	valid := (param.Latitude >= -90.0 && param.Latitude <= 90.0) &&
				(param.Longitude >= -180.0 && param.Longitude <= 180.0)
	if valid == true {
		return nil
	} else {
		return errors.New("Invalid coordinate.")
	}
}



