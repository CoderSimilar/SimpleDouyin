package mysql

import (
	"SimpleDouyin/module"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func InsertComment(commentRecord *module.Comment) (err error) {
	// DB.Migrator().CreateTable(&module.Comment{})
	// DB.Model(&module.Comment{}).Add
	commentRecord.CreatedAt = time.Now()
	commentRecord.CreatedAtString = commentRecord.CreatedAt.Format("2006-01")
	fmt.Println(commentRecord)
	commentRecord.UserId = 3753915153649664
	err = DB.Create(commentRecord).Error
	if err != nil {
		fmt.Println("创建记录出错：", err)
	}
	
	return
}

func DeleteComment(commentRecord *module.Comment) (err error) {
	err = DB.Where("comment_id=?", commentRecord.Id).Delete(&module.Comment{}).Error
	return
}

func UpdateCommentCount(commentRecord *module.Comment) (err error) {
	if commentRecord.ActionType == "1" {
		return addCommentCount(commentRecord)
	}
	return minusCommentCount(commentRecord)
}

func addCommentCount(commentRecord *module.Comment) (err error) {
	video := findCommentCount(commentRecord.VideoId)
	return DB.Model(video).Update("comment_count", video.CommentCount+1).Error
	//DB.Model(voteVideo).Update("favorite_count", favoriteCount+1)
	//fmt.Printf("video.FavoriteCount=%d\n", voteVideo.FavoriteCount)
}

func minusCommentCount(commentRecord *module.Comment) (err error) {
	video := findCommentCount(commentRecord.VideoId)
	return DB.Model(video).Update("comment_count", video.CommentCount-1).Error
}

func findCommentCount(videoId int64) (video *module.Video) {
	video = new(module.Video)
	DB.Model(&module.Video{}).Where("id=?", videoId).First(video)
	return
}

func CommentsQuery(videoId int64) (commentlist *module.CommentList, err error) {
	commentlist = new(module.CommentList)
	commentlist.VideoId = videoId
	var comments []module.Comment
	err = DB.Where("video_id=?", videoId).Order("created_at desc").Find(&comments).Error
	for _, comment := range comments {
		comment.User.Id = comment.UserId
		GetUserInfo(&comment.User)
		fmt.Println(comment)
		commentlist.AllComments = append(commentlist.AllComments, comment)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("video doesn't have comments")
	}
	fmt.Println(commentlist)
	return
}
