package service

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
)

func CommentAction(commentRecord *module.Comment) (err error) {
	// 1，写入评论
	if commentRecord.ActionType == "1" {
		commentRecord.Id = module.GenID()
		err = mysql.GetUserInfo(&commentRecord.User)
		if err != nil {
			return err
		}
		err = mysql.InsertComment(commentRecord)
		if err != nil {
			return err
		}
	}

	// 2，删除评论
	if commentRecord.ActionType == "2" {
		err = mysql.DeleteComment(commentRecord)
		return
	}

	// 3，更新video的CommentCount
	return mysql.UpdateCommentCount(commentRecord)
}

func CommentList(videoId int64) (commentlist *module.CommentList, err error) {
	commentlist,err = mysql.CommentsQuery(videoId)
	return 
}
