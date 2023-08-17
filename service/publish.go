package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
	"net/http"
	"path/filepath"
)

func Publish(video *module.Video, c *gin.Context) (err error) {
	// 1.上传视频
	// 获取上传的视频
	data, err := c.FormFile("data") // c.FormFile()接收文件，c.PostForm()接收param
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// mysql中存储saveFile，data是手机上的视频文件
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", video.AuthorId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return err
	}

	// 2.生成video id
	video.Id = module.GenID()

	// 3.写入数据库
	return mysql.CreatePublish(video, saveFile)

	// 问题1：视频的上传，本地是否成功？
	// 问题2：数据库中存放在本地存储的视频位置是否可行？ -- 【排序问题，是否用redis来实现】
	// 问题3：将视频转移存储到HDFS中，如何操作？

}
