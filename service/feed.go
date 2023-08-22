package service

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
)

func Feed() (*[]module.Video, error) {
	return mysql.FeedVideos()
}