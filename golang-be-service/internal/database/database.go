package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBInstance struct {
	DB *gorm.DB
}

var (
	database DBInstance
)

// NewDatabase - returns a pointer to a new database connection
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Connecting to database...")
	dbUsername := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	tcpString := os.Getenv("MYSQL_DB_CONTAINER")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUsername, dbPassword, tcpString, dbName)
	db, err := gorm.Open("mysql", connectionString)

	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	database = DBInstance{
		DB: db,
	}

	return db, nil
}

func GetConnection() *DBInstance {
	return &database
}
