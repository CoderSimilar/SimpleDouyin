package controller

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository"
	"SimpleDouyin/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]module.User{
	"zhangleidouyin": {
		UserId: 1,
		Name:   "zhanglei",
		//FollowCount:   10,
		//FollowerCount: 5,
		//IsFollow:      true,
	},
}

// var userIdSequence = int64(1)

func Register(c *gin.Context) {
	// 1.获取参数和参数校验
	username := c.Query("username")
	password := c.Query("password")
	if len(username) == 0 || len(password) < 5 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "parameter error",
		})
		return
	}

	// 2.业务处理
	newUser, err := service.Register(username, password)

	if err != nil {
		// 用户不存在
		if errors.Is(err, repository.ErrorUserExist) {
			c.JSON(http.StatusOK, module.UserLoginResponse{
				Response: module.Response{StatusCode: 1, StatusMsg: "User already exist"},
			})
			return
		}
		// 注册失败
		if errors.Is(err, repository.ErrorRegister) {
			c.JSON(http.StatusOK, module.UserLoginResponse{
				Response: module.Response{StatusCode: 1, StatusMsg: "User registration failed"},
			})
			return
		}
		// 其他错误
		c.JSON(http.StatusOK, module.UserLoginResponse{
			Response: module.Response{StatusCode: 1, StatusMsg: "Server Busy"},
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, module.UserLoginResponse{
		// 注册成功
		Response: module.Response{StatusCode: 0},
		UserId:   newUser.UserId,
		Token:    newUser.Token,
	})
}

func Login(c *gin.Context) {
	// 1.获取参数并验证
	username := c.Query("username")
	password := c.Query("password")
	if len(username) == 0 || len(password) < 5 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "parameter error",
		})
		return
	}

	// 2.业务处理
	user, err := service.Login(username, password)
	if err != nil {
		// 用户不存在
		if errors.Is(err, repository.ErrorUserInfo) {
			c.JSON(http.StatusOK, module.UserLoginResponse{
				Response: module.Response{StatusCode: 1, StatusMsg: "User doesn't exist or Error password"},
			})
			return
		}

		// 其他错误
		c.JSON(http.StatusOK, module.UserLoginResponse{
			Response: module.Response{StatusCode: 1, StatusMsg: "Server Busy"},
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, module.UserLoginResponse{
		Response: module.Response{StatusCode: 0},
		UserId:   user.UserId,
		Token:    user.Token,
	})

}

func UserInfo(c *gin.Context) {
	// 1.获取当前用户的id
	userId, err := GetCurrentUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, module.UserResponse{
			Response: module.Response{StatusCode: 1, StatusMsg: "User need login"},
		})
		return
	}

	// 2.业务逻辑
	user, err := service.UserInfo(userId)

	if err != nil {
		c.JSON(http.StatusOK, module.UserResponse{
			Response: module.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}

	c.JSON(http.StatusOK, module.UserResponse{
		Response: module.Response{StatusCode: 0},
		User:     *user,
	})
}
