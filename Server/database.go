package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var db *gorm.DB

func connectDb() {
	var err error
	dsn := "root:mehdi007@tcp(127.0.0.1:3306)/detweak?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	

}
func findEmail( email string)(bool){
	var users User
     result:=db.Where("email = ?", email).First(&users)
	if result.Error.Error() == "record not found"{
		fmt.Println(false)
		return false
		
	}else{
		fmt.Println(true)
		return true
	}
	
}