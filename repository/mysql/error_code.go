package mysql

import "errors"

var (
	ErrorUserExist = errors.New("user already exist")
	ErrorUserInfo  = errors.New("user doesn't exist or Error password")

	ErrorUserNotLogin = errors.New("user doesn't login")
	ErrorRegister     = errors.New("user registration failed")
)
