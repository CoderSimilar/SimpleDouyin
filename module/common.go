package module

import (
	"gorm.io/gorm"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	gorm.Model    `json:"-"`
	Id            int64  `json:"id" gorm:"primarykey"`
	AuthorId      int64  `json:"-" gorm:"not null"`
	Author        User   `json:"author" gorm:"foreignKey:AuthorId"`
	PlayUrl       string `json:"play_url,omitempty" binding:"required"`
	CoverUrl      string `json:"cover_url,omitempty" binding:"required"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite" gorm:"column:is_favorite"`
	Title         string `json:"title"`
	StorePath     string `json:"-"`
}

type VideoList struct {
	AuthorId  int64 `json:"-"`
	AllVideos []Video
}

type UserVideoRelation struct {
	gorm.Model
	UserId     int64 `json:"user_id"`
	VideoId    int64 `json:"video_id" gorm:"column:video_id"`
	IsFavorite bool  `json:"is_favorite" gorm:"column:is_favorite"`
}

func (UserVideoRelation) TableName() string {
	return "user_video_relations"
}

type Comment struct {
	gorm.Model      `json:"-"`
	Id       		int64  `json:"id,omitempty" gorm:"unique"`    // 评论id，当ActionType为false时有效
	VideoId         int64  `json:"-" bind:"required"`             // 所属的视频Id
	UserId          int64  `json:"-" gorm:"column:user_id"`       // 用户id，将其设置成外键
	User            User   `json:"user" gorm:"foreignKey:UserId"` // 用户
	Content         string `json:"content,omitempty" binding:"required;oneof=1 2" gorm:"column:content"` // 评论内容，当ActionType为true时有效
	ActionType      string `json:"-"`                                                                    // 发布时为"1"，删除时为"2"
	CreatedAtString string `json:"create_date"`
}

type CommentList struct {
	VideoId     int64 `json:"-"`
	AllComments []Comment
}

type User struct {
	gorm.Model      `json:"-"`
	Id              int64  `json:"id,omitempty" gorm:"primaryKey index:idx_user_id"`                                // 用户id
	Name            string `json:"name" binding:"required" gorm:"column:username; unique; not null; index:idx_name"`// 用户名称
	Password        string `json:"-" binding:"required" gorm:"not null"`
	FollowCount     int64  `json:"follow_count" gorm:"default:0"`    // 关注总数
	FollowerCount   int64  `json:"follower_count" gorm:"default:0"`  // 粉丝总数
	IsFollow        bool   `json:"is_follow" gorm:"default:true"`    // true-已关注，false-未关注
	Avatar          string `json:"avatar" gorm:"default:null"`       // 用户头像
	BackgroundImage string `json:"background_image"`                 // 用户个人页顶部大图
	Signature       string `json:"signature"`                        // 个人简介
	TotalFavorited  int    `json:"total_favorited" gorm:"default:0"` // 获赞数量
	WorkCount       int    `json:"work_count"`                       // 作品数量
	FavoriteCount   int    `json:"favorite_count"`                   // 点赞数量
	Token           string `json:"-"`
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
