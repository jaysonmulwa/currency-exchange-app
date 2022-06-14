package helper_service

import (
	db "github.com/jaysonmulwa/golang-be-service/internal/database"
	"github.com/jaysonmulwa/golang-be-service/internal/model"
)

type HelperService interface {
	GetIDFromUsername(username string) (int, error)
	AdjustBalance(balance model.Balance, user_id int) (error)
	RecordTransaction(transaction model.Transaction) (error)
}

type Helper struct {
	err error
}

func (h *Helper)GetIDFromUsername(username string) (int, error) {
	var user model.User
	_db := db.GetConnection().DB
	if result := _db.Where("username = ?", username).First(&user); result.Error != nil {
		h.err = result.Error
		return 0, result.Error
	}
	return user.User_id, nil
}

func (h *Helper)AdjustBalance(balance model.Balance, user_id int) (error) {
	_db := db.GetConnection().DB
	if err := _db.Create(&balance).Error; err != nil {
		return err
	}
	return nil
}

func (h *Helper)RecordTransaction(transaction model.Transaction) (error) {
	_db := db.GetConnection().DB
	if err := _db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func (h *Helper)GetDefaultCurrency(user_id int) (string, error) {
	var user model.User
	_db := db.GetConnection().DB
	if result := _db.Where("user_id = ?", user_id).First(&user); result.Error != nil {
		return "", result.Error
	}
	return user.Default_currency, nil
} 
