package middlewares

import (
	"errors"
	"mrrizal/wallet-service/app/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	temp := c.Get("Authorization")
	if !strings.HasPrefix(temp, "Token") {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(errors.New("Invalid token.")))
	}

	if len(strings.Split(temp, " ")) != 2 {
		c.SendStatus(400)
		return c.JSON(models.ErrResponse(errors.New("Invalid token.")))
	}

	return c.Next()
}
