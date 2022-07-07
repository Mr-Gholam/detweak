package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/smtp"
	"sort"
	"strconv"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var OnlineUserIds = map[uint]*net.Conn{}
var OnlineChatUsers = map[string]*net.Conn{}
var jwtSecret = []byte("detweak")

func (user *User) jsonUser(r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	handleError(decoder.Decode(&user))
}
func (post *Post) jsonPost(r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	handleError(decoder.Decode(&post))
}
func (likedPost *LikedPost) jsonLikedPost(r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	handleError(decoder.Decode(&likedPost))
}

func (user *User) encryptPassword() {
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	handleError(err)
	user.Password = string(hashedPassword)
}
func (user *User) validatePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err != nil
}

// file Handler
func sendFileToClient(w http.ResponseWriter, r *http.Request) {
	imgUrl := mux.Vars(r)["imgUrl"]
	fileBytes, err := ioutil.ReadFile("./images/" + imgUrl)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

// webSocket
func get_ws(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	username, firstname, lastname, imgUrl := findUserById(userId)
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	OnlineUserIds[userId] = &conn
	handleError(err)
	var friendIds []uint
	var senderIds []uint
	var receiverIds []uint
	db.Model(&FriendShip{}).Select([]string{"receiver_id"}).Where("sender_id =? AND status =?", userId, true).Find(&receiverIds)
	db.Model(&FriendShip{}).Select([]string{"sender_id"}).Where("receiver_id  =? AND status =?", userId, true).Find(&senderIds)
	friendIds = append(friendIds, senderIds...)
	friendIds = append(friendIds, receiverIds...)
	go func() {
		for i := 0; i < len(friendIds); i++ {
			onlineFriend, err := json.Marshal(map[string]interface{}{"Friend": map[string]interface{}{"Firstname": firstname, "Lastname": lastname, "Username": username, "ImgUrl": imgUrl}})
			handleError(err)
			ws, ok := OnlineUserIds[friendIds[i]]
			if ok {
				wsutil.WriteServerText(*ws, onlineFriend)
			}
		}
	}()
	go func() {
		defer conn.Close()
		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if len(msg) == 0 {
				continue
			}
			if err != nil {
				for i := 0; i < len(friendIds); i++ {
					offlineFriend, err := json.Marshal(map[string]interface{}{"offlineFriend": map[string]interface{}{"Username": username}})
					handleError(err)
					ws, ok := OnlineUserIds[friendIds[i]]
					if ok {
						wsutil.WriteServerText(*ws, offlineFriend)
					}
				}
				delete(OnlineUserIds, userId)
				break
			}
		}
	}()
}
func get_chatRoom_ws(w http.ResponseWriter, r *http.Request) {
	chatroomId := mux.Vars(r)["chatRoomId"]
	userId := getIdFromCookie(w, r)
	var targetId uint
	db.Model(&Room{}).Select([]string{"sender_id"}).Where("id =? AND receiver_id =? ", chatroomId, userId).First(&targetId)
	db.Model(&Room{}).Select([]string{"receiver_id"}).Where("id =? AND sender_id =? ", chatroomId, userId).First(&targetId)
	var userIdString string = strconv.FormatUint(uint64(userId), 10)
	var targetIdString string = strconv.FormatUint(uint64(targetId), 10)
	chatRoomUserId := "RoomId:" + chatroomId + "/" + "UserId:" + userIdString
	chatRoomTargetId := "RoomId:" + chatroomId + "/" + "UserId:" + targetIdString
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	OnlineChatUsers[chatRoomUserId] = &conn
	handleError(err)
	go func() {

	}()
	go func() {
		defer conn.Close()
		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if len(msg) == 0 {
				continue
			}
			var msgContent map[string]interface{}
			err = json.Unmarshal(msg, &msgContent)
			message, err := json.Marshal(map[string]interface{}{"Message": msgContent})
			handleError(err)
			ws, ok := OnlineChatUsers[chatRoomTargetId]
			if ok {
				wsutil.WriteServerText(*ws, message)
			}

			if err != nil {
				delete(OnlineChatUsers, chatRoomUserId)
				break
			}
		}
	}()
}

// auth controllers
func get_jwt(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*User)
	res, err := json.Marshal(user)
	handleError(err)
	w.Write(res)
}

func get_logOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: "", HttpOnly: true, Secure: true, MaxAge: -1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
}
func post_signup(w http.ResponseWriter, r *http.Request) {

	var user User
	user.jsonUser(r)

	// invalid input
	if user.Username == "" || user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Invalid Email,username,password"}})
		handleError(err)
		w.Write(e)
		return
	}
	foundEmail := findDuplicateEmail(user.Email)
	// find duplicate email
	if foundEmail {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Email already exists"}})
		handleError(err)
		w.Write(e)
		return
	}
	// find duplicate username
	if findDuplicateUsername(user.Username) {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Username already exists"}})
		handleError(err)
		w.Write(e)
		return
	}
	// hashing the password
	user.encryptPassword()
	// creating the user in database
	db.Create(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": user.Username, "id": user.ID, "imgUrl": "", "firstname": "", "lastname": ""})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusCreated)
	e, err := json.Marshal(map[string]interface{}{"username": user.Username, "ImgUrl": ""})
	handleError(err)
	w.Write(e)
}
func post_login(w http.ResponseWriter, r *http.Request) {
	var user User
	user.jsonUser(r)
	FoundEmail, foundPassword, id, username, ImgUrl := findUserByEmail(user.Email)
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "invalid Email or Password"}})
		handleError(err)
		w.Write(e)
		return
	}
	if len(FoundEmail) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Email Not found"}})
		handleError(err)
		w.Write(e)
		return
	}
	wrongPassword := user.validatePassword(foundPassword, user.Password)
	if wrongPassword {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Incorrect Password"}})
		handleError(err)
		w.Write(e)
		return
	}
	db.Model(&User{}).Select([]string{"firstname", "lastname"}).Where("id =?", id).First(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": id, "imgUrl": ImgUrl, "firstname": user.Firstname, "lastname": user.Lastname})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"username": username, "ImgUrl": ImgUrl})
	handleError(err)
	w.Write(e)
}
func post_set_profile(w http.ResponseWriter, r *http.Request) {
	var userData User
	userData.jsonUser(r)
	userId := getIdFromCookie(w, r)
	if userData.Firstname == "" || userData.Lastname == "" || userData.Bio == "" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Invalid Firstname, Lastname,Bio"}})
		handleError(err)
		w.Write(e)
		return
	}
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Invalid UserId"}})
		handleError(err)
		w.Write(e)
		return
	}
	db.Model(&userData).
		Where("id = ?", userId).Updates(User{Firstname: userData.Firstname, Lastname: userData.Lastname, Bio: userData.Bio, Birthday: userData.Birthday, Location: userData.Location})
	db.Model(&User{}).Select([]string{"username", "img_url"}).Where("id = ?", userData.ID).First(&userData)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": userData.Username, "id": userId, "imgUrl": userData.ImgUrl, "firstname": userData.Firstname, "lastname": userData.Lastname})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
}
func post_set_profile_img(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Invalid userId"}})
		handleError(err)
		w.Write(e)
		return
	}
	ImgUrl, err := FileUpload(r)
	handleError(err)
	username, firstname, lastname, _ := findUserById(userId)
	db.Model(&User{}).Where("id = ?", userId).Updates(User{ImgUrl: ImgUrl})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": userId, "imgUrl": ImgUrl, "firstname": firstname, "lastname": lastname})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"username": username, "ImgUrl": ImgUrl})
	handleError(err)
	w.Write(e)
}
func post_set_resume(w http.ResponseWriter, r *http.Request) {
	var userData User
	userData.jsonUser(r)
	userId := getIdFromCookie(w, r)
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Invalid UserId"}})
		handleError(err)
		w.Write(e)
		return
	}
	db.Model(&User{}).
		Where("id = ?", userId).
		Updates(User{GitHubUsername: userData.GitHubUsername, Language: userData.Language, FrameWork: userData.FrameWork, Field: userData.Field, Experience: userData.Experience})
	w.WriteHeader(http.StatusOK)
}
func post_reset_password(w http.ResponseWriter, r *http.Request) {
	var user map[string]interface{}
	var userId uint
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &user)
	email := user["email"].(string)
	db.Model(&User{}).Select([]string{"id"}).Where("email = ?", email).First(&userId)
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Email does not  exists"}})
		handleError(err)
		w.Write(e)
		return
	}
	token := uuid.New().String()
	location := "localhost/new-password?token=" + token
	from := "detweak@gmail.com"
	password := "Mehdi007!"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(location)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err1 := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

// profile controller
func get_profile(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	username := mux.Vars(r)["username"]
	profileInfo := findUserInfoByUsername(username, userId)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(profileInfo)
	handleError(err)
	w.Write(e)
}
func get_get_notification(w http.ResponseWriter, r *http.Request) {
	var friendship int64
	var unseenMsg int64
	var NewLikes int64
	userId := getIdFromCookie(w, r)
	db.Model(&FriendShip{}).Where("receiver_id =? AND status =?", userId, false).Count(&friendship)
	db.Model(&Chat{}).Where("receiver_id = ? AND seen =?", userId, false).Count(&unseenMsg)
	db.Model(&LikedPost{}).Where("owner_id = ? AND seen =?", userId, false).Count(&NewLikes)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"Friendship": friendship, "Messages": unseenMsg, "NewLiked": NewLikes})
	handleError(err)
	w.Write(e)
}

// post controllers
func post_create_post(w http.ResponseWriter, r *http.Request) {
	var post Post
	post.jsonPost(r)
	userId := getIdFromCookie(w, r)
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Invalid UserId"}})
		handleError(err)
		w.Write(e)
		return
	}
	if post.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Invalid Description"}})
		handleError(err)
		w.Write(e)
		return
	}
	post.OwnerId = userId
	username, firstname, lastname, profileImg := findUserById(userId)
	db.Create(&post)
	w.WriteHeader(http.StatusCreated)
	e, err := json.Marshal(
		PostJSON{
			PostId:        post.ID,
			Description:   post.Description,
			AllowComments: post.AllowComments,
			Likes:         post.Likes,
			CreatedAt:     post.CreatedAt,
			ProfileImg:    profileImg,
			Username:      username,
			Firstname:     firstname,
			Lastname:      lastname,
		})
	handleError(err)
	w.Write(e)
}
func post_create_post_img(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	imgUrl, err := FileUpload(r)
	handleError(err)
	db.Model(&Post{}).Where("id = ?", postId).Updates(Post{PostImgUrl: imgUrl})
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"ImgUrl": imgUrl})
	handleError(err)
	w.Write(e)
}
func get_available_posts(w http.ResponseWriter, r *http.Request) {
	var friendIds []uint
	var availablePosts []PostJSON
	userId := getIdFromCookie(w, r)
	friends := findFriendsByUserId(userId)
	friendIds = append(friendIds, friends...)
	friendIds = append(friendIds, userId)
	for i := 0; i < len(friendIds); i++ {
		posts := getPostsByUserId(friendIds[i], userId)
		availablePosts = append(availablePosts, posts...)
	}

	sort.Slice(availablePosts, func(i, j int) bool {
		return availablePosts[j].PostId < availablePosts[i].PostId
	})
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(availablePosts)
	handleError(err)
	w.Write(e)
}
func post_like_post(w http.ResponseWriter, r *http.Request) {
	var likedPost LikedPost
	var ownerId uint
	likedPost.jsonLikedPost(r)
	userId := getIdFromCookie(w, r)
	likedPost.UserId = userId
	db.Model(&Post{}).Select([]string{"owner_id"}).Where("id =?", likedPost.PostId).Find(&ownerId)
	likedPost.OwnerId = ownerId
	result := getPostLikedByUserId(userId, likedPost.PostId)
	if result {
		dislikePost(likedPost.PostId)
		db.Where("user_id = ? AND post_id = ?", userId, likedPost.PostId).Delete(&LikedPost{})
		if ownerId != userId {
			ws, ok := OnlineUserIds[ownerId]
			if ok {
				username := findUsernameById(userId)
				notification, err := json.Marshal(map[string]interface{}{"notification": map[string]interface{}{"Disliked": map[string]interface{}{"Username": username, "PostId": likedPost.PostId}}})
				handleError(err)
				wsutil.WriteServerText(*ws, notification)
			}
		}
		w.WriteHeader(http.StatusOK)
		e, err := json.Marshal(map[string]interface{}{"Added": false})
		handleError(err)
		w.Write(e)
	} else {
		likePost(likedPost.PostId)
		db.Create(&likedPost)
		if ownerId != userId {
			ws, ok := OnlineUserIds[ownerId]
			if ok {
				var postImg string
				db.Model(&Post{}).Select([]string{"post_img_url"}).Where("id = ?", likedPost.PostId).Find(&postImg)
				username, _, _, imgUrl := findUserById(userId)
				notification, err := json.Marshal(map[string]interface{}{"notification": map[string]interface{}{"Liked": map[string]interface{}{"Username": username, "ImgUrl": imgUrl, "PostImgUrl": postImg, "CreatedAt": likedPost.CreatedAt, "PostId": likedPost.PostId}}})
				handleError(err)
				wsutil.WriteServerText(*ws, notification)
			}
		}
		w.WriteHeader(http.StatusOK)
		e, err := json.Marshal(map[string]interface{}{"Added": true})
		handleError(err)
		w.Write(e)
	}
}
func post_update_post(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var post map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &post)
	postId := uint(post["postId"].(float64))
	NewDescription := post["newDescription"]
	db.Model(&Post{}).Where("id =? AND owner_id =?", postId, userId).Update("description", NewDescription)
	w.WriteHeader(http.StatusOK)
}
func post_delete_post(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var post map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &post)
	postId := uint(post["postId"].(float64))
	imgUrl := getPostImgUrl(postId)
	db.Where("post_id = ?", postId).Delete(&LikedPost{})
	db.Where("id = ? AND owner_id = ?", postId, userId).Delete(&Post{})
	DeleteFile(imgUrl)
	w.WriteHeader(http.StatusOK)
}
func get_single_post(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["postId"]
	postId, err := strconv.ParseUint(postID, 10, 32)
	handleError(err)
	userId := getIdFromCookie(w, r)
	post := getPostByPostId(uint(postId), userId)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(post)
	handleError(err)
	w.Write(e)
}
func get_liked_posts(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	posts := getLikedPostsByUserId(userId)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(posts)
	handleError(err)
	w.Write(e)
}
func get_liked_by(w http.ResponseWriter, r *http.Request) {
	var result []map[string]string
	var users []uint
	postID := mux.Vars(r)["postId"]
	userId := getIdFromCookie(w, r)
	postId, err := strconv.ParseUint(postID, 10, 32)
	handleError(err)
	db.Model(&LikedPost{}).Select([]string{"user_id"}).Where("post_id =?", postId).Find(&users)
	for i := 0; i < len(users); i++ {
		var userInfo User
		user := make(map[string]string)
		db.Model(&User{}).Select([]string{"firstname", "lastname", "img_url", "username", "id"}).Where("id =?", users[i]).First(&userInfo)
		user["Firstname"] = userInfo.Firstname
		user["Lastname"] = userInfo.Lastname
		user["ImgUrl"] = userInfo.ImgUrl
		user["Username"] = userInfo.Username
		user["IsFriend"] = findFriendShip(users[i], userId)
		result = append(result, user)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(result)
	handleError(err)
	w.Write(e)
}
func get_new_likes(w http.ResponseWriter, r *http.Request) {
	var posts []LikedPost
	var result []map[string]interface{}
	var newLikes int64
	userId := getIdFromCookie(w, r)
	db.Model(&LikedPost{}).Select([]string{"post_id", "user_id", "created_at"}).Where("owner_id=?", userId).Find(&posts)
	db.Model(&LikedPost{}).Where("owner_id = ? AND seen =?", userId, false).Count(&newLikes)
	db.Model(&LikedPost{}).Where("owner_id = ? AND seen =?", userId, false).Update("seen", true)
	for i := 0; i < len(posts); i++ {
		if posts[i].UserId == userId {
			continue
		}
		user := make(map[string]interface{})
		var postImg string
		username, _, _, imgUrl := findUserById(posts[i].UserId)
		db.Model(&Post{}).Select([]string{"post_img_url"}).Where("id =?", posts[i].PostId).Find(&postImg)
		user["Username"] = username
		user["ImgUrl"] = imgUrl
		user["PostImgUrl"] = postImg
		user["CreatedAt"] = posts[i].CreatedAt
		user["PostId"] = posts[i].PostId
		result = append(result, user)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"Likes": result, "NewLikes": newLikes})
	handleError(err)
	w.Write(e)
}

// Comment controller
func post_create_comment(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var commentInfo map[string]interface{}
	var commentAdded CommentJSON
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &commentInfo)
	postId := uint(commentInfo["postId"].(float64))
	commentContent := commentInfo["comment"].(string)
	username, _, _, imgUrl := findUserById(userId)
	var comment Comment
	comment.OwnerId = userId
	comment.PostId = postId
	comment.Comment = commentContent
	db.Create(&comment)
	commentAdded.Comment = commentContent
	commentAdded.CommentId = comment.ID
	commentAdded.CommenterImgUrl = imgUrl
	commentAdded.CommenterUsername = username
	commentAdded.CreatedAt = comment.CreatedAt
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(commentAdded)
	handleError(err)
	w.Write(e)
}
func post_delete_comment(w http.ResponseWriter, r *http.Request) {
	var commentInfo map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &commentInfo)
	commentId := uint(commentInfo["commentId"].(float64))
	db.Where("id = ?", commentId).Delete(&Comment{})
	w.WriteHeader(http.StatusOK)
}

// Chat controller
func post_create_room(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var receiverInfo map[string]interface{}
	var room Room
	var existingRoom Room
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &receiverInfo)
	targetUsername := receiverInfo["targetUsername"].(string)
	receiverId := findUserIdByUsername(targetUsername)
	db.Where("sender_id = ? AND receiver_id = ?", userId, receiverId).Or("sender_id = ? AND receiver_id = ?", receiverId, userId).Find(&existingRoom)
	if existingRoom.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": "Room already exists"})
		handleError(err)
		w.Write(e)
		return
	}
	room.ReceiverId = receiverId
	room.SenderId = userId
	db.Create(&room)
	targetUsername, targetFirstname, targetLastname, targetImgUrl := findUserById(receiverId)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"ImgUrl": targetImgUrl, "Firstname": targetFirstname, "Lastname": targetLastname, "Username": targetUsername, "RoomId": room.ID})
	handleError(err)
	w.Write(e)
}
func get_load_chatRooms(w http.ResponseWriter, r *http.Request) {
	var chatRooms []Room
	var Ids []map[string]uint
	var Rooms []RoomJSON
	userId := getIdFromCookie(w, r)
	db.Where("sender_id = ?", userId).Or("receiver_id = ?", userId).Find(&chatRooms)
	for i := 0; i < len(chatRooms); i++ {
		if chatRooms[i].SenderId == userId {
			target := map[string]uint{"chatRoomId": chatRooms[i].ID, "targetId": chatRooms[i].ReceiverId}
			Ids = append(Ids, target)
		} else {
			target := map[string]uint{"chatRoomId": chatRooms[i].ID, "targetId": chatRooms[i].SenderId}
			Ids = append(Ids, target)
		}
	}
	for i := 0; i < len(Ids); i++ {
		var roomDetail RoomJSON
		var chats int64
		db.Model(&Chat{}).Where("seen =? AND chat_room_id = ? AND receiver_id = ?", false, Ids[i]["chatRoomId"], userId).Count(&chats)
		username, firstname, lastname, imgUrl := findUserById(Ids[i]["targetId"])
		roomDetail.Firstname = firstname
		roomDetail.ImgUrl = imgUrl
		roomDetail.Username = username
		roomDetail.Lastname = lastname
		roomDetail.UnseenMsg = chats
		roomDetail.RoomId = Ids[i]["chatRoomId"]
		Rooms = append(Rooms, roomDetail)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(Rooms)
	handleError(err)
	w.Write(e)
}
func post_get_chat(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var post map[string]interface{}
	var room Room
	var chats []Chat
	var roomJSON RoomJSON
	var chatJson []ChatJSON
	var targetId uint
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &post)
	roomId := uint(post["RoomId"].(float64))
	db.Where("chat_room_id = ?", roomId).Find(&chats)
	db.Where("sender_id =? AND id = ?", userId, roomId).Or("receiver_id =? AND id =?", userId, roomId).Find(&room)

	for i := 0; i < len(chats); i++ {
		var chat ChatJSON
		if chats[i].ReceiverId == userId {
			chat.Receive = true
			chat.Message = chats[i].Message
			chat.ImgUrl = chats[i].ImgUrl
			chat.CreatedAt = chats[i].CreatedAt
			chatJson = append(chatJson, chat)
		} else {
			chat.Receive = false
			chat.ImgUrl = chats[i].ImgUrl
			chat.Message = chats[i].Message
			chat.CreatedAt = chats[i].CreatedAt
			chatJson = append(chatJson, chat)
		}
	}
	if room.ReceiverId == userId {
		targetId = room.SenderId
	} else {
		targetId = room.ReceiverId
	}
	username, firstname, lastname, imgUrl := findUserById(targetId)
	roomJSON.Chat = chatJson
	roomJSON.Firstname = firstname
	roomJSON.Lastname = lastname
	roomJSON.Username = username
	roomJSON.ImgUrl = imgUrl
	roomJSON.RoomId = roomId
	roomJSON.TargetId = targetId
	db.Model(&Chat{}).Where("chat_room_id = ? AND receiver_id =? AND seen =?", roomId, userId, false).Update("seen", true)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(roomJSON)
	handleError(err)
	w.Write(e)
}
func post_create_message(w http.ResponseWriter, r *http.Request) {
	var messageInfo map[string]interface{}
	var chat Chat
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &messageInfo)
	roomId := uint(messageInfo["RoomId"].(float64))
	targetId := uint(messageInfo["TargetId"].(float64))
	ws, ok := OnlineUserIds[targetId]
	if ok {
		unseenMsg, err := json.Marshal(map[string]interface{}{"UnSeenMsg": map[string]interface{}{"RoomId": roomId}})
		handleError(err)
		wsutil.WriteServerText(*ws, unseenMsg)
	}
	message := messageInfo["Message"].(string)
	chat.ChatRoomId = roomId
	chat.Message = message
	chat.ReceiverId = targetId
	db.Create(&chat)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"CreatedAt": chat.CreatedAt})
	handleError(err)
	w.Write(e)
}
func post_create_message_img(w http.ResponseWriter, r *http.Request) {
	var room Room
	var chat Chat
	targetId := mux.Vars(r)["TargetId"]
	receiverId, err := strconv.ParseUint(targetId, 0, 64)
	handleError(err)
	userId := getIdFromCookie(w, r)
	db.Where("sender_id =? AND receiver_id = ?", userId, targetId).Or("sender_id =? AND receiver_id = ?", targetId, userId).Find(&room)
	imgUrl, err := FileUpload(r)
	handleError(err)
	chat.ImgUrl = imgUrl
	chat.ReceiverId = uint(receiverId)
	chat.ChatRoomId = room.ID
	db.Create(&chat)
	ws, ok := OnlineUserIds[uint(receiverId)]
	if ok {
		unseenMsg, err := json.Marshal(map[string]interface{}{"UnSeenMsg": map[string]interface{}{"RoomId": room.ID}})
		handleError(err)
		wsutil.WriteServerText(*ws, unseenMsg)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"CreatedAt": chat.CreatedAt, "ImgUrl": imgUrl, "MessageId": chat.ID})
	handleError(err)
	w.Write(e)
}
func post_update_message(w http.ResponseWriter, r *http.Request) {
	var messageInfo map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &messageInfo)
	messageId := uint(messageInfo["messageId"].(float64))
	message := messageInfo["message"].(string)
	db.Model(&Chat{}).Where("id = ?", messageId).Update("message", message)
	w.WriteHeader(http.StatusOK)
}

// Setting controller
func get_setting(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var user User
	db.Where("id =? ", userId).Find(&user)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{
		"Email":      user.Email,
		"Username":   user.Username,
		"Firstname":  user.Firstname,
		"Lastname":   user.Lastname,
		"Bio":        user.Bio,
		"GitHub":     user.GitHubUsername,
		"Field":      user.Field,
		"Language":   user.Language,
		"FrameWork":  user.FrameWork,
		"Birthday":   user.Birthday,
		"Location":   user.Location,
		"Experience": user.Experience,
		"ImgUrl":     user.ImgUrl,
	})
	handleError(err)
	w.Write(e)
}
func post_remove_profileImg(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &data)
	ImgUrl := data["ImgUrl"].(string)
	userId := getIdFromCookie(w, r)
	db.Model(&User{}).Where("id =?", userId).Update("ImgUrl", "")
	DeleteFile(ImgUrl)
	username, firstname, lastname, _ := findUserById(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": userId, "imgUrl": " ", "firstname": firstname, "lastname": lastname})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
}
func post_update_profileImg(w http.ResponseWriter, r *http.Request) {
	var user User
	userId := getIdFromCookie(w, r)
	db.Where("id = ?", userId).Find(&user)
	if user.ImgUrl != "" {
		DeleteFile(user.ImgUrl)
	}
	NewImgUrl, err := FileUpload(r)
	handleError(err)
	db.Model(&User{}).Where("id = ?", userId).Update("ImgUrl", NewImgUrl)
	username, firstname, lastname, _ := findUserById(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": userId, "imgUrl": NewImgUrl, "firstname": firstname, "lastname": lastname})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"ImgUrl": NewImgUrl})
	handleError(err)
	w.Write(e)
}
func post_update_personal(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	updated := make(map[string]interface{})
	userId := getIdFromCookie(w, r)
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &data)
	firstname, firstOk := data["Firstname"]
	if firstOk {
		updated["Firstname"] = firstname
	}
	lastname, lastOk := data["Lastname"]
	if lastOk {
		updated["Lastname"] = lastname
	}
	bio, bioOk := data["Bio"]
	if bioOk {
		updated["Bio"] = bio
	}
	birthday, birthOk := data["Birthday"]
	if birthOk {
		updated["Birthday"] = birthday
	}
	location, locationOk := data["Location"]
	if locationOk {
		updated["Location"] = location
	}
	db.Model(&User{}).Where("id = ?", userId).Updates(updated)
	w.WriteHeader(http.StatusOK)
}
func post_update_account(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	updated := make(map[string]interface{})
	userId := getIdFromCookie(w, r)
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &data)
	email, emailOk := data["Email"]
	if emailOk {
		duplicateEmail := findDuplicateEmail(email.(string))
		if duplicateEmail {
			w.WriteHeader(http.StatusBadRequest)
			e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Email already exists"}})
			handleError(err)
			w.Write(e)
			return
		}
		updated["Email"] = email
	}
	username, usernameOk := data["Username"]
	if usernameOk {
		duplicateUsername := findDuplicateUsername(username.(string))
		if duplicateUsername {
			w.WriteHeader(http.StatusBadRequest)
			e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Username already exists"}})
			handleError(err)
			w.Write(e)
			return
		}
		if len(username.(string)) <= 3 {
			w.WriteHeader(http.StatusBadRequest)
			e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Username must be 3 character "}})
			handleError(err)
			w.Write(e)
			return
		}
		updated["Username"] = username
	}
	db.Model(&User{}).Where("id = ?", userId).Updates(updated)
	newUsername, firstname, lastname, imgUrl := findUserById(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": newUsername, "id": userId, "imgUrl": imgUrl, "firstname": firstname, "lastname": lastname})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"message": "user updated"})
	handleError(err)
	w.Write(e)
}
func post_update_password(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var user User
	userId := getIdFromCookie(w, r)
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &data)
	db.Where("id = ?", userId).Find(&user)
	oldPassword := data["currentPassword"]
	wrongPassword := user.validatePassword(user.Password, oldPassword.(string))
	if wrongPassword {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Incorrect Password"}})
		handleError(err)
		w.Write(e)
		return
	}
	newPassword := data["newPassword"]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword.(string)), bcrypt.DefaultCost)
	handleError(err)
	db.Model(&User{}).Where("id = ?", userId).Update("password", hashedPassword)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"message": "password updated"})
	handleError(err)
	w.Write(e)
}
func post_update_professional(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	updated := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &data)
	userId := getIdFromCookie(w, r)
	github, githubOk := data["GitHub"]
	if githubOk {
		updated["GitHubUsername"] = github
	}
	xp, xpOk := data["Experience"]
	if xpOk {
		updated["Experience"] = xp
	}
	language, languageOk := data["Language"]
	if languageOk {
		updated["Language"] = language
	}
	field, fieldOk := data["Field"]
	if fieldOk {
		updated["Field"] = field
	}
	framework, frameworkOk := data["FrameWork"]
	if frameworkOk {
		updated["FrameWork"] = framework
	}
	db.Model(&User{}).Where("id =?", userId).Updates(updated)
	w.WriteHeader(http.StatusOK)
}
func post_delete_account(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	db.Where("receiver_id = ?", userId).Delete(&Chat{})
	db.Where("owner_id = ?", userId).Delete(&Comment{})
	db.Where("user_id = ?", userId).Delete(&LikedPost{})
	db.Where("owner_id = ?", userId).Delete(&Post{})
	db.Where("sender_id = ?", userId).Or("receiver_id = ?", userId).Delete(&FriendShip{})
	db.Where("sender_id = ?", userId).Or("receiver_id = ?", userId).Delete(&Room{})
	db.Where("id = ?", userId).Delete(&User{})
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: "", HttpOnly: true, Secure: true, MaxAge: -1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
}

// Friendship Controller
func post_add_friend(w http.ResponseWriter, r *http.Request) {
	var user User
	user.jsonUser(r)
	userId := getIdFromCookie(w, r)
	targetId := findUserIdByUsername(user.Username)
	status := findFriendShip(userId, targetId)
	if status == "Friend" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Friends"}})
		handleError(err)
		w.Write(e)
		return
	}
	if status == "Pending" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(map[string]interface{}{"error": map[string]interface{}{"message": "Pending"}})
		handleError(err)
		w.Write(e)
		return
	}
	var friendShip FriendShip
	friendShip.SenderId = userId
	friendShip.ReceiverId = targetId
	friendShip.Status = false
	db.Create(&friendShip)
	ws, ok := OnlineUserIds[targetId]
	if ok {
		username, firstname, lastname, imgUrl := findUserById(userId)
		notification, err := json.Marshal(map[string]interface{}{"notification": map[string]interface{}{"NewFriendReq": map[string]interface{}{"Username": username, "Firstname": firstname, "Lastname": lastname, "ImgUrl": imgUrl, "RequestId": friendShip.ID}}})
		handleError(err)
		wsutil.WriteServerText(*ws, notification)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"status": "Pending"})
	handleError(err)
	w.Write(e)
}
func get_friend_requests(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	requests := getFriendRequests(userId)
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(requests)
	handleError(err)
	w.Write(e)

}
func get_online_friends(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var result []map[string]interface{}
	var friendIds []uint
	var senderIds []uint
	var receiverIds []uint
	db.Model(&FriendShip{}).Select([]string{"receiver_id"}).Where("sender_id =? AND status =?", userId, true).Find(&receiverIds)
	db.Model(&FriendShip{}).Select([]string{"sender_id"}).Where("receiver_id  =? AND status =?", userId, true).Find(&senderIds)
	friendIds = append(friendIds, senderIds...)
	friendIds = append(friendIds, receiverIds...)
	for i := 0; i < len(friendIds); i++ {
		_, ok := OnlineUserIds[friendIds[i]]
		if ok {
			var user = make(map[string]interface{})
			username, firstname, lastname, imgUrl := findUserById(friendIds[i])
			user["Username"] = username
			user["Firstname"] = firstname
			user["Lastname"] = lastname
			user["ImgUrl"] = imgUrl
			result = append(result, user)
		}
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(result)
	handleError(err)
	w.Write(e)
}
func post_accept_request(w http.ResponseWriter, r *http.Request) {
	var post map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &post)
	requestId := uint(post["requestId"].(float64))
	db.Model(&FriendShip{}).Where("id = ?", requestId).Update("status", true)
	w.WriteHeader(http.StatusOK)
}
func post_reject_request(w http.ResponseWriter, r *http.Request) {
	var post map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &post)
	requestId := uint(post["requestId"].(float64))
	db.Where("id = ?", requestId).Delete(&FriendShip{})
	w.WriteHeader(http.StatusOK)
}
func get_friends_by_username(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	var result []map[string]string
	var friendsId []uint
	var receiverIds []uint
	var senderIds []uint
	userId := getIdFromCookie(w, r)
	targetId := findUserIdByUsername(username)
	db.Model(&FriendShip{}).Select("receiver_id").Where("sender_id =? AND status = ?", targetId, true).Find(&receiverIds)
	db.Model(&FriendShip{}).Select("sender_id").Where("receiver_id =? AND status = ?", targetId, true).Find(&senderIds)
	friendsId = append(friendsId, receiverIds...)
	friendsId = append(friendsId, senderIds...)
	for i := 0; i < len(friendsId); i++ {
		var userInfo User
		user := make(map[string]string)
		db.Model(&User{}).Select([]string{"firstname", "lastname", "img_url", "username", "id"}).Where("id =?", friendsId[i]).First(&userInfo)
		user["Firstname"] = userInfo.Firstname
		user["Lastname"] = userInfo.Lastname
		user["ImgUrl"] = userInfo.ImgUrl
		user["Username"] = userInfo.Username
		user["IsFriend"] = findFriendShip(friendsId[i], userId)
		result = append(result, user)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(result)
	handleError(err)
	w.Write(e)
}

// Job
func post_create_job(w http.ResponseWriter, r *http.Request) {
	var job Job
	var user map[string]interface{}
	userId := getIdFromCookie(w, r)
	body, err := ioutil.ReadAll(r.Body)
	handleError(err)
	err = json.Unmarshal(body, &user)
	job.ProjectName = user["ProjectName"].(string)
	job.Language = user["language"].(string)
	job.Title = user["title"].(string)
	job.Field = user["field"].(string)
	frameWork, ok := user["frameWork"]
	if ok {
		job.FrameWork = frameWork.(string)
	}
	job.Deadline = user["deadline"].(string)
	job.Budget = int64(user["budget"].(float64))
	job.Description = user["description"].(string)
	job.OwnerId = userId
	db.Create(&job)
}
func get_my_jobs(w http.ResponseWriter, r *http.Request) {
	var jobs []Job
	var result []map[string]interface{}
	userId := getIdFromCookie(w, r)
	db.Model(&Job{}).Select([]string{"id", "title", "project_name", "budget", "created_at"}).Where("owner_id =?", userId).Find(&jobs)
	username, _, _, imgUrl := findUserById(userId)
	for i := 0; i < len(jobs); i++ {
		var unseen int64
		job := make(map[string]interface{})
		db.Model(&Offer{}).Where("job_id = ? AND seen =? ", jobs[i].ID, false).Count(&unseen)
		job["Id"] = jobs[i].ID
		job["Unseen"] = unseen
		job["Title"] = jobs[i].Title
		job["ProjectName"] = jobs[i].ProjectName
		job["Budget"] = jobs[i].Budget
		job["CreatedAt"] = jobs[i].CreatedAt
		job["OwnerImg"] = imgUrl
		job["OwnerUsername"] = username
		result = append(result, job)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(result)
	handleError(err)
	w.Write(e)
}
func get_find_jobs(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var jobs []Job
	var jsonJobs []JobJSON
	var language string
	db.Model(&User{}).Select([]string{"language"}).Where("id = ?", userId).Find(&language)
	db.Model(&Job{}).Select([]string{"id", "title", "project_name", "budget", "owner_id", "created_at"}).Where("language = ?", language).Find(&jobs)
	for i := 0; i < len(jobs); i++ {
		var job JobJSON
		username, _, _, imgUrl := findUserById(jobs[i].OwnerId)
		job.Id = jobs[i].ID
		job.Title = jobs[i].Title
		job.ProjectName = jobs[i].ProjectName
		job.Budget = jobs[i].Budget
		job.CreatedAt = jobs[i].CreatedAt
		job.OwnerImg = imgUrl
		job.OwnerUsername = username
		jsonJobs = append(jsonJobs, job)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(jsonJobs)
	handleError(err)
	w.Write(e)
}
func get_job_byId(w http.ResponseWriter, r *http.Request) {
	var offers []Offer
	var job Job
	var hasAcceptedOffer Offer
	result := make(map[string]interface{})
	jobId := mux.Vars(r)["jobId"]
	JobId, err := strconv.ParseUint(jobId, 10, 32)
	handleError(err)
	userId := getIdFromCookie(w, r)
	db.Where("id = ?", JobId).First(&job)
	username, firstname, lastname, imgUrl := findUserById(job.OwnerId)
	db.Where("job_id = ? AND accepted = ?", JobId, true).First(&hasAcceptedOffer)
	if job.OwnerId == userId {
		result["isOwner"] = true
		if hasAcceptedOffer.ID == 0 {
			db.Where("job_id =? AND rejected = ?", JobId, false).Find(&offers)
			var detailedOffer []map[string]interface{}
			for i := 0; i < len(offers); i++ {
				offer := make(map[string]interface{})
				username, firstname, lastname, imgUrl := findUserById(offers[i].OwnerId)
				offer["Description"] = offers[i].Description
				offer["Id"] = offers[i].ID
				offer["OwnerUsername"] = username
				offer["OwnerFirstname"] = firstname
				offer["OwnerLastname"] = lastname
				offer["OwnerImgUrl"] = imgUrl
				offer["CreatedAt"] = offers[i].CreatedAt
				detailedOffer = append(detailedOffer, offer)
			}
			orderedOffer := reverseArray(detailedOffer)
			result["Offers"] = orderedOffer
			db.Model(&Offer{}).Where("job_id =? AND seen = ?", JobId, false).Update("seen", true)
		} else {
			offer := make(map[string]interface{})
			username, firstname, lastname, imgUrl := findUserById(hasAcceptedOffer.OwnerId)
			offer["Description"] = hasAcceptedOffer.Description
			offer["Id"] = hasAcceptedOffer.ID
			offer["OwnerUsername"] = username
			offer["OwnerFirstname"] = firstname
			offer["OwnerLastname"] = lastname
			offer["OwnerImgUrl"] = imgUrl
			offer["CreatedAt"] = hasAcceptedOffer.CreatedAt
			result["AcceptedOffer"] = offer
		}
	} else {
		result["isOwner"] = false
		db.Model(&Offer{}).Select([]string{"owner_id"}).Where("job_id =?", JobId).Find(&offers)
		if hasAcceptedOffer.ID != 0 {
			offer := make(map[string]interface{})
			username, firstname, lastname, _ := findUserById(hasAcceptedOffer.OwnerId)
			offer["OwnerUsername"] = username
			offer["OwnerFirstname"] = firstname
			offer["OwnerLastname"] = lastname
			result["AcceptedOffer"] = offer
		}
		var detailedOffer []map[string]interface{}
		for i := 0; i < len(offers); i++ {
			offer := make(map[string]interface{})
			username, _, _, imgUrl := findUserById(offers[i].OwnerId)
			offer["OwnerUsername"] = username
			offer["OwnerImgUrl"] = imgUrl
			detailedOffer = append(detailedOffer, offer)
		}
		result["Others"] = detailedOffer
	}
	result["HasAcceptedOffer"] = hasAcceptedOffer.Accepted
	result["OwnerFirstname"] = firstname
	result["OwnerLastname"] = lastname
	result["OwnerUsername"] = username
	result["OwnerImgUrl"] = imgUrl
	result["Title"] = job.Title
	result["Budget"] = job.Budget
	result["CreatedAt"] = job.CreatedAt
	result["Deadline"] = job.Deadline
	result["Description"] = job.Description
	result["Field"] = job.Field
	result["FrameWork"] = job.FrameWork
	result["ID"] = job.ID
	result["Language"] = job.Language
	result["ProjectName"] = job.ProjectName
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(result)
	handleError(err)
	w.Write(e)
}
func post_create_application(w http.ResponseWriter, r *http.Request) {
	var application Offer
	var offer map[string]interface{}
	var hasOffer uint
	userId := getIdFromCookie(w, r)
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &offer)
	handleError(err)
	JobId, err := strconv.ParseUint(offer["jobId"].(string), 10, 32)
	handleError(err)
	db.Model(&Offer{}).Select([]string{"id"}).Where("owner_id = ? AND job_id =?", userId, JobId).First(&hasOffer)
	if hasOffer != 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	application.Description = offer["description"].(string)
	application.OwnerId = userId
	application.JobId = uint(JobId)
	db.Create(&application)
	w.WriteHeader(http.StatusOK)
}
func post_reject_Offer(w http.ResponseWriter, r *http.Request) {
	var offer map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &offer)
	handleError(err)
	db.Model(&Offer{}).Where("id =?", offer["offerId"]).Update("rejected", true)
	w.WriteHeader(http.StatusOK)
}
func post_accept_Offer(w http.ResponseWriter, r *http.Request) {
	var request Offer
	var offer map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &offer)
	handleError(err)
	db.Model(&Offer{}).Select([]string{"job_id", "owner_id"}).Where("id = ?", offer["offerId"]).Find(&request)
	db.Model(&Offer{}).Where("id =?", offer["offerId"]).Update("accepted", true)
	db.Model(&Job{}).Where("id = ?", request.JobId).Updates(Job{WorkerId: request.OwnerId, OfferId: request.ID})
	w.WriteHeader(http.StatusOK)
}
func get_my_Offers(w http.ResponseWriter, r *http.Request) {
	var result []map[string]interface{}
	var Offers []Offer
	var job Job
	userId := getIdFromCookie(w, r)
	db.Where("owner_id = ?", userId).Find(&Offers)
	for i := 0; i < len(Offers); i++ {
		offer := make(map[string]interface{})
		offer["JobId"] = Offers[i].JobId
		if Offers[i].Accepted {
			offer["Status"] = "Accepted"
			if !Offers[i].AcceptedSeen {
				offer["AcceptedSeen"] = true
				db.Where("id = ? AND accepted_seen = ?", Offers[i].ID, false).Update("accepted_seen", true)
			}
		} else if Offers[i].Rejected {
			offer["Status"] = "Rejected"
			if !Offers[i].RejectedSeen {
				offer["RejectedSeen"] = true
				db.Where("id = ? AND rejected_seen = ?", Offers[i].ID, false).Update("rejected_seen", true)
			}
		} else {
			offer["Status"] = "Pending"
		}
		db.Model(&Job{}).Select([]string{"title", "project_name"}).Where("id=?", Offers[i].JobId).First(&job)
		offer["ProjectName"] = job.ProjectName
		offer["Title"] = job.Title
		result = append(result, offer)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(result)
	handleError(err)
	w.Write(e)
}

// search
func get_search(w http.ResponseWriter, r *http.Request) {
	userId := getIdFromCookie(w, r)
	var users []User
	var posts []Post
	var postResult []PostJSON
	var userResults []map[string]string
	userInput := mux.Vars(r)["userInput"]
	db.Where("description LIKE ?", "%"+userInput+"%").Find(&posts)
	db.Model(&User{}).Select([]string{"firstname", "lastname", "img_url", "username", "id"}).Where("firstname LIKE ?", "%"+userInput+"%").Or("lastname LIKE ?", "%"+userInput+"%").Or("username LIKE ?", "%"+userInput+"%").Find(&users)
	for i := 0; i < len(users); i++ {
		user := make(map[string]string)
		user["Firstname"] = users[i].Firstname
		user["Lastname"] = users[i].Lastname
		user["ImgUrl"] = users[i].ImgUrl
		user["Username"] = users[i].Username
		user["IsFriend"] = findFriendShip(users[i].ID, userId)
		userResults = append(userResults, user)
	}
	for i := 0; i < len(posts); i++ {
		post := getPostByPostId(posts[i].ID, userId)
		post.IsFriend = findFriendShip(posts[i].OwnerId, userId)
		postResult = append(postResult, post)
	}
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(map[string]interface{}{"Users": userResults, "Posts": postResult})
	handleError(err)
	w.Write(e)
}
