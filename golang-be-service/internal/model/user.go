package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	user_id          int    `json:"user_id"`
	username         string `json:"username"`
	email            string `json:"email"`
	firstname        string `json:"firstname"`
	lastname         string `json:"lastname"`
	password         string `json:"password"`
	profile_pic      string `json:"profile_pic"`
	default_currency string `json:"default_currency"`
}
