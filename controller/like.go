package controller

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	token := c.Query("token")

	// 检测用户token是否legal
	if _, err := module.ParseToken(token); err != nil {
		c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	// 检测视频Id是否合法
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
	}
	actionType := c.Query("action_type")

	switch actionType {
	case "1":
		err = like(token, videoId)
		if err != nil {
			c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		}
	case "2":
		err = cancelLike(token, videoId)
		if err != nil {
			c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		}
	default:
		c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: "Illegal action-type"})
	}
}

// LikeList all users have same like video list
func LikeList(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	videoList := []module.Video{}

	var relations []module.UserVideoRelation
	if err := mysql.DB.Where("token = ? AND video_id = ?", token, videoId).Find(&relations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, module.Response{StatusCode: 1, StatusMsg: "Error fetching data"})
		return
	}

	for _, relation := range relations {
		var video module.Video
		if err := mysql.DB.First(&video, relation.VideoId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, module.Response{StatusCode: 1, StatusMsg: "Error fetching video data"})
			return
		}
		videoList = append(videoList, video)
	}

	c.JSON(http.StatusOK, LikeListResponse{Response: module.Response{StatusCode: 0, StatusMsg: "List success"}, VideoList: videoList})
}

func like(token string, videoId int) error {

	// 创建 UserVideoRelation 记录
	newRelation := module.UserVideoRelation{
		Token:   token,
		VideoId: videoId,
	}

	// 在数据库中创建记录
	result := mysql.DB.Create(&newRelation)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func cancelLike(token string, videoId int) error {
	var relation module.UserVideoRelation

	// 根据 token 和 videoId 查询 user_video_relation 表
	if err := mysql.DB.Where("token = ? AND video_id = ?", token, videoId).First(&relation).Error; err != nil {
		return err
	}

	// 删除点赞记录
	if err := mysql.DB.Delete(&relation).Error; err != nil {
		return err
	}

	return nil
}
