package services

import (
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/repositories"
	"time"

	"github.com/google/uuid"
)

type WalletService interface {
	Enable(userID, token string) (map[string]interface{}, error)
	ReEnable(wallet *models.Wallet) (map[string]interface{}, error)
	GetWallet(userID string) models.Wallet
}

type walletService struct {
	walletRepository repositories.WalletRepository
}

func NewWalletService(walletRepository repositories.WalletRepository) walletService {
	return walletService{walletRepository: walletRepository}
}

func (w *walletService) Enable(userID, token string) (map[string]interface{}, error) {
	walletData := models.Wallet{
		ID:        uuid.NewString(),
		UserID:    userID,
		Status:    models.WalletStatusEnabled,
		EnabledAt: time.Now(),
		Balance:   0,
	}

	if err := w.walletRepository.Enable(walletData); err != nil {
		return make(map[string]interface{}), err
	}

	return models.ParseWallet(walletData), nil
}

func (w *walletService) ReEnable(wallet *models.Wallet) (map[string]interface{}, error) {
	wallet.Status = models.WalletStatusEnabled
	if err := w.walletRepository.ReEnable(wallet.ID, wallet.Status); err != nil {
		return make(map[string]interface{}), err
	}
	return models.ParseWallet(*wallet), nil
}

func (w *walletService) GetWallet(userID string) models.Wallet {
	return w.walletRepository.GetWallet(userID)
}
