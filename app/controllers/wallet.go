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

func (u *WalletController) Enable(token string) (map[string]interface{}, error) {
	userID, err := u.userTokenValidator.ValidateToken(token)
	if err != nil {
		return map[string]interface{}{}, err
	}

	walletDataObj := u.walletValidator.IsHaveWallet(userID)
	if err != nil {
		return make(map[string]interface{}), err
	}

	if walletDataObj.ID != "" {
		if walletDataObj.Status == models.WalletStatusDisabled {
			return u.walletService.ReEnable(&walletDataObj)
		} else if walletDataObj.Status == models.WalletStatusEnabled {
			return make(map[string]interface{}), errors.New("Already enabled")
		}
	}

	walletData, err := u.walletService.Enable(userID, token)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return walletData, nil
}
