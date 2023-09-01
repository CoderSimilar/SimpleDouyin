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
	token := c.Query("token")
	// 如果token不是空，代表用户已经登录，需要做身份验证
	if token != "" {
		// 获取用户id
		userId, err := middleware.GetCurrentUserId(c)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response:  module.Response{StatusCode: 1, StatusMsg: "Fail to GetCurrentUserId"},
				VideoList: demoData.DemoVideos,
				NextTime:  time.Now().Unix(),
			})
			return
		}
		video_list, err := service.Feed(userId, token)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response:  module.Response{StatusCode: 1, StatusMsg: "Fail to Get Videos!"},
				VideoList: demoData.DemoVideos,
				NextTime:  time.Now().Unix(),
			})
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  module.Response{StatusCode: 0},
			VideoList: *video_list,
			NextTime:  time.Now().Unix(),
		})
	}else {
		// 用户未登录
		video_list, err := service.Feed(0, "")
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response:  module.Response{StatusCode: 1, StatusMsg: "Fail to Get Videos!"},
				VideoList: demoData.DemoVideos,
				NextTime:  time.Now().Unix(),
			})
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  module.Response{StatusCode: 0},
			VideoList: *video_list,
			NextTime:  time.Now().Unix(),
		})
	}
	
	
	
}
