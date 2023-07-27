package validators

import (
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/services"
)

type WalletValidator interface {
	IsHaveWallet(userID string) models.Wallet
}

type walletValidator struct {
	walletService services.WalletService
}

func NewWalletValidator(walletService services.WalletService) walletValidator {
	return walletValidator{walletService: walletService}
}

func (w *walletValidator) IsHaveWallet(userID string) models.Wallet {
	return w.walletService.GetWallet(userID)
}
