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
		if token == "" {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 404, StatusMsg: "User didn't login!"},
			})
			c.Abort()
			return
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

func AuthFeedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户token
		tokenString := c.Query("token")
		if tokenString == "" {
			// 用户未登录
			return
		}
		mc, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 401, StatusMsg: "Invalid token"},
			})
			c.Abort()
		}
		c.Set(CurUserId, mc.UserId)
		c.Next() 
	}	
}
