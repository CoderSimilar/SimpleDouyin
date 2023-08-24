package module

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		mc, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 401, StatusMsg: "Invalid token"},
			})

			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(CurUserId, mc.UserId)
		c.Next() // 后续的处理请求函数中 可以使用c.Get(CtxtUserIDKey)来获取当前请求的用户信息
	}
}
