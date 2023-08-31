package mysql

import (
	"errors"
	"fmt"
	// "fmt"
	"SimpleDouyin/module"
)

// 用户没有登录时，直接返回
func FeedVideos() (videos *[]module.Video, err error) {
	// 视频按时间顺序逆置显示，最多显示30个
	videos = new([]module.Video)
	err = DB.Order("created_at desc").Limit(30).Find(videos).Error
	for index := range *videos {
		(*videos)[index].Author.Id = (*videos)[index].AuthorId
		GetUserInfo(&(*videos)[index].Author)
	}
	// fmt.Println(videos)
	if err != nil { // 数据库中没有视频
		return nil, errors.New("no videos")
	}
	return
}

// 用户登录，将他喜欢的视频的is_favorited设置为true
func CheckLikeVideo(videos *[]module.Video, userId int64) (videosWithLike *[]module.Video, err error) {
	// 在user_video_relation表中找出用户喜欢的视频
	// var videoWithLike module.Video
	var favoriteVideoIDs []int64
	if err := DB.Model(&module.UserVideoRelation{}).Select("video_id").Where("user_id = ? AND is_favorite = ?", userId, true).Pluck("video_id", &favoriteVideoIDs).Error; err != nil {
		// 错误处理
		return nil, err
	}
	fmt.Println(favoriteVideoIDs)
	for i := range *videos {
		for _, favoriteID := range favoriteVideoIDs {
			if (*videos)[i].Id == favoriteID {
				(*videos)[i].IsFavorite = true
				break
			}
		}
	}
	return videos, nil

}
