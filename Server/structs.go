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
type friendShip struct {
	ID         uint `gorm:"primaryKey"`
	SenderId   uint
	ReceiverId uint
	Status     string
}
type LikedPost struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint
	PostId uint `json:"postId"`
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
	IsFriend       bool
}
type ErrorRespone struct {
	Error ErrorMessage `json:"error,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}
type UserJSON struct {
	Username string `json:"username,omitempty"`
	ImgUrl   string `json:"imgUrl,omitempty"`
}
type PostJSON struct {
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
}
type PostImg struct {
	ImgUrl string
}
type LikedPostJSON struct {
	Added bool
}
