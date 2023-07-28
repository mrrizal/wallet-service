package handler

import (
	"mrrizal/wallet-service/app/controllers"
	"mrrizal/wallet-service/app/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionController controllers.TransactionController
}

func NewTransactionHandler(transactionController controllers.TransactionController) TransactionHandler {
	return TransactionHandler{transactionController: transactionController}
}

func (t *TransactionHandler) Deposit(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]
	amount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(err))
	}

	transactionRequest := models.TransactionRequest{
		Amount:      amount,
		ReferenceID: c.FormValue("reference_id"),
	}

	transactionData, err := t.transactionController.Deposit(token, transactionRequest)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(err))
	}

	c.SendStatus(201)
	return c.JSON(models.Response(transactionData, "success"))
}

func (t *TransactionHandler) Withdraw(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]
	amount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(err))
	}

	transactionRequest := models.TransactionRequest{
		Amount:      amount,
		ReferenceID: c.FormValue("reference_id"),
	}

	transactionData, err := t.transactionController.Withdraw(token, transactionRequest)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(err))
	}

	c.SendStatus(201)
	return c.JSON(models.Response(transactionData, "success"))
}

func (t *TransactionHandler) FetchAll(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]
	transactionData, err := t.transactionController.FetchAll(token)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(err))
	}

	c.SendStatus(200)
	return c.JSON(models.Response(transactionData, "success"))
}
