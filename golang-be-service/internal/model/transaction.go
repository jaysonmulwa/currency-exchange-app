package model

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Transaction_id   int    `json:"transaction_id"`
	Entry_id         int    `json:"entry_id"`
	Transaction_type string `json:"transaction_type"`
	Entry            string `json:"entry"`
	Amount           float64    `json:"amount"`
	Currency         string `json:"currency"`
	Amount_in_usd    float64    `json:"amount_in_usd"`
	Transaction_date string `json:"transaction_date"`
	User_id          int    `json:"user_id"`
	Initiated_by     int `json:"initiated_by"`
}