package transfer

import (
	rand "math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	balance "github.com/jaysonmulwa/golang-be-service/internal/handler/balance"
	"github.com/jaysonmulwa/golang-be-service/internal/model"
	converter "github.com/jaysonmulwa/golang-be-service/pkg/conversion_service"
	helper "github.com/jaysonmulwa/golang-be-service/pkg/helper_service"
)


func Transfer (c *fiber.Ctx) error {
	
	type TransferRequest struct {
		From int `json : "from"`
		Username string `json : "username"`
		Amount float64 `json : "amount"`
	}

	var input TransferRequest

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	helper := helper.Helper{}
	_from := input.From
	_to, _ := helper.GetIDFromUsername(input.Username)
	_amount := input.Amount

	fromCurrency, _ := helper.GetDefaultCurrency(_from)
	toCurrency, _ := helper.GetDefaultCurrency(_to)

	if fromCurrency == "" || toCurrency == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Default Currency of either parties is missing. Please update", "data": nil})
	}

	status, err := transferCash(_from, fromCurrency, _to, toCurrency, _amount)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error Transfering cash", "data": status})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cash trasnfered", "data": status})
}


func transferCash(from int, fromCurrency string, to int, toCurrency string, amount float64) (bool, error) {

	var new_amount float64
	var err error
	time_now := time.Now().Format("2006-01-02 15:04:05")
	new_amount = amount
	helper := helper.Helper{}
	if fromCurrency != toCurrency {
		conversion := converter.Converter{}
		new_amount, _ = conversion.Convert(amount, fromCurrency, toCurrency)
	}


	//Adjust Balance of Sender
	balanceModel, _ := balance.FetchBalance(from)
	balanceAmount := balanceModel.Amount
	newBalanceAmount := balanceAmount - new_amount
	newSenderBalance := model.Balance{
		Balance_id:   rand.Intn(1000000),
		User_id:      from,
		Amount:       newBalanceAmount,
		Last_updated: time_now,
	}
	_ = helper.AdjustBalance(newSenderBalance, from)


	//Adjust Balance of Receiver
	balanceModel_2, _ := balance.FetchBalance(to)
	balanceAmount_2 := balanceModel_2.Amount
	newBalanceAmount_2 := balanceAmount_2 + new_amount
	newReceiverBalance := model.Balance{
		Balance_id:   rand.Intn(1000000),
		User_id:      from,
		Amount:       newBalanceAmount_2,
		Last_updated: time_now,
	}
	_ = helper.AdjustBalance(newReceiverBalance, to)

	transaction_id := rand.Intn(1000000)
	entry_id := rand.Intn(1000000)
	newConversion := converter.Converter{}
	in_usd, err := newConversion.Convert(amount, fromCurrency, "USD") 
	
	//Record Transaction DR
	DR := model.Transaction{
		Transaction_id  : transaction_id,
		Entry_id        : entry_id,
		Transaction_type : "Transfer",
		Entry            : "DR",
		Amount           : amount,
		Currency        : fromCurrency,
		Amount_in_usd    : in_usd,
		Transaction_date : time_now,
		User_id          : from,
		Initiated_by     : from,
	}
	_ = helper.RecordTransaction(DR)



	//Record Transaction CR
	transaction_id = rand.Intn(1000000)
	CR := model.Transaction{
		Transaction_id   : transaction_id,
		Entry_id         : entry_id,
		Transaction_type : "Transfer",
		Entry            : "CR",
		Amount           : new_amount,
		Currency         : toCurrency,
		Amount_in_usd    : in_usd,
		Transaction_date : time_now,
		User_id          : to,
		Initiated_by     : to,
	}
	_ = helper.RecordTransaction(CR)
	
	return true, err
}

