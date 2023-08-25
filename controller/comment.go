package controller

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
)

type CommentActionResponse struct {
	module.Response
	Comment module.Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	module.Response
	CommentList []module.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, err := module.ParseToken(token); err != nil {
		c.JSON(http.StatusBadRequest, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			module.Response{StatusCode: 1, StatusMsg: "Illegal video-id"},
			module.Comment{},
		})
		return
	}
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	ci := c.Query("comment_id")
	commentId, err := strconv.Atoi(ci)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			module.Response{StatusCode: 1, StatusMsg: "Illegal comment-id"},
			module.Comment{},
		})
	}

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			module.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			module.Comment{},
		})
		return
	}

	switch actionType {
	case "1":
		comment, err := createComment(videoId, commentText)
		if err != nil {
			c.JSON(http.StatusBadRequest, CommentActionResponse{
				module.Response{StatusCode: 1, StatusMsg: "Comment failed"},
				module.Comment{},
			})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				module.Response{StatusCode: 0, StatusMsg: "Comment success"},
				comment,
			})
		}
	case "2":
		err := deleteComment(videoId, commentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, CommentActionResponse{
				module.Response{StatusCode: 1, StatusMsg: "Delete fail"},
				module.Comment{},
			})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				module.Response{StatusCode: 0, StatusMsg: "Delete success"},
				module.Comment{},
			})
		}
	default:
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			module.Response{StatusCode: 1, StatusMsg: "Illegal action-type"},
			module.Comment{},
		})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	videoId := c.Query("video_id")

	var comments []module.Comment
	result := mysql.DB.First(&comments, videoId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, CommentListResponse{
			module.Response{StatusCode: 1, StatusMsg: result.Error.Error()},
			comments,
		})
	}
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].Model.CreatedAt.After(comments[j].Model.CreatedAt)
	})
	c.JSON(http.StatusOK, CommentListResponse{
		module.Response{StatusCode: 0, StatusMsg: "List comment success"},
		comments,
	})
}

func createComment(videoId int, commentText string) (module.Comment, error) {
	comment := module.Comment{
		VideoId:    strconv.Itoa(videoId),
		Content:    commentText,
		ActionType: "1",
	}
	var video module.Video
	result := mysql.DB.First(&video, videoId)
	if result.Error != nil {
		return comment, result.Error
	}

	video.CommentCount++
	err := mysql.DB.Create(&comment).Error
	if err != nil {
		return comment, err
	}
	result = mysql.DB.Save(&video)
	if result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}

func deleteComment(videoId int, commentId int) error {
	err := mysql.DB.Where("VideoId = ? AND CommentId = ?1", videoId, commentId).Delete(&module.Comment{}).Error
	return err
}
