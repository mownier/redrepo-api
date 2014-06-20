package errors

import (
	"log"
	"errors"
	"time"
	)

func Log(e error) {
	log.Printf("Error [%+v]: %+v\n", time.Now().String(), e)
}

func LogErrorMessage(message string) {
	Log(errors.New(message))
}
