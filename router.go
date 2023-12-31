package main

import (
	"SimpleDouyin/controller"
	"SimpleDouyin/module"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	r.LoadHTMLGlob("templates/*")

	// home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	apiRouter := r.Group("/douyin")

	// basic apis
	// apiRouter.GET("/LarryTest/", repository.LarryTest)
	
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	
	apiRouter.GET("/comment/list/", controller.CommentList)

	
	apiRouter.Use(module.AuthFeedMiddleware()) 
	{
		apiRouter.GET("/feed/", controller.Feed)
	}

	// 使用中间件 -- 查询用户的信息，并返回
	apiRouter.Use(module.AuthMiddleWare())
	{
		
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/publish/action/", controller.Publish)
		apiRouter.GET("/publish/list/", controller.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.LikeAction)
		apiRouter.GET("/favorite/list/", controller.LikeList)
		apiRouter.POST("/comment/action/", controller.CommentAction)

		// extra apis - II
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)
		apiRouter.GET("/relation/friend/list/", controller.FriendList)
		apiRouter.GET("/message/chat/", controller.MessageChat)
		apiRouter.POST("/message/action/", controller.MessageAction)
	}

}
