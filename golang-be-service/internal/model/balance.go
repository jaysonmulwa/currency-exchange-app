package model

import "github.com/jinzhu/gorm"

type Balance struct {
	gorm.Model
	balance_id   int    `json:"balance_id"`
	user_id      int    `json:"user_id"`
	amount       int    `json:"amount"`
	currency     string `json:"currency"`
	last_updated string `json:"last_updated"`
}