package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": user.Username, "id": user.ID, "imgUrl": ""})
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": id, "imgUrl": ImgUrl})
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
	username := findUsernameById(userId)
	db.Model(&User{}).Where("id = ?", userId).Updates(User{ImgUrl: ImgUrl})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": userId, "imgUrl": ImgUrl})
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
	likedPost.jsonLikedPost(r)
	userId := getIdFromCookie(w, r)
	likedPost.UserId = userId
	result := getPostLikedByUserId(userId, likedPost.PostId)
	if result {
		dislikePost(likedPost.PostId)
		db.Where("user_id = ? AND post_id = ?", userId, likedPost.PostId).Delete(&LikedPost{})
		w.WriteHeader(http.StatusOK)
		e, err := json.Marshal(map[string]interface{}{"Added": false})
		handleError(err)
		w.Write(e)
	} else {
		likePost(likedPost.PostId)
		db.Create(&likedPost)
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
		username, firstname, lastname, imgUrl := findUserById(Ids[i]["targetId"])
		roomDetail.Firstname = firstname
		roomDetail.ImgUrl = imgUrl
		roomDetail.Username = username
		roomDetail.Lastname = lastname
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

// search
func get_search(w http.ResponseWriter, r *http.Request) {
	// userInput := mux.Vars(r)["userInput"]

}
