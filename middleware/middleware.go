package middleware

import (
	// "SimpleDouyin/module"
	"SimpleDouyin/module"
	"SimpleDouyin/repository"
	// "strconv"

	"github.com/gin-gonic/gin"
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
