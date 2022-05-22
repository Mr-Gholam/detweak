package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("detweak")

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Bio       string `json:"bio,omitempty"`
	Location  string `json:"location,omitempty"`
	ImgUrl    string
	Field     string `json:"field,omitempty"`
	Language  string `json:"language,omitempty"`
	FrameWork string `json:"framework,omitempty"`
	Birthday  string `json:"birthday,omitempty"`
}

type ErrorRespone struct {
	Error ErrorMessage `json:"error,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}
type UserJson struct {
	Username string `json:"username,omitempty"`
	ImgUrl   string `json:"imgUrl,omitempty"`
}

func (user *User) json(r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	handleError(decoder.Decode(&user))
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
	user.json(r)

	// invalid input
	if user.Username == "" || user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"invalid username, password, email"}})
		handleError(err)
		w.Write(e)
		return
	}
	foundEmail, _, _, _ := findEmail(user.Email)
	// find duplicate email
	if len(foundEmail) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(ErrorRespone{ErrorMessage{"email already exists"}})
		handleError(err)
		w.Write(e)
		return
	}
	// find duplicate username
	if findUsername(user.Username) {
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": user.Username, "id": user.ID, "email": user.Email})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusCreated)
	e, err := json.Marshal(UserJson{Username: user.Username, ImgUrl: ""})
	handleError(err)
	w.Write(e)
}
func post_login(w http.ResponseWriter, r *http.Request) {
	var user User
	user.json(r)
	FoundEmail, foundPassword, id, username := findEmail(user.Email)
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": id, "email": user.Email})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusOK)
	e, err := json.Marshal(UserJson{Username: username, ImgUrl: ""})
	handleError(err)
	w.Write(e)
}
func post_set_profile(w http.ResponseWriter, r *http.Request) {
	var userData User
	userData.json(r)
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
	db.Model(&userData).Where("id = ?", userId).Updates(User{Firstname: userData.Firstname, Lastname: userData.Lastname, Bio: userData.Bio, Birthday: userData.Birthday, Location: userData.Location})
	w.WriteHeader(http.StatusOK)
}
func post_set_resume(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
