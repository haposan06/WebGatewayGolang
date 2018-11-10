package errors

import (
	"errors"
	"fmt"
)

var ErrNotFound 		= errors.New("No data found in storage")
var ErrDBOpsFailed		= errors.New("DB operation failed")
var ErrPrimarKey		= errors.New("Not valid primary key")
var ErrSystemInternal	= errors.New("System Internal Error")
var ErrUnauthorized		= errors.New("Unauthorized access")
var ErrInvalidToken		= errors.New("Invalid token")

type Error struct {
	Code	string
	Message	string
	Detail 	interface{}
}

func (e Error) Error() string{
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}