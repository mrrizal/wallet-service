package routes

import (
	"mrrizal/wallet-service/app/controllers"
	"mrrizal/wallet-service/app/database"
	"mrrizal/wallet-service/app/handler"
	"mrrizal/wallet-service/app/middlewares"
	"mrrizal/wallet-service/app/repositories"
	"mrrizal/wallet-service/app/services"
	"mrrizal/wallet-service/app/validators"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db database.DB) {
	// repository
	userTokenRepository := repositories.NewUserTokenRepository(db)
	walletRepository := repositories.NewWalletRepository(db)

	// service
	userTokenService := services.NewUserTokenService(&userTokenRepository)
	walletService := services.NewWalletService(&walletRepository)

	// validator
	userTokenValidator := validators.NewUserTokenValidator(&userTokenService)
	walletValidator := validators.NewWalletValidator(&walletService)

	// controller
	userTokenController := controllers.NewUserTokenController(&userTokenValidator, &userTokenService)
	walletController := controllers.NewWalletController(&userTokenValidator, &walletValidator, &walletService)

	// handler
	userTokenHandler := handler.NewUserTokenHandler(userTokenController)
	walletHandler := handler.NewWalletHandler(walletController)

	// routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/init", userTokenHandler.Init)
	v1.Post("/wallet", middlewares.AuthMiddleware, walletHandler.Enable)
	v1.Get("/wallet", middlewares.AuthMiddleware, walletHandler.GetWallet)
}
