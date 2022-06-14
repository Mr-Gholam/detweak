package main

import (
	"time"
)

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Email          string `json:"email,omitempty"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	Firstname      string `json:"firstname,omitempty"`
	Lastname       string `json:"lastname,omitempty"`
	Bio            string `json:"bio,omitempty"`
	GitHubUsername string `json:"githubUsername,omitempty"`
	Field          string `json:"field,omitempty"`
	Language       string `json:"language,omitempty"`
	FrameWork      string `json:"framework,omitempty"`
	Birthday       string `json:"birthday,omitempty"`
	Location       string `json:"location,omitempty"`
	Experience     string `json:"experience,omitempty"`
	ImgUrl         string
	Token          string
	TokenExpire    time.Time
	CreatedAt      time.Time
}
type Post struct {
	ID            uint `gorm:"primaryKey"`
	OwnerId       uint
	Description   string `json:"description"`
	AllowComments bool   `json:"allowComments"`
	Likes         uint
	PostImgUrl    string
	CreatedAt     time.Time
}
type FriendShip struct {
	ID         uint `gorm:"primaryKey"`
	SenderId   uint
	ReceiverId uint
	Status     bool
}

type LikedPost struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	PostId    uint `json:"postId"`
	OwnerId   uint
	Seen      bool
	CreatedAt time.Time
}
type SearchResult struct {
	Email        string
	Username     string
	Firstname    string
	Lastname     string
	ImgUrl       string
	friendStatus string
}
type ProfileInfo struct {
	Username       string
	Firstname      string
	Lastname       string
	Bio            string
	Location       string
	Birthday       string
	Language       string
	Field          string
	FrameWork      string
	GitHubUsername string
	ImgUrl         string
	Experience     string
	Posts          []PostJSON
	PostCount      int
	FriendCount    int
	IsFriend       string
}
type Comment struct {
	ID        uint `gorm:"primaryKey"`
	OwnerId   uint
	PostId    uint
	Comment   string
	CreatedAt time.Time
}
type FriendReq struct {
	RequestId uint
	Username  string
	Firstname string
	Lastname  string
	ImgUrl    string
}
type Job struct {
	ID          uint `gorm:"primaryKey"`
	ProjectName string
	Language    string
	FrameWork   string
	Description string
	Budget      int64
	OwnerId     uint
	Deadline    string
	CreatedAt   time.Time
}
type PostJSON struct {
	IsFriend      string
	PostId        uint
	Description   string
	AllowComments bool
	Likes         uint
	PostImgUrl    string
	CreatedAt     time.Time
	Username      string
	Firstname     string
	Lastname      string
	ProfileImg    string
	Liked         bool
	Comments      []CommentJSON
}

type CommentJSON struct {
	CommentId         uint
	CommenterImgUrl   string
	CommenterUsername string
	Comment           string
	CreatedAt         time.Time
}
type Room struct {
	ID         uint
	SenderId   uint
	ReceiverId uint
}
type Chat struct {
	ID         uint
	ReceiverId uint
	ChatRoomId uint
	Seen       bool
	Message    string
	ImgUrl     string
	CreatedAt  time.Time
}
type RoomJSON struct {
	RoomId    uint
	Username  string
	Firstname string
	Lastname  string
	ImgUrl    string
	TargetId  uint
	UnseenMsg int64
	Chat      []ChatJSON
}

type ChatJSON struct {
	ImgUrl    string
	Receive   bool
	Message   string
	CreatedAt time.Time
}
