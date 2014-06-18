package parameter

import (
	"regexp"
	)

type BaseParam struct {
	
}

func (param BaseParam) HasErrors() bool {
	return true
}

func (param BaseParam) IsEmpty(value string) bool {
	if len(value) == 0 { 
		return true 
	} else { 
		return false 
	}
}

func (param BaseParam) IsAlphaNumeric(value string) bool {
	_, err := regexp.Compile("[^A-Za-z0-9]+")
    if err == nil { 
    	return true 
    } else { 
    	return false 
    }
}