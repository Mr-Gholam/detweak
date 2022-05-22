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
	db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

}
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
func getUsernameById(userId uint) string {
	var user User
	db.Where("id = ?", userId).First(&user)
	return user.Username
}
