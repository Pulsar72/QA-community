package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName    string `form:"username"  json:"username" binding:"required"`
	Email       string `form:"email" json:"email" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	PhoneNumber string `form:"phonenumber" json:"phonenumber" binding:"required"`
}

type UserRegister struct {
	UserName    string `form:"username"  json:"username" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	Email       string `form:"email" json:"email" binding:"required"`
	PhoneNumber string `form:"phonenumber" json:"phonenumber" binding:"required"`
}

type PhoneLogin struct {
	PhoneNumber string `form:"phonenumber" json:"phonenumber" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
}

type UserNameLogin struct {
	UserName string `form:"username"  json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type EmailLogin struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
