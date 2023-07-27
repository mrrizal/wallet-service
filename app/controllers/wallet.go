package controllers

import (
	"mrrizal/wallet-service/app/service"
	"mrrizal/wallet-service/app/validators"
)

type WalletController struct {
	userTokenValidator validators.UserTokenValidator
	walletService      service.WalletService
}

func NewWalletController(
	userTokenValidator validators.UserTokenValidator,
	walletService service.WalletService) WalletController {
	return WalletController{
		userTokenValidator: userTokenValidator,
		walletService:      walletService,
	}
}

func (u *WalletController) Enable(token string) (map[string]interface{}, error) {
	userID, err := u.userTokenValidator.ValidateToken(token)
	if err != nil {
		return map[string]interface{}{}, err
	}

	walletData, err := u.walletService.Enable(userID, token)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return walletData, nil
}
