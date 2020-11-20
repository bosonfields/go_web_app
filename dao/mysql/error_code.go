package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("user already exist")
	ErrorUserNotExist    = errors.New("user not exist")
	ErrorInvalidPassword = errors.New("incorrect password")
	ErrorInvalidID       = errors.New("invalid ID")
)
