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
	re := regexp.MustCompile("[0-9A-Za-z]")
	return re.Match([]byte(value))
}
func (param BaseParam) IsAlpha(value string) bool {
	re := regexp.MustCompile("[A-Za-z]")
	return re.Match([]byte(value))
}