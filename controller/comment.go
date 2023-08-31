package controller

import (
	"SimpleDouyin/middleware"
	"SimpleDouyin/module"
	"SimpleDouyin/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	// 1，获取参数
	commentRecord := new(module.Comment)
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			module.Response{StatusCode: 1, StatusMsg: "Illegal video-id"},
			module.Comment{},
		})
		return
	}
	commentRecord.VideoId = videoId
	actionType := c.Query("action_type")
	commentRecord.ActionType = actionType

	switch actionType {
	case "1":
		commentText := c.Query("comment_text")
		commentRecord.Content = commentText
	case "2":
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64) // 要删除的评论id
		if err != nil {
			c.JSON(http.StatusBadRequest, CommentActionResponse{
				module.Response{StatusCode: 1, StatusMsg: "Illegal comment-id"},
				module.Comment{},
			})
		}
		commentRecord.CommentId = commentId
	default:
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			module.Response{StatusCode: 1, StatusMsg: "Illegal action-type"},
			module.Comment{},
		})

	}

	// 2，获取当前请求用户的id
	userId, err := middleware.GetCurrentUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: "User need login"})
		return
	}
	fmt.Println(userId)
	commentRecord.User.UserId = userId
	commentRecord.UserId = userId
	
	
	// 3，业务处理
	if err = service.CommentAction(commentRecord); err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	// 4，返回响应
	if commentRecord.ActionType == "1" {
		c.JSON(http.StatusOK, CommentActionResponse{Response: module.Response{StatusCode: 0},
			Comment: module.Comment{
				CommentId: commentRecord.CommentId,
				UserId: commentRecord.UserId,
				User: commentRecord.User,
				Content: commentRecord.Content,
				ActionType: "1",
				CreatedAtString: commentRecord.CreatedAtString,
			},
		})
	}
	c.JSON(http.StatusOK, module.Response{StatusCode: 0, StatusMsg: "successfully"})
	

	
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	// 1，获取参数
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64) 
	if err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	// 2，业务处理
	commentList, err := service.CommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	// 3，返回响应
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    module.Response{StatusCode: 0},
		CommentList: commentList.AllComments,
		//CommentList: DemoComments,
	})
	
}

