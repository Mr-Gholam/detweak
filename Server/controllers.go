package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Email  string `json:"email,omitempty"`
	Username string  `json:"username,omitempty"`
	Password   string   `json:"password,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Lastname string 	`json:"lastname,omitempty"`
	Bio string      `json:"bio,omitempty"`
	Location string `json:"location,omitempty"`
	ImgUrl string 
	Field string `json:"field,omitempty"`
	Language  string `json:"language,omitempty"`
	FrameWork string `json:"framework,omitempty"`
	Birthday time.Time `json:"birthday,omitempty"`

}
func post_signup(w http.ResponseWriter , r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var user User
    err := decoder.Decode(&user)
    if err != nil {
        panic(err)
    }
	findEmail(user.Email)
}
func post_login(w http.ResponseWriter , r *http.Request) {
	fmt.Println(r)
	}
func post_set_profile(w http.ResponseWriter , r *http.Request) {
fmt.Println(r)
}
func post_set_resume(w http.ResponseWriter , r *http.Request) {
	fmt.Println(r)
	}