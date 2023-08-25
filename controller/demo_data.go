package controller

import (
	"SimpleDouyin/module"
)

var DemoVideos = []module.Video{
	{
		AuthorId:      DemoUser.UserId,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
	},
}

var DemoComments = []module.Comment{
	{
		User:    DemoUser,
		Content: "Test Comment",
	},
}

var DemoUser = module.User{
	UserId:        1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
