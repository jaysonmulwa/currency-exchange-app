package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
	auth "github.com/jaysonmulwa/golang-be-service/internal/handler/auth"
	balance "github.com/jaysonmulwa/golang-be-service/internal/handler/balance"
	profile "github.com/jaysonmulwa/golang-be-service/internal/handler/profile"
	transact "github.com/jaysonmulwa/golang-be-service/internal/handler/transact"
	transfer "github.com/jaysonmulwa/golang-be-service/internal/handler/transfer"
)
func SetupRoutes() {

	app := fiber.New()

	// Default config
	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "*",
	}))
	
	app.Post("/api/v1/login", auth.Login)
	app.Post("/api/v1/signup", auth.Register)
	app.Get("/api/v1/profile/:user_id", profile.GetProfile)
	app.Put("/api/v1/profile/:user_id", profile.UpdateProfile)

	app.Get("/api/v1/balance/:user_id", balance.GetBalance)
	app.Post("/api/v1/transfer", transfer.Transfer)
	app.Post("/api/v1/transact", transact.Transact)

	if err := app.Listen(":3001"); err != nil {
		fmt.Println(err)
		panic(err)
	}

}