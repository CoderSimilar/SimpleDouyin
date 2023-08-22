package repository

import "errors"

var (
	// 用户
	ErrorUserExist = errors.New("user already exist")
	ErrorUserInfo  = errors.New("user doesn't exist or Error password")

	ErrorUserNotLogin = errors.New("user doesn't login")
	ErrorRegister     = errors.New("user registration failed")

	// 视频
	ErrorVideoExist = errors.New("video already exist")

	ErrorInvalidVideoFormat = errors.New("invalid Video Format")

	ErrorGenPicture = errors.New("picture generate error")
)
