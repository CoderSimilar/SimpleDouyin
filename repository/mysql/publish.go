package mysql

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository"
	"errors"

	"gorm.io/gorm"
)

func CheckVideoExist(authorId int64, playUrl string) (err error) {
	err = DB.Where("author_id=? and play_url=?", authorId, playUrl).First(&module.Video{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	
	return repository.ErrorVideoExist
}

// CreatePublish 发布的视频信息写入数据库中
func CreatePublish(video *module.Video) (err error) {
	// 视频信息插入数据库
	return DB.Create(video).Error
}

func GetPublishListByUserId(authorId int64) (videolist *module.VideoList, err error) {
	videolist = new(module.VideoList)
	// 视频按时间顺序逆置显示
	var videos []module.Video
	err = DB.Where("author_id=?", authorId).Order("created_at desc").Find(&videos).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // find未找到
		return nil, errors.New("user doesn't have videos")
	}
	videolist.AllVideos = videos
	return
}
