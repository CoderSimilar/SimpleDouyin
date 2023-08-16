package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response{StatusCode: 1, StatusMsg: "Illegal video-id"},
			Comment{},
		})
		return
	}
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	ci := c.Query("comment_id")
	commentId, err := strconv.Atoi(ci)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response{StatusCode: 1, StatusMsg: "Illegal comment-id"},
			Comment{},
		})
	}

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			Comment{},
		})
		return
	}

	switch actionType {
	case "1":
		comment, err := createComment(videoId, commentText)
		if err != nil {
			c.JSON(http.StatusBadRequest, CommentActionResponse{
				Response{StatusCode: 1, StatusMsg: "Comment failed"},
				Comment{},
			})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response{StatusCode: 0, StatusMsg: "Comment success"},
				comment,
			})
		}
	case "2":
		err := deleteComment(videoId, commentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, CommentActionResponse{
				Response{StatusCode: 1, StatusMsg: "Delete fail"},
				Comment{},
			})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response{StatusCode: 0, StatusMsg: "Delete success"},
				Comment{},
			})
		}
	default:
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response{StatusCode: 1, StatusMsg: "Illegal action-type"},
			Comment{},
		})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	videoId := c.Query("video_id")

	var comments []Comment
	result := db.First(&comments, videoId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, CommentListResponse{
			Response{StatusCode: 1, StatusMsg: result.Error.Error()},
			comments,
		})
	}
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreateDate.After(comments[j].CreateDate)
	})
	c.JSON(http.StatusOK, CommentListResponse{
		Response{StatusCode: 0, StatusMsg: "List comment success"},
		comments,
	})
}

func createComment(videoId int, commentText string) (Comment, error) {
	comment := Comment{
		VideoId:    strconv.Itoa(videoId),
		Content:    commentText,
		ActionType: "1",
		CreateDate: time.Now(),
	}
	var video Video
	result := db.First(&video, videoId)
	if result.Error != nil {
		return comment, result.Error
	}

	video.CommentCount++
	err := db.Create(&comment).Error
	if err != nil {
		return comment, err
	}
	result = db.Save(&video)
	if result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}

func deleteComment(videoId int, commentId int) error {
	err := db.Where("VideoId = ? AND CommentId = ?1", videoId, commentId).Delete(&Comment{}).Error
	return err
}
