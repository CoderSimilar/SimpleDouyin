package controller

import (
	"SimpleDouyin/middleware"
	"SimpleDouyin/module"
	"SimpleDouyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LikeActionResponse struct {
	module.Response
}

type LikeListResponse struct {
	module.Response
	VideoList []module.Video
}

// LikeAction has no practical effect, just check if token is valid
func LikeAction(c *gin.Context) {

	// 1，获取参数
	// 将video_id转换成int
	newRelation := new(module.UserVideoRelation)
	var err error
	newRelation.VideoId, err = strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
	}

	// 根据传入的actionType决定点赞或者取消点赞
	actionType := c.Query("action_type")
	if actionType == "1" {
		newRelation.IsFavorite = true
	} else if actionType == "2" {
		newRelation.IsFavorite = false
	} else {
		c.JSON(http.StatusOK, module.UserResponse{
			Response: module.Response{StatusCode: 1, StatusMsg: "Parameter out of bounds"},
		})
		return
	}

	// 2, 获取当前请求的用户id
	userId, err := middleware.GetCurrentUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, module.UserResponse{
			Response: module.Response{StatusCode: 1, StatusMsg: "User need login"},
		})
		return
	}
	newRelation.UserId = userId

	// 3，业务处理
	if err := service.LikeAction(newRelation); err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: err.Error()})
	}

	// 4，返回响应
	c.JSON(http.StatusOK, module.Response{StatusCode: 0, StatusMsg: "successfully"})
}

// LikeList all users have same like video list
func LikeList(c *gin.Context) {

	// 1.获取userId
	curUserId := c.Query("user_id")
	userId, err := strconv.ParseInt(curUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 2，业务处理
	data, err := service.LikeList(userId)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 3，返回响应

	c.JSON(http.StatusOK, VideoListResponse{
		Response: module.Response{
			StatusCode: 0,
			StatusMsg: "successfully",
		},
		VideoList: data.AllVideos,
	})
	
}




