package controller

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
	"net/http"
	"strconv"
	"time"

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
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, module.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
	}
	actionType := c.Query("action_type")

	switch actionType {
	case "1":
		err = like(videoId)
		if err != nil {
			c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		}
	case "2":
		err = cancelLike(videoId)
		if err != nil {
			c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		}
	default:
		c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: "Illegal action-type"})
	}
}

// LikeList all users have same like video list
func LikeList(c *gin.Context) {
	userId := c.Query("user_id")

	var relationList []module.UserVideoRelation
	var likedVideoList []module.Video

	result := mysql.DB.Find(&relationList, "UserId = ? AND IsLiked = ?", userId, true)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: result.Error.Error()})
	}
	for _, r := range relationList {
		var video module.Video
		result = mysql.DB.Find(&video, "VideoId = ?", r.VideoId)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: result.Error.Error()})
		}
		likedVideoList = append(likedVideoList, video)
	}
	c.JSON(http.StatusOK, LikeListResponse{module.Response{StatusCode: 0, StatusMsg: "List success"}, likedVideoList})
}

func like(videoId int) error {
	var video module.Video
	var relation module.UserVideoRelation
	result := mysql.DB.First(&video, videoId)
	if result.Error != nil {
		return result.Error
	}
	result = mysql.DB.First(&relation, videoId)
	if result.Error != nil {
		return result.Error
	}

	video.FavoriteCount++
	// video.UpdateTime = time.Now()
	relation.IsLiked = true
	relation.UpdateDatetime = time.Now()

	result = mysql.DB.Save(&video)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func cancelLike(videoId int) error {
	var video module.Video
	var relation module.UserVideoRelation
	result := mysql.DB.First(&video, videoId)
	if result.Error != nil {
		return result.Error
	}
	result = mysql.DB.First(&relation, videoId)
	if result.Error != nil {
		return result.Error
	}

	video.FavoriteCount--
	// video.UpdateTime = time.Now()
	relation.IsLiked = false
	relation.UpdateDatetime = time.Now()

	result = mysql.DB.Save(&video)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
