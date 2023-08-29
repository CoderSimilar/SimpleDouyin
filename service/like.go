package service

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
)

func LikeAction(newRelation *module.UserVideoRelation) (err error) {

	if exists, err := mysql.CheckLikeExist(newRelation.UserId, newRelation.VideoId); err != nil {
		// 处理错误
		return err
	} else if exists {
		// 点赞已存在
		return nil
	} else {
		// 点赞不存在
		// 2，处理点赞纪录
		if err := mysql.LikeHandle(newRelation); err != nil {
			return err
		}

		// 3，更新video的FavoriteCount
		return mysql.UpdateFavotiteCount(newRelation)
	}

	
}


func LikeList(userId int64) (videoList *module.VideoList, err error) {

	videoList, err = mysql.GetLikeListByUserId(userId)
	if err != nil {
		return
	}

	for index := range videoList.AllVideos {
		videoList.AllVideos[index].Author.UserId = userId
		mysql.GetUserInfo(&videoList.AllVideos[index].Author)
	}
	return videoList, err
	
}