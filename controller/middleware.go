package controller

import (
	// "SimpleDouyin/module"
	"SimpleDouyin/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCurrentUserId(c *gin.Context) (uId int64, err error) {
	uid := c.Query("user_id")
	uId, err = strconv.ParseInt(uid, 10, 64)
	if err != nil {
		err = repository.ErrorUserNotLogin
		return
	}
	// if !ok {
	// 	err = repository.ErrorUserNotLogin
	// 	return
	// }
	// uId, ok = uid.(int64)
	// if !ok {
	// 	err = repository.ErrorUserNotLogin
	// 	return
	// }
	return
}
