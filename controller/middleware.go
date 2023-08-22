package controller

import (
	"github.com/gin-gonic/gin"
	"SimpleDouyin/module"
	"SimpleDouyin/repository"
)

func GetCurrentUserId(c *gin.Context) (uId int64, err error) {
	uid, ok := c.Get(module.CurUserId)
	if !ok {
		err = repository.ErrorUserNotLogin
		return
	}
	uId, ok = uid.(int64)
	if !ok {
		err = repository.ErrorUserNotLogin
		return
	}
	return
}
