package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewDatabase - returns a pointer to a new database connection
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Connecting to database...")
	dbUsername := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbPort := os.Getenv("MYSQL_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require", dbHost, dbPort, dbUsername, dbName, dbPassword)
	fmt.Println(connectionString)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}