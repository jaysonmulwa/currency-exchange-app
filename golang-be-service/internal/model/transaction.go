package model

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	transaction_id   int    `json:"transaction_id"`
	entry_id         int    `json:"entry_id"`
	transaction_type string `json:"transaction_type"`
	entry            string `json:"entry"`
	amount           int    `json:"amount"`
	currency         string `json:"currency"`
	amount_in_usd    int    `json:"amount_in_usd"`
	transaction_date string `json:"transaction_date"`
	user_id          int    `json:"user_id"`
	initiated_by     string `json:"initiated_by"`
}