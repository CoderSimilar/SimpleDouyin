package mysql

import (
	"SimpleDouyin/module"
	"crypto/md5"
	"encoding/hex"
	"SimpleDouyin/repository"
	"gorm.io/gorm"
)

const secret = "SimpleDouyin"

// checkUserExist 判断用户是否存在
func CheckUserExist(username string) (err error) {
	// 定义一个具体的user
	err = DB.Where("name=?", username).First(&module.User{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	// first查询不到返回错误，而find不会返回错误
	return repository.ErrorUserExist
}

func InsertUser(newUser *module.User) (err error) {
	// 对密码进行加密
	newUser.Password = encryptPassword(newUser.Password)

	// 新用户插入数据库
	return DB.Create(newUser).Error
}

func CheckLoginUser(user *module.User) (err error) {
	user.Password = encryptPassword(user.Password)
	// 定义一个具体的user
	err = DB.Where("name=? and password=?", user.Name, user.Password).First(user).Error
	return
}

func GetUserInfo(user *module.User) (err error) {
	
	return DB.Where("user_id=?", user.UserId).Find(user).Error
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New() 
	h.Write([]byte(secret)) // 加盐的字符串
	return hex.EncodeToString(h.Sum([]byte(oPassword))) // 字节 转 十六进制
}
