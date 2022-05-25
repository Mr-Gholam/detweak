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
	db.AutoMigrate(&User{}, &Post{}, &LikedPost{}, &FriendShip{})
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
func findUserByEmail(email string) (string, string, uint, string, string) {
	var user User
	db.Where("email = ?", email).First(&user)
	if len(user.Email) > 0 {
		return user.Email, user.Password, user.ID, user.Username, user.ImgUrl
	} else {
		return "", "", 0, "", ""
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
func findUserIdByUsername(username string) uint {
	var user User
	db.Where("username = ?", username).Find(&user)
	return user.ID
}

// post
func getPostsByUserId(ownerId uint, userId uint) []PostJSON {
	var postsInfo []Post
	db.Where("owner_id = ?", ownerId).Find(&postsInfo)
	username, firstname, lastname, userImgUrl := findUserById(ownerId)
	var posts []PostJSON
	for i := 0; i < len(postsInfo); i++ {
		var post PostJSON
		liked := getPostLikedByUserId(userId, postsInfo[i].ID)
		post.AllowComments = postsInfo[i].AllowComments
		post.CreatedAt = postsInfo[i].CreatedAt
		post.Description = postsInfo[i].Description
		post.PostImgUrl = postsInfo[i].PostImgUrl
		post.Likes = postsInfo[i].Likes
		post.PostId = postsInfo[i].ID
		post.Liked = liked
		post.Username = username
		post.Firstname = firstname
		post.Lastname = lastname
		post.ProfileImg = userImgUrl
		posts = append(posts, post)
	}
	return posts
}

func getPostLikedByUserId(userId uint, postId uint) bool {
	var liked LikedPost
	db.Where("user_id = ? AND post_id = ?", userId, postId).Find(&liked)
	if liked.ID == 0 {
		return false
	} else {
		return true
	}
}
func likePost(postId uint) {
	var post Post
	db.Where("id=?", postId).First(&post)
	newLikes := post.Likes + 1
	db.Model(&Post{}).Where("id=?", postId).Update("likes", newLikes)
}
func dislikePost(postId uint) {
	var post Post
	db.Where("id=?", postId).First(&post)
	if post.Likes == 0 {
		return
	}
	newLikes := post.Likes - 1
	db.Model(&Post{}).Where("id=?", postId).Update("likes", newLikes)
}
func getPostImgUrl(postId uint) string {
	var post Post
	db.Where("id = ?", postId).Find(&post)
	return post.PostImgUrl
}

// profile
func findUserInfoByUsername(username string, userId uint) ProfileInfo {
	var user User
	db.Where("username = ?", username).Find(&user)
	posts := getPostsByUserId(user.ID, userId)
	friends := findFriendsByUserId(user.ID)
	friendshipStatus := findFriendShip(user.ID, userId)
	var profileInfo ProfileInfo
	profileInfo.Username = user.Username
	profileInfo.Firstname = user.Firstname
	profileInfo.Lastname = user.Lastname
	profileInfo.Bio = user.Bio
	profileInfo.ImgUrl = user.ImgUrl
	profileInfo.FrameWork = user.FrameWork
	profileInfo.Language = user.Language
	profileInfo.Field = user.Field
	profileInfo.Birthday = user.Birthday
	profileInfo.Location = user.Location
	profileInfo.Field = user.Field
	profileInfo.GitHubUsername = user.GitHubUsername
	profileInfo.Experience = user.Experience
	profileInfo.Posts = posts
	profileInfo.IsFriend = friendshipStatus
	profileInfo.PostCount = len(posts)
	profileInfo.FriendCount = len(friends)
	return profileInfo
}

// search
func findUserByName(name string, userId uint) []SearchResult {
	var users []User
	var result []SearchResult
	db.Where("firstname LIKE ?", name).Or("lastname LIKE ?", name).Find(&users)
	for i := 0; i < len(users); i++ {
		var userInfo SearchResult
		userInfo.Email = users[i].Email
		userInfo.Firstname = users[i].Firstname
		userInfo.Lastname = users[i].Lastname
		userInfo.ImgUrl = users[i].ImgUrl
		result = append(result, userInfo)
	}
	return result
}

// friend ship
func findFriendShip(senderId uint, receiverId uint) string {
	var friendShip FriendShip
	db.Where("sender_id =? AND receiver_id = ?", senderId, receiverId).Or("sender_id =? AND receiver_id = ?", receiverId, senderId).Find(&friendShip)
	if friendShip.ID == 0 {
		return "Not Friend"
	} else if friendShip.ID != 0 && friendShip.Status == false {
		return "Pending"
	} else {
		return "Friend"
	}
}
func getFriendRequests(receiverId uint) []FriendReq {
	var friendship []FriendShip
	db.Where("receiver_id = ? AND status = ?", receiverId, false).Find(&friendship)
	var requests []FriendReq
	for i := 0; i < len(friendship); i++ {
		var request FriendReq
		username, firstname, lastname, imgUrl := findUserById(friendship[i].SenderId)
		request.Firstname = firstname
		request.Lastname = lastname
		request.Username = username
		request.ImgUrl = imgUrl
		request.RequestId = friendship[i].ID
		requests = append(requests, request)
	}
	return requests
}
func findFriendsByUserId(userId uint) []uint {
	var friendship []FriendShip
	var friendsId []uint
	db.Where("sender_id=? AND status = ?", userId, true).Or("receiver_id =? AND status = ?", userId, true).Find(&friendship)

	for i := 0; i < len(friendship); i++ {
		if friendship[i].SenderId == userId {
			friendsId = append(friendsId, friendship[i].ReceiverId)
		} else {
			friendsId = append(friendsId, friendship[i].SenderId)
		}
	}
	return friendsId
}
