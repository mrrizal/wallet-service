package handler

import (
	"mrrizal/wallet-service/app/controllers"
	"mrrizal/wallet-service/app/models"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

type WalletHandler struct {
	walletController controllers.WalletController
}

func NewWalletHandler(walletController controllers.WalletController) WalletHandler {
	return WalletHandler{walletController: walletController}
}

func (w *WalletHandler) Enable(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]
	walletData, err := w.walletController.Enable(token)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(err))
	}
	c.SendStatus(201)
	return c.JSON(models.Response(walletData, "success"))
}

func (w *WalletHandler) GetWallet(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]
	walletData, err := w.walletController.GetWallet(token)
	if err != nil {
		c.SendStatus(404)
		return c.JSON(models.ErrResponse(err))
	}
	c.SendStatus(200)
	return c.JSON(models.Response(walletData, "success"))
}

func (w *WalletHandler) Disable(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]
	isDisable := c.FormValue("is_disabled")
	if isDisable == "true" {
		walletData, err := w.walletController.Disable(token)
		if err != nil {
			c.SendStatus(400)
			return c.JSON(models.ErrResponse(err))
		}
		c.SendStatus(200)
		return c.JSON(models.Response(walletData, "success"))
	}
	return c.SendStatus(204)
}
