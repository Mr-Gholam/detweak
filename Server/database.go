package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func connectDb() {
	var err error
	dsn := "root:mehdi007@tcp(127.0.0.1:3306)/detweak?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	db.AutoMigrate(&User{}, &Post{})
	if err != nil {
		panic(err)
	}

}

// user
func findDuplicateEmail(email string) bool {
	var user User
	db.Where("email = ?", email).First(&user)
	if len(user.Email) > 0 {
		return true
	} else {
		return false
	}
}
func findUserByEmail(email string) (string, string, uint, string) {
	var user User
	db.Where("email = ?", email).First(&user)
	if len(user.Email) > 0 {
		return user.Email, user.Password, user.ID, user.Username
	} else {
		return "", "", 0, ""
	}
}
func findDuplicateUsername(username string) bool {
	var user User
	db.Where("username = ?", username).First(&user)
	if len(user.Username) > 0 {
		return true
	} else {
		return false
	}
}
func findUsernameById(userId uint) string {
	var user User
	db.Where("id = ?", userId).First(&user)
	return user.Username
}
func findUserById(userId uint) (string, string, string, string) {
	var user User
	db.Where("id = ?", userId).First(&user)
	return user.Username, user.Firstname, user.Lastname, user.ImgUrl
}

// post
func getPostsByUserId(ownerId uint) []PostJSON {
	var postsInfo []Post
	db.Where("owner_id = ?", ownerId).Find(&postsInfo)
	username, firstname, lastname, userImgUrl := findUserById(ownerId)
	var posts []PostJSON
	for i := 0; i < len(postsInfo); i++ {
		var post PostJSON
		post.AllowComments = postsInfo[i].AllowComments
		post.CreatedAt = postsInfo[i].CreatedAt
		post.Description = postsInfo[i].Description
		post.PostImgUrl = postsInfo[i].PostImgUrl
		post.Likes = postsInfo[i].Likes
		post.PostId = postsInfo[i].ID
		post.Username = username
		post.Firstname = firstname
		post.Lastname = lastname
		post.ProfileImg = userImgUrl
		posts = append(posts, post)
	}
	return posts
}
