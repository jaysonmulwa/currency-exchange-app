package transact

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jaysonmulwa/golang-be-service/internal/handler/balance"
	"github.com/jaysonmulwa/golang-be-service/internal/model"
	converter "github.com/jaysonmulwa/golang-be-service/pkg/conversion_service"
	helper "github.com/jaysonmulwa/golang-be-service/pkg/helper_service"
)

func Transact (c *fiber.Ctx) error {

	type TransactRequest struct {
		Username string `json : "username"`
		Amount float64 `json : "amount"`
		Entry string `json : "entry"`
	}

	var input TransactRequest

	if err := c.BodyParser(input); err != nil {
		return err
	}

	helper := helper.Helper{}
	user_id, err := helper.GetIDFromUsername(input.Username)
	fromCurrency, err := helper.GetDefaultCurrency(user_id)
	amount := input.Amount
	entry := input.Entry
	time_now := time.Now().Format("2006-01-02 15:04:05")

	if user_id < 0 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	transaction_id := rand.Seed(time.Now().UnixNano())
	entry_id := rand.Seed(time.Now().UnixNano())
	newConversion := converter.Converter{}
	in_usd, err := newConversion.Convert(amount, fromCurrency, "USD") 

	//Record Transaction DR
	Entry := model.Transaction{
		Transaction_id  : transaction_id,
		Entry_id        : entry_id,
		Transaction_type : "Transact",
		Entry            : entry,
		Amount           : amount,
		Currency        : fromCurrency,
		Amount_in_usd    : in_usd,
		Transaction_date : time_now,
		User_id          : user_id,
		Initiated_by     : user_id,
	}
	helper.RecordTransaction(Entry)
	

	//Adjust balance
	balanceModel, err := balance.FetchBalance(user_id)
	balanceAmount := balanceModel.Amount
	newBalanceAmount := 0.0

	if entry == "DR" {
		newBalanceAmount = balanceAmount - amount
	} else {
		newBalanceAmount = balanceAmount + amount
	}
	newSenderBalance := model.Balance{
		Balance_id:   rand.Seed(time.Now().UnixNano()),
		User_id:      user_id,
		Amount:       newBalanceAmount,
		Last_updated: time_now,
	}
	helper.AdjustBalance(newSenderBalance, user_id)
	

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error occurred", "data": nil})
	}

	return nil
}