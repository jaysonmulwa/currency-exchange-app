package main

import (
	"fmt"

	http "github.com/jaysonmulwa/golang-be-service/internal/http"
)

type App struct {

}

func(a *App) Run () error {
	http.SetupRoutes();
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
