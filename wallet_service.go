package main

import (
	"fmt"
	"log"
	"mrrizal/wallet-service/app/configs"
	"mrrizal/wallet-service/app/database"
	"mrrizal/wallet-service/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type AppInstance struct {
	app *fiber.App
	db  database.DB
}

func NewAppInstance(config configs.Config) AppInstance {
	db, err := database.NewDB(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	return AppInstance{app: app, db: db}
}

func main() {
	config := configs.GetConfig()
	app := NewAppInstance(*config)
	defer app.db.Close()

	routes.SetupRoutes(app.app, app.db)
	app.app.Listen(fmt.Sprintf(":%s", config.Port))
}
