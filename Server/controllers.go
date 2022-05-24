package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

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
func (request *FriendJSON) jsonFriend(r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	handleError(decoder.Decode(&request))
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
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid username, password, email"}})
		handleError(err)
		w.Write(e)
		return
	}
	foundEmail := findDuplicateEmail(user.Email)
	// find duplicate email
	if foundEmail {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"email already exists"}})
		handleError(err)
		w.Write(e)
		return
	}
	// find duplicate username
	if findDuplicateUsername(user.Username) {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"username already exists"}})
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
	e, err := json.Marshal(UserJSON{Username: user.Username, ImgUrl: ""})
	handleError(err)
	w.Write(e)
}
func post_login(w http.ResponseWriter, r *http.Request) {
	var user User
	user.jsonUser(r)
	FoundEmail, foundPassword, id, username, ImgUrl := findUserByEmail(user.Email)
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid email or password"}})
		handleError(err)
		w.Write(e)
		return
	}
	if len(FoundEmail) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"email not found"}})
		handleError(err)
		w.Write(e)
		return
	}
	wrongPassword := user.validatePassword(foundPassword, user.Password)
	if wrongPassword {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"incorrect password"}})
		handleError(err)
		w.Write(e)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": id, "imgUrl": ImgUrl})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(UserJSON{Username: username, ImgUrl: ImgUrl})
	handleError(err)
	w.Write(e)
}
func post_set_profile(w http.ResponseWriter, r *http.Request) {
	var userData User
	userData.jsonUser(r)
	userId := getIdFromCookie(w, r)
	if userData.Firstname == "" || userData.Lastname == "" || userData.Bio == "" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid Firstname , Lastname, Bio"}})
		handleError(err)
		w.Write(e)
		return
	}
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid Id"}})
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
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid Id"}})
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
	e, err := json.Marshal(UserJSON{Username: username, ImgUrl: ImgUrl})
	handleError(err)
	w.Write(e)
}
func post_set_resume(w http.ResponseWriter, r *http.Request) {
	var userData User
	userData.jsonUser(r)
	userId := getIdFromCookie(w, r)
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid Id"}})
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
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid Id"}})
		handleError(err)
		w.Write(e)
		return
	}
	if post.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid Description"}})
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
	e, err := json.Marshal(PostImg{ImgUrl: imgUrl})
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
		e, err := json.Marshal(LikedPostJSON{Added: false})
		handleError(err)
		w.Write(e)
	} else {
		likePost(likedPost.PostId)
		db.Create(&likedPost)
		w.WriteHeader(http.StatusOK)
		e, err := json.Marshal(LikedPostJSON{Added: true})
		handleError(err)
		w.Write(e)
	}
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
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"Friends"}})
		handleError(err)
		w.Write(e)
		return
	}
	if status == "Pending" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"Pending"}})
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
	var request FriendJSON
	request.jsonFriend(r)
	db.Model(&FriendShip{}).Where("id = ?", request.ID).Update("status", true)
	w.WriteHeader(http.StatusOK)
}
func post_reject_request(w http.ResponseWriter, r *http.Request) {
	var request FriendJSON
	request.jsonFriend(r)
	db.Where("id = ?", request.ID).Delete(&FriendShip{})
	w.WriteHeader(http.StatusOK)
}

// search
func get_search(w http.ResponseWriter, r *http.Request) {
	// userInput := mux.Vars(r)["userInput"]

}
