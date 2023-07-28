package validators

import (
	"log"
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/services"
)

type WalletValidator interface {
	IsHaveWallet(userID string) (models.Wallet, bool)
	IsWalletEnabled(wallet models.Wallet) bool
}

type walletValidator struct {
	walletService services.WalletService
}

func NewWalletValidator(walletService services.WalletService) walletValidator {
	return walletValidator{walletService: walletService}
}

func (w *walletValidator) IsHaveWallet(userID string) (models.Wallet, bool) {
	walletData, err := w.walletService.GetWallet(userID)
	if err != nil {
		log.Fatal(err.Error())
	}
	wallet := models.MapTopWallet(walletData)
	return wallet, wallet.ID != ""
}

func (w *walletValidator) IsWalletEnabled(wallet models.Wallet) bool {
	if wallet.Status == models.WalletStatusDisabled {
		return false
	}

	return true
}
