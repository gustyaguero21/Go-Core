package svc

import (
	"fmt"
)

type Error struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

func (er *Error) Error() string {
	return fmt.Sprintf("%s, %s", er.Message, er.Err)
}

func NewError(message, err string) error {
	errorStruct := &Error{
		Message: message,
		Err:     err,
	}

	return errorStruct
}
