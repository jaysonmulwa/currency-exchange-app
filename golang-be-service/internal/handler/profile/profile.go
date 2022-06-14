package profile

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/golang-be-service/internal/database"
	"github.com/jaysonmulwa/golang-be-service/internal/model"
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

func GetProfile (c *fiber.Ctx) error {

	user_id := c.Params("user_id")
	i_user_id , _ := strconv.Atoi(user_id)

	fmt.Println(user_id, i_user_id)
	result, err := GetUserDetails(i_user_id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error fetching user details", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User details fetched", "data": result})
}


func UpdateProfile (c *fiber.Ctx) error {

	user_id := c.Params("user_id")
	i_user_id , _ := strconv.Atoi(user_id)

	type ProfileRequest struct {
		Default_currency string `json:"default_currency"`
	}

	var input ProfileRequest

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	err := UpdateUserDetails(i_user_id, input.Default_currency)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error updating user details", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User details updated", "data": nil})
}

func GetUserDetails(user_id int) (User, error) {
	user := User{}
	_db := db.GetConnection().DB
	if result := _db.Where("user_id = ?", user_id).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func UpdateUserDetails(user_id int, default_currency string) error {
	_db := db.GetConnection().DB
	if err := _db.Model(&model.User{}).Where("user_id = ?", user_id).Update("default_currency", default_currency).Error; err != nil {
		return err
	}
	return nil
}