package controller

import "time"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id             int64     `json:"id,omitempty" gorm:"column:VideoId;primaryKey"`
	Author         User      `json:"author" gorm:"foreignKey:AuthorID"`
	PlayUrl        string    `json:"play_url" gorm:"column:PlayUrl"`
	CoverUrl       string    `json:"cover_url,omitempty" gorm:"column:CoverUrl"`
	LikeCount      int       `json:"LikeCount" gorm:"column:LikeCount"`
	CommentCount   int       `json:"comment_count" gorm:"column:CommentCount"`
	UpdateDatetime time.Time `json:"update-datetime" gorm:"column:UpdateDatetime"`
}

type UserVideoRelation struct {
	Id             int64     `json:"id,omitempty" gorm:"column:Id;primaryKey"`
	UserId         int       `json:"user_id" gorm:"column:UserId"`
	VideoId        int       `json:"video_id" gorm:"column:VideoId"`
	UpdateDatetime time.Time `json:"update_datetime" gorm:"column:UpdateDatetime"`
	IsLiked        bool      `json:"is_liked" gorm:"column:IsLiked"`
}

type Comment struct {
	Id         string `json:"id,omitempty"`      // 评论Id
	CommentId  string `json:"comment_id"`        // 要删除的评论id，当ActionType为false时有效
	VideoId    string `json:"video_id"`          // 所属的视频Id
	User       User   `json:"user"`              // 用户Id
	Content    string `json:"content,omitempty"` // 评论内容，当ActionType为true时有效
	ActionType string `json:"action_type"`       // 发布时为"1"，删除时为"2"
	CreateDate string `json:"create_date,omitempty"`
	Token      string `json:"token"` // 鉴权token
}

type Like struct {
	Id         string `json:"id"`          // 点赞Id
	VideoId    string `json:"video_id"`    // 所属的视频Id
	ActionType string `json:"action_type"` // 行为类型，点赞时为"1"，取消点赞时为“2”
	Token      string `json:"token"`       // 鉴权token
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
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
