package controller

import (
	"SimpleDouyin/demoData"
	"SimpleDouyin/middleware"
	"SimpleDouyin/module"
	"SimpleDouyin/service"
	"fmt"

	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	module.Response
	VideoList []module.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	fmt.Println("Welcome to Douyin!")
	userId, err := middleware.GetCurrentUserId(c)
	fmt.Println(userId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  module.Response{StatusCode: 0},
			VideoList: demoData.DemoVideos,
			NextTime:  time.Now().Unix(),
		})
		return
	}
	token := c.Query("token")
	video_list, err := service.Feed(userId, token)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  module.Response{StatusCode: 1, StatusMsg: "Fail to Get Videos!"},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  module.Response{StatusCode: 0},
		VideoList: *video_list,
		NextTime:  time.Now().Unix(),
	})
}
