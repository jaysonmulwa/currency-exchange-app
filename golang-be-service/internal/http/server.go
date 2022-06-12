package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	auth "github.com/jaysonmulwa/golang-be-service/internal/handler/auth"
	balance "github.com/jaysonmulwa/golang-be-service/internal/handler/balance"
	profile "github.com/jaysonmulwa/golang-be-service/internal/handler/profile"
	transact "github.com/jaysonmulwa/golang-be-service/internal/handler/transact"
	transfer "github.com/jaysonmulwa/golang-be-service/internal/handler/transfer"
)
func SetupRoutes() {

	app := fiber.New()

	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
	app.Get("/profile", profile.GetProfile)
	app.Put("/profile", profile.UpdateProfile)

	app.Get("/balance", balance.GetBalance)
	app.Post("/transfer", transfer.Transfer)
	app.Post("/transact", transact.Transact)



	if err := app.Listen(":3001"); err != nil {
		fmt.Println(err)
		panic(err)
	}

}