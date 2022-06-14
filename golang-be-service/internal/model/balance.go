package model

import "github.com/jinzhu/gorm"

type Balance struct {
	gorm.Model
	Balance_id   int    `json:"balance_id"`
	User_id      int    `json:"user_id"`
	Amount       float64   `json:"amount"`
	Currency     string `json:"currency"`
	Last_updated string `json:"last_updated"`
}