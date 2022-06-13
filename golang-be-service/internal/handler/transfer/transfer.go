package transfer

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/golang-be-service/internal/database"
	"github.com/jaysonmulwa/golang-be-service/internal/model"
	converter "github.com/jaysonmulwa/golang-be-service/pkg/conversion_service"
)


func Transfer (c *fiber.Ctx) error {
	
	type TransferRequest struct {
		From int `json : "from"`
		To int `json : "to"`
		Amount float64 `json : "amount"`
	}

	var input TransferRequest

	if err := c.BodyParser(input); err != nil {
		return err
	}

	_from := input.From
	_to := input.To
	_amount := input.Amount

	fromCurrency, err := getDefaultCurrency(_from)
	toCurrency, err := getDefaultCurrency(_to)

	if fromCurrency == nil || toCurrency == nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Default Currency of either parties is missing. Please update", "data": nil})
	}

	status, err := transferCash(_from, fromCurrency, _to, toCurrency, amount)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error Transfering cash", "data": status})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cash trasnfered", "data": status})
}

func getDefaultCurrency(id int) (string, error) {
	return "", nil
} 

func transferCash(from int, fromCurrency string, to int, toCurrency string, amount float64) (bool, error) {

	new_amount := 0
	if fromCurrency != toCurrency {
		new_amount, err = converter.Convert(amount, fromCurrency, toCurrency)
	}
	
	
	var balance model.Balance
	_db := db.GetConnection().DB
	if result := _db.Where("user_id = ?", user_id).First(&balance); result.Error != nil {
		return &balance, result.Error
	}
	return &balance, nil
	
}