package controllers

import (
	"errors"
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/services"
	"mrrizal/wallet-service/app/validators"
	"time"
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

	walletDataObj, haveWallet := w.walletValidator.IsHaveWallet(userID)
	if err != nil {
		return make(map[string]interface{}), err
	}

	if haveWallet {
		if w.walletValidator.IsWalletEnabled(walletDataObj) {
			return make(map[string]interface{}), errors.New("Already enabled")
		}
		return w.walletService.Update(&walletDataObj, models.WalletStatusEnabled)
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

	wallet, haveWallet := w.walletValidator.IsHaveWallet(userID)

	if !w.walletValidator.IsWalletEnabled(wallet) || !haveWallet {
		return make(map[string]interface{}), errors.New("Wallet Disabled")
	}

	return models.ParseWallet(wallet), nil
}

func (w *WalletController) Disable(token string) (map[string]interface{}, error) {
	userID, err := w.userTokenValidator.ValidateToken(token)
	if err != nil {
		return map[string]interface{}{}, err
	}

	walletDataObj, haveWallet := w.walletValidator.IsHaveWallet(userID)
	if err != nil {
		return make(map[string]interface{}), err
	}

	if !haveWallet {
		return make(map[string]interface{}), errors.New("User doens't have wallet.")
	}

	walletDataObj.Status = models.WalletStatusDisabled
	walletData, err := w.walletService.Update(&walletDataObj, models.WalletStatusDisabled)
	if err != nil {
		return map[string]interface{}{}, err
	}

	walletData["wallet"].(map[string]interface{})["disabled_at"] = time.Now().Format("2006-01-02T15:04:05-07:00")
	delete(walletData["wallet"].(map[string]interface{}), "enabled_at")

	return walletData, nil
}
