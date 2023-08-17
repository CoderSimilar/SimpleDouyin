package controller

import (
	"github.com/gin-gonic/gin"
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
)

func GetCurrentUserId(c *gin.Context) (uId int64, err error) {
	uid, ok := c.Get(module.CurUserId)
	if !ok {
		err = mysql.ErrorUserNotLogin
		return
	}
	uId, ok = uid.(int64)
	if !ok {
		err = mysql.ErrorUserNotLogin
		return
	}
	return
}
