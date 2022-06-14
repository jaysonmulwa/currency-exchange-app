package auth

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/golang-be-service/internal/database"
)

type User struct {
	User_id          int    `json:"user_id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Password         string `json:"password"`
	Profile_pic      string `json:"profile_pic"`
	Default_currency string `json:"default_currency"`
}

func Login (c *fiber.Ctx) error {

	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
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

	input := RegisterInput{} 

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	_user_id := rand.Intn(100000)
	_username := input.Username
	_email := input.Email
	_firstname := input.Firstname
	_lastname := input.Lastname
	_password := input.Password
	_default_currency := input.Default_currency

	user, err := createUser(_user_id, _username, _email, _firstname, _lastname, _password, _default_currency)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error creating user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User created successfully", "data": user})

}

func getUserByCredentials(username string, password string) (User, error) {

	user := User{}
	_db := db.GetConnection().DB
	if result := _db.Where("username = ?", username).Where("password = ?", password).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func createUser(userid int, username string, email string, firstname string, lastname string, password string, default_currency string) (User, error) {

	user := User{}
	user.User_id = userid
	user.Username = username
	user.Email = email
	user.Firstname = firstname
	user.Lastname = lastname
	user.Password = password
	user.Default_currency = default_currency

	_db := db.GetConnection().DB
	if result := _db.Create(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
	
}