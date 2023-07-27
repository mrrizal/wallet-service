package controllers

import (
	"errors"
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/services"
	"mrrizal/wallet-service/app/validators"
)

type WalletController struct {
	userTokenValidator validators.UserTokenValidator
	walletValidator    validators.WalletValidator
	walletService      services.WalletService
}

func NewWalletController(
	userTokenValidator validators.UserTokenValidator,
	walletValidator validators.WalletValidator,
	walletService services.WalletService) WalletController {
	return WalletController{
		userTokenValidator: userTokenValidator,
		walletValidator:    walletValidator,
		walletService:      walletService,
	}
}

func (w *WalletController) Enable(token string) (map[string]interface{}, error) {
	userID, err := w.userTokenValidator.ValidateToken(token)
	if err != nil {
		return map[string]interface{}{}, err
	}

	walletDataObj := w.walletValidator.IsHaveWallet(userID)
	if err != nil {
		return make(map[string]interface{}), err
	}

	if walletDataObj.ID != "" {
		if walletDataObj.Status == models.WalletStatusDisabled {
			return w.walletService.ReEnable(&walletDataObj)
		} else if walletDataObj.Status == models.WalletStatusEnabled {
			return make(map[string]interface{}), errors.New("Already enabled")
		}
	}

	walletData, err := w.walletService.Enable(userID, token)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return walletData, nil
}

func (w *WalletController) GetWallet(token string) (map[string]interface{}, error) {
	userID, err := w.userTokenValidator.ValidateToken(token)
	if err != nil {
		return map[string]interface{}{}, err
	}

	wallet := w.walletService.GetWallet(userID)
	if wallet.ID == "" || wallet.Status == models.WalletStatusDisabled {
		return make(map[string]interface{}), errors.New("Wallet Disabled")
	}

	return models.ParseWallet(wallet), nil
}
