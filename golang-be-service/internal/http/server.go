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

	app.Post("/api/v1/login", auth.Login)
	app.Post("/api/v1/register", auth.Register)
	app.Get("/api/v1/profile", profile.GetProfile)
	app.Put("/api/v1/profile", profile.UpdateProfile)

	app.Get("/api/v1/balance", balance.GetBalance)
	app.Post("/api/v1/transfer", transfer.Transfer)
	app.Post("/api/v1/transact", transact.Transact)

	if err := app.Listen(":3001"); err != nil {
		fmt.Println(err)
		panic(err)
	}

}