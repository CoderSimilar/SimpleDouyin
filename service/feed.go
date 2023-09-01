package service

import (
	"SimpleDouyin/demoData"
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
	"fmt"
)

func Feed(userId int64, token string) (*[]module.Video, error) {

	fmt.Println("userid = ", userId)
	if token == "" {
		// 用户未登录
		return mysql.FeedVideos()
	} else {
		// 用户登录，在表中查询其喜欢的视频
		videolist, err := mysql.FeedVideos()
		if err != nil {
			return &demoData.DemoVideos, err
		}
		videosWithLike, err := mysql.CheckLikeVideo(videolist, userId)
		if err != nil {
			return &demoData.DemoVideos, err
		}
		if videosWithLike == nil {
			return &demoData.DemoVideos, err
		}
		// fmt.Println("videosWithLike = ", videosWithLike)
		return videosWithLike, nil
	}
}
