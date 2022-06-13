package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	User_id          int    `json:"user_id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Password         string `json:"password"`
	Profile_pic      string `json:"profile_pic"`
	Default_currency string `json:"default_currency"`
}
