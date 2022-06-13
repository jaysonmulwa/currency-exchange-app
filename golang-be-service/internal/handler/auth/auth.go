package auth

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/golang-be-service/internal/database"
	"github.com/jaysonmulwa/golang-be-service/internal/model"
)

func Login (c *fiber.Ctx) error {

	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(input); err != nil {
		return err
	}

	_username := input.Username
	_password := input.Password

	user, err := getUserByCredentials(_username, _password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid username or password", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Login successful", "data": user})
}


func Register (c *fiber.Ctx) error {
	
	type RegisterInput struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Password string `json:"password"`
		Default_currency string `json:"default_currency"`
	}

	var input RegisterInput

	if err := c.BodyParser(input); err != nil {
		return err
	}

	_username := input.Username
	_email := input.Email
	_firstname := input.Firstname
	_lastname := input.Lastname
	_password := input.Password
	_default_currency := input.Default_currency

	user, err := createUser(_username, _email, _firstname, _lastname, _password, _default_currency)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error creating user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User created successfully", "data": user})

}

func getUserByCredentials(username string, password string) model.User, error {

	user := new(model.User)
	_db := db.GetConnection().DB
	if result := _db.Where("username = ?", username).Where("password = ?", password).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func createUser(username string, email string, firstname string, lastname string, password string, default_currency string) model.User, error {
	user := new(model.User)
	_db := db.GetConnection().DB
	if result := _db.Create(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}