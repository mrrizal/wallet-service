package controllers

import (
	"errors"
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/services"
	"mrrizal/wallet-service/app/validators"
)

type TransactionController struct {
	userTokenValidator validators.UserTokenValidator
	walletValidator    validators.WalletValidator
	transactionService services.TransactionService
}

func NewTransactionController(userTokenValidator validators.UserTokenValidator,
	walletValidator validators.WalletValidator,
	transactionService services.TransactionService) TransactionController {
	return TransactionController{
		userTokenValidator: userTokenValidator,
		walletValidator:    walletValidator,
		transactionService: transactionService,
	}
}

func (t *TransactionController) Deposit(token string,
	transactionReq models.TransactionRequest) (map[string]interface{}, error) {
	userID, err := t.userTokenValidator.ValidateToken(token)
	if err != nil {
		return map[string]interface{}{}, err
	}

	walletDataObj, haveWallet := t.walletValidator.IsHaveWallet(userID)
	walletEnabled := t.walletValidator.IsWalletEnabled(walletDataObj)

	if !haveWallet {
		return make(map[string]interface{}), errors.New("Wallet Disabled")
	}

	if !walletEnabled {
		return make(map[string]interface{}), errors.New("Wallet Disabled")
	}

	transactionData, err := t.transactionService.Deposit(walletDataObj.ID, userID, transactionReq)
	if err != nil {
		return make(map[string]interface{}), err
	}
	return transactionData, nil
}

func (t *TransactionController) Withdraw(token string,
	transactionReq models.TransactionRequest) (map[string]interface{}, error) {
	userID, err := t.userTokenValidator.ValidateToken(token)
	if err != nil {
		return map[string]interface{}{}, err
	}

	walletDataObj, haveWallet := t.walletValidator.IsHaveWallet(userID)
	walletEnabled := t.walletValidator.IsWalletEnabled(walletDataObj)

	if !haveWallet {
		return make(map[string]interface{}), errors.New("Wallet Disabled")
	}

	if !walletEnabled {
		return make(map[string]interface{}), errors.New("Wallet Disabled")
	}

	transactionData, err := t.transactionService.Withdraw(walletDataObj.ID, userID, transactionReq)
	if err != nil {
		return make(map[string]interface{}), err
	}
	return transactionData, nil
}
