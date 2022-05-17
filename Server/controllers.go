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
	Field     string    `json:"field,omitempty"`
	Language  string    `json:"language,omitempty"`
	FrameWork string    `json:"framework,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
}

type ErrorRespone struct {
	Error ErrorMessage `json:"error,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
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
	// hasing the password
	user.encryptPassword()
	// creating the user in datebase
	db.Create(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": user.Username, "id": user.ID, "email": user.Email})
	tokenStr, _ := token.SignedString(jwtSecret)
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenStr, HttpOnly: true, Secure: true, MaxAge: 3600 * 24 * 1, SameSite: http.SameSiteNoneMode})
	w.WriteHeader(http.StatusCreated)
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
	worngPassword := user.validatePassword(foundPassword, user.Password)
	if worngPassword {
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
}
func post_set_profile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
func post_set_resume(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
