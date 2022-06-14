package balance

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/golang-be-service/internal/database"
)

type Balance struct {
	Balance_id   int    `json:"balance_id"`
	User_id      int    `json:"user_id"`
	Amount       float64   `json:"amount"`
	Currency     string `json:"currency"`
	Last_updated string `json:"last_updated"`
}

func GetBalance (c *fiber.Ctx) error {
	user_id := c.Params("user_id")
	i_user_id , _ := strconv.Atoi(user_id)

	balance, err := FetchBalance(i_user_id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error fetching balance", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Balance fetched", "data": balance})
}

func FetchBalance(user_id int) (Balance, error) {
	
	balance := Balance{}
	_db := db.GetConnection().DB
	if result := _db.Where("user_id = ?", user_id).First(&balance); result.Error != nil {
		return balance, result.Error
	}
	return balance, nil
}