package controller

import (
	"SimpleDouyin/module"
	"SimpleDouyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	module.Response
	VideoList []module.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 1.获取参数并验证参数
	video := new(module.Video)
	video.Title = c.PostForm("title")

	userId, err := GetCurrentUserId(c) // 获取视频发布者的ID
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  "User need login",
		})
	}
	video.AuthorId = userId

	// 2.业务处理
	if err := service.Publish(video, c); err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  "Upload error",
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, module.Response{
		StatusCode: 0,
		StatusMsg:  "Work is uploaded successfully",
	})
}

// PublishList all users have same publish video list
// PublishList 查询到user的所有视频，使用列表返回
func PublishList(c *gin.Context) {
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

	// 2.业务处理
	data, err := service.PublishList(userId)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, VideoListResponse{
		Response: module.Response{
			StatusCode: 0,
		},
		VideoList: data.AllVideoes,
	})
}
