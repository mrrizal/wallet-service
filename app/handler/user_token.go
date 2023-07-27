package handler

import (
	"mrrizal/wallet-service/app/controllers"
	"mrrizal/wallet-service/app/models"

	fiber "github.com/gofiber/fiber/v2"
)

type UserTokenHandler struct {
	userTokenController controllers.UserTokenController
}

func NewUserTokenHandler(userTokenController controllers.UserTokenController) UserTokenHandler {
	return UserTokenHandler{userTokenController: userTokenController}
}

func (u *UserTokenHandler) Init(c *fiber.Ctx) error {
	customerXID := c.FormValue("customer_xid")
	token, err := u.userTokenController.Init(customerXID)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(err))
	}

	c.SendStatus(201)
	return c.JSON(models.Response(map[string]interface{}{"token": token}, "success"))
}
