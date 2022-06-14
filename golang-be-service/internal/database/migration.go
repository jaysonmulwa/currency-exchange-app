package database

import (
	"github.com/jaysonmulwa/golang-be-service/internal/model"
	"github.com/jinzhu/gorm"
)

//Migrate DB - migrates the database and create our comment table
func MigrateDB(db *gorm.DB) error {

	if result := db.AutoMigrate(&model.Balance{}); result.Error != nil {
		return result.Error
	}

	if result := db.AutoMigrate(&model.User{}); result.Error != nil {
		return result.Error
	}

	if result := db.AutoMigrate(&model.Transaction{}); result.Error != nil {
		return result.Error
	}
	return nil
}