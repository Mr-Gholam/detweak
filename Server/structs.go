package main

import "time"

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
	ImgUrl         string
	CreatedAt      time.Time
}
type Post struct {
	ID            uint `gorm:"primaryKey"`
	OwnerId       uint
	Description   string `json:"description"`
	AllowComments bool   `json:"allowComments"`
	Likes         int
	PostImgUrl    string
	CreatedAt     time.Time
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
	PostId uint
}
