package service

import (
	"SimpleDouyin/module"
	"SimpleDouyin/repository"
	"SimpleDouyin/repository/mysql"
	"bytes"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"
	// "time"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var (
	suffix = map[string]bool{
		".mp4":  true,
		".avi":  true,
		".wmv":  true,
		".flv":  true,
		".mpeg": true,
		".mov":  true,
	}

	coverSuffix = ".jpg"

	frame = 1

	videoPrefix = "http://127.0.0.1:8080/static/videos/"
	coverPrefix = "http://127.0.0.1:8080/static/covers/"
)

func Publish(video *module.Video, c *gin.Context) (err error) {
	// 1.提取视频
	data, err := c.FormFile("data") // c.FormFile()接收文件，c.PostForm()接收param。data是手机上的视频文件。
	if err != nil {
		return
	}

	// 2.构建视频和封面的上传路径
	fileName := filepath.Base(data.Filename)
	fileSuffix := filepath.Ext(fileName) // 判断视频后缀
	if _, ok := suffix[fileSuffix]; !ok {
		return repository.ErrorInvalidVideoFormat
	}

	finalName := fmt.Sprintf("%d_%s", video.AuthorId, strings.TrimSuffix(fileName, fileSuffix)) // 去除后缀
	videoName := finalName + fileSuffix
	coverName := finalName + coverSuffix
	saveVideoFile := filepath.Join("./public/videos/", videoName) // 视频存储路径
	saveCoverFile := filepath.Join("./public/covers/", coverName) // 封面存储路径
	fmt.Printf("finalName = %s, saveVideoFile = %s, saveCoverFile = %s\n ", finalName, saveVideoFile, saveCoverFile)

	// 3.上传视频和封面
	// 上传视频
	if err = c.SaveUploadedFile(data, saveVideoFile); err != nil {
		return
	}
	// 截取视频帧作为封面
	img, err := GetSnapshot(saveVideoFile, finalName, frame)
	if err != nil {
		return
	}
	// 上传封面
	err = imaging.Save(img, saveCoverFile)
	if err != nil {
		//log.Fatal("生成缩略图失败：", err)
		return
	}

	// 4.保存视频信息到数据库
	video.PlayUrl = videoPrefix + videoName
	video.CoverUrl = coverPrefix + coverName
	if err = SavePublishToMysql(video); err != nil {
		return
	}
	return

	// 3和4是原子操作 ？
	// 问题3：将视频转移存储到HDFS中，如何操作？
}

func PublishList(userId int64) (videoList *module.VideoList, err error) {
	fmt.Printf("author_id = %v\n", userId)
	videoList, err = mysql.GetPublishListByUserId(userId)
	if err != nil {
		return
	}
	for index := range videoList.AllVideos {
		videoList.AllVideos[index].Author.UserId = userId
		mysql.GetUserInfo(&videoList.AllVideos[index].Author)
	}
	return videoList, err
}

// GetSnapshot 从视频中截取帧作为封面
func GetSnapshot(videoPath, snapshotPath string, frameNum int) (img image.Image, err error) {
	buf := bytes.NewBuffer(nil)
	// 使用ffmpeg库进行视频处理
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		//log.Fatal("生成缩略图失败：", err)
		return nil, err
	}

	img, err = imaging.Decode(buf)
	if err != nil {
		//log.Fatal("生成缩略图失败：", err)
		return nil, err
	}
	return
}

// SavePublish 数据库的操作
func SavePublishToMysql(video *module.Video) (err error) {
	// 数据库的操作
	// 1.判断相同用户的同一作品是否发布过
	if err = mysql.CheckVideoExist(video.AuthorId, video.PlayUrl); err != nil {
		fmt.Println("Repated Publish!")
		return
	}

	// 2.生成video id
	video.ID = uint(module.GenID())
	// video.UpdateTime = time.Time{}
	// 3.写入数据库
	return mysql.CreatePublish(video)
}
