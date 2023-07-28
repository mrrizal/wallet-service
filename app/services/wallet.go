package services

import (
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/repositories"
	"time"

	"github.com/google/uuid"
)

type WalletService interface {
	Enable(userID, token string) (map[string]interface{}, error)
	Update(wallet *models.Wallet, status models.WalletStatus) (map[string]interface{}, error)
	GetWallet(userID string) (map[string]interface{}, error)
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

func (w *walletService) Update(wallet *models.Wallet, status models.WalletStatus) (map[string]interface{}, error) {
	wallet.Status = status
	if err := w.walletRepository.Update(wallet); err != nil {
		return make(map[string]interface{}), err
	}
	return models.ParseWallet(*wallet), nil
}

func (w *walletService) GetWallet(userID string) (map[string]interface{}, error) {
	wallet := w.walletRepository.GetWallet(userID)
	return models.ParseWallet(wallet), nil
}
