package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type LikeActionResponse struct {
	Response
}

type LikeListResponse struct {
	Response
	VideoList []Video
}

// LikeAction has no practical effect, just check if token is valid
func LikeAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: err.Error()})
	}
	actionType := c.Query("action_type")

	switch actionType {
	case "1":
		err = like(videoId)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: err.Error()})
		}
	case "2":
		err = cancelLike(videoId)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: err.Error()})
		}
	default:
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "Illegal action-type"})
	}
}

// LikeList all users have same like video list
func LikeList(c *gin.Context) {
	userId := c.Query("user_id")

	var relationList []UserVideoRelation
	var likedVideoList []Video

	result := db.Find(&relationList, "UserId = ? AND IsLiked = ?", userId, true)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: result.Error.Error()})
	}
	for _, r := range relationList {
		var video Video
		result = db.Find(&video, "VideoId = ?", r.VideoId)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: result.Error.Error()})
		}
		likedVideoList = append(likedVideoList, video)
	}
	c.JSON(http.StatusOK, LikeListResponse{Response{StatusCode: 0, StatusMsg: "List success"}, likedVideoList})
}

func like(videoId int) error {
	var video Video
	var relation UserVideoRelation
	result := db.First(&video, videoId)
	if result.Error != nil {
		return result.Error
	}
	result = db.First(&relation, videoId)
	if result.Error != nil {
		return result.Error
	}

	video.LikeCount++
	video.UpdateDatetime = time.Now()
	relation.IsLiked = true
	relation.UpdateDatetime = time.Now()

	result = db.Save(&video)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func cancelLike(videoId int) error {
	var video Video
	var relation UserVideoRelation
	result := db.First(&video, videoId)
	if result.Error != nil {
		return result.Error
	}
	result = db.First(&relation, videoId)
	if result.Error != nil {
		return result.Error
	}

	video.LikeCount--
	video.UpdateDatetime = time.Now()
	relation.IsLiked = false
	relation.UpdateDatetime = time.Now()

	result = db.Save(&video)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
