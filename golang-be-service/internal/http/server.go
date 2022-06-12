package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)
func SetupRoutes() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err)
		panic(err)
	}

}