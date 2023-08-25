package main

import (
	"SimpleDouyin/controller"
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
	"SimpleDouyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 配置mysql连接
	fmt.Println(" I am here")
	if err := mysql.Init(); err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}
	defer mysql.Close()

	// 初始化snowflake算法
	if err := module.SnowflakeInit("2023-08-07", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

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
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.GET("/comment/list/", controller.CommentList)

	// 使用中间件 -- 查询用户的信息，并返回
	apiRouter.Use(module.AuthMiddleWare())
	{
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/publish/action/", controller.Publish)
		apiRouter.GET("/publish/list/", controller.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.FavoriteList)
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
