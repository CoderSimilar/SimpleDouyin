package module

import (
	"time"

	"gorm.io/gorm"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	gorm.Model    `json:"-"`
	Id            int64  `json:"id,omitempty" gorm:"unique;not null;index:idx_video_id"`
	AuthorId      int64  `json:"author_id" gorm:"not null"`
	Author 		  User   `json:"author" gorm:"foreignKry:AuthorId"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url,omitempty" binding:"required"`
	CoverUrl      string `json:"cover_url,omitempty" binding:"required"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	StorePath     string
}

type VideoList struct {
	AuthorId   	int64 `json:"-"`
	AllVideoes 	[]Video
}

type UserVideoRelation struct {
	Id             int64     `json:"id,omitempty" gorm:"column:Id;primaryKey"`
	UserId         int       `json:"user_id" gorm:"column:UserId"`
	VideoId        int       `json:"video_id" gorm:"column:VideoId"`
	UpdateDatetime time.Time `json:"update_datetime" gorm:"column:UpdateDatetime"`
	IsLiked        bool      `json:"is_liked" gorm:"column:IsLiked"`
}

type Comment struct {
	Id         int64     `json:"id,omitempty"`      // 评论Id
	CommentId  string    `json:"comment_id"`        // 要删除的评论id，当ActionType为false时有效
	VideoId    string    `json:"video_id"`          // 所属的视频Id
	User       User      `json:"user"`              // 用户Id
	Content    string    `json:"content,omitempty"` // 评论内容，当ActionType为true时有效
	ActionType string    `json:"action_type"`       // 发布时为"1"，删除时为"2"
	CreateDate time.Time `json:"create_date,omitempty"`
	Token      string    `json:"token"` // 鉴权token
}

type Like struct {
	Id         string `json:"id"`          // 点赞Id
	VideoId    string `json:"video_id"`    // 所属的视频Id
	ActionType string `json:"action_type"` // 行为类型，点赞时为"1"，取消点赞时为“2”
	Token      string `json:"token"`       // 鉴权token
}

type User struct {
	gorm.Model      `json:"-"`
	UserId          int64  `json:"id,omitempty" gorm:"index:idx_user_id"`                           // 用户id
	Name            string `json:"name" binding:"required" gorm:"unique;not null;index:idx_name"` 	// 用户名称
	Password        string `json:"-" binding:"required" gorm:"not null"`
	FollowCount     int64  `json:"follow_count" gorm:"default:0"` 									// 关注总数
	FollowerCount   int64  `json:"follower_count" gorm:"default:0"`                  				// 粉丝总数
	IsFollow        bool   `json:"is_follow" gorm:"default:true"`                  					// true-已关注，false-未关注
	Avatar          string `json:"avatar" gorm:"default:null"`                  					// 用户头像
	BackgroundImage string `json:"background_image"`                								// 用户个人页顶部大图
	Signature       string `json:"signature"`                  										// 个人简介
	TotalFavorited  int    `json:"total_favorited" gorm:"default:0"`                 				// 获赞数量
	WorkCount       int    `json:"work_count"`                  									// 作品数量
	FavoriteCount   int64    `json:"favorite_count"`                  								// 点赞数量
	Token           string `json:"-" gorm:"-"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
