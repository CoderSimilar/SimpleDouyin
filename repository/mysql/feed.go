package mysql

import (
	"errors"
	"fmt"
	"SimpleDouyin/module"
)

func FeedVideos() (videos *[]module.Video, err error) {
	// 视频按时间顺序逆置显示，最多显示30个
	videos = new([]module.Video)
	err = DB.Order("created_at desc").Limit(30).Find(videos).Error
	fmt.Println(videos)
	if err != nil { // 数据库中没有视频
		return nil, errors.New("no videos")
	}
	return
}