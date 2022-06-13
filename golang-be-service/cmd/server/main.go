package main

import (
	"fmt"

	"github.com/jaysonmulwa/golang-be-service/internal/database"
	http "github.com/jaysonmulwa/golang-be-service/internal/http"
	"github.com/jinzhu/gorm"
)

type App struct {
	db *gorm.DB
	err error
}

func(a *App) Run () error {
	a.setupDatabase()
	a.setupMigrations()
	http.SetupRoutes();
	return a.err
}

func(a *App) setupDatabase() {

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		a.err = err
	}
	a.db = db
	a.err = nil
	
}

func(a *App) setupMigrations() {

	err := database.MigrateDB(a.db)
	if err != nil {
		a.err = err
	}
	a.err = nil

}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}