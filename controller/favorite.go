package controller

import (
	"SimpleDouyin/demoData"
	"SimpleDouyin/module"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	c.JSON(http.StatusOK, module.Response{
		StatusCode: 0,
		StatusMsg: "successfully",
	})
	// token := c.Query("token")

	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, module.Response{StatusCode: 0})
	// } else {
	// 	c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// }
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: module.Response{
			StatusCode: 0,
		},
		VideoList: demoData.DemoVideos,
	})
}
