package controller

import (
	"SimpleDouyin/module"
	"time"
)

var DemoVideos = []module.Video{
	{
		Id:           		1,
		AuthorId:      		DemoUser.UserId,
		PlayUrl:      		"https://www.w3schools.com/html/movie.mp4",
		CoverUrl:     		"https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount:    	0,
		CommentCount: 		0,
	},
}

var DemoComments = []module.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: time.Now(),
	},
}

var DemoUser = module.User{
	UserId:        1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
