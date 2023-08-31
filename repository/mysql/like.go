package mysql

import (
	"SimpleDouyin/module"
	"errors"

	"gorm.io/gorm"
)

func CheckLikeExist(userId, videoId int64) (exists bool, err error) {
	// 如果在点赞记录中发现该视频不存在或者用户没有对该视频点过赞，返回
	var count int64
	result := DB.Model(&module.UserVideoRelation{}).Where("user_id = ? AND video_id = ? AND is_favorite = ?", userId, videoId, true).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func LikeHandle(newRelation *module.UserVideoRelation) (err error) {
	if newRelation.IsFavorite {
		return createLike(newRelation)
	}
	return deleteLike(newRelation)
}

// 添加点赞视频
func createLike(newRelation *module.UserVideoRelation) (err error) {
	return DB.Create(newRelation).Error
}

// 取消赞的视频
func deleteLike(newRelation *module.UserVideoRelation) (err error) {

	var relation module.UserVideoRelation
	
	// 查找点赞记录
	if err = DB.Where("user_id = ? AND video_id = ?", newRelation.UserId, newRelation.VideoId).First(&relation).Error; err != nil {
		return err
	}

	// 设置is_favorited为false
	if err = DB.Model(&module.UserVideoRelation{}).Where("user_id = ? AND video_id = ? AND deleted_at IS NOT NULL", newRelation.UserId, newRelation.VideoId).
	Update("is_favorite", false).Error; err != nil {
		return err
	}

	// 删除点赞记录
	if err := DB.Delete(&relation).Error; err != nil {
		return err
	}

	return nil
}

// 更新视频点赞数量
func UpdateFavotiteCount(newRelation *module.UserVideoRelation) (err error) {
	// 更新 video 类下的 FavoriteCount
	if newRelation.IsFavorite {
		return addLike(newRelation)
	}
	return minusLike(newRelation)
}
// 视频点赞数 + 1
func addLike(newRelation *module.UserVideoRelation) (err error) {
	video := findFavoriteCount(newRelation.VideoId)
	return DB.Model(video).Update("favorite_count", video.FavoriteCount + 1).Error
}
// 视频点赞数 - 1
func minusLike(newRelation *module.UserVideoRelation) (err error) {
	video := findFavoriteCount(newRelation.VideoId)
	return DB.Model(video).Update("favorite_count", video.FavoriteCount - 1).Error
}
// 根据视频id找到对应视频
func findFavoriteCount(videoId int64) (video *module.Video) {
	video = new(module.Video)
	DB.Where("id=?", videoId).First(video)
	return
}


func GetLikeListByUserId(userId int64) (videolist *module.VideoList, err error) {
	
	var relations []module.UserVideoRelation
	err = DB.Where("user_id=? and is_favorite=?", userId, 1).Find(&relations).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // find未找到
		return nil, errors.New("user doesn't like any videos")
	}

	var video module.Video
	var videos []module.Video
	for index := range relations {
		video.Id = relations[index].VideoId
		err = DB.Where("is_favorite=?", 0).Order("created_at desc").First(&video).Error
		if errors.Is(err, gorm.ErrRecordNotFound) { // find未找到
			return nil, errors.New("videos doesn't exist! ")
		}
		video.IsFavorite = true
		videos = append(videos, video)
	}
	
	videolist = new(module.VideoList)
	videolist.AllVideos = videos
	return
}
