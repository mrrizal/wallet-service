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

func (t *TransactionController) validate(token string) (models.Wallet, string, error) {
	userID, err := t.userTokenValidator.ValidateToken(token)
	if err != nil {
		return models.Wallet{}, "", err
	}

	walletDataObj, haveWallet := t.walletValidator.IsHaveWallet(userID)
	walletEnabled := t.walletValidator.IsWalletEnabled(walletDataObj)

	if !haveWallet {
		return models.Wallet{}, "", errors.New("Wallet Disabled")
	}

	if !walletEnabled {
		return models.Wallet{}, "", errors.New("Wallet Disabled")
	}
	return walletDataObj, userID, nil
}

func (t *TransactionController) Deposit(token string,
	transactionReq models.TransactionRequest) (map[string]interface{}, error) {
	walletDataObj, userID, err := t.validate(token)
	if err != nil {
		return make(map[string]interface{}), err
	}

	transactionData, err := t.transactionService.Deposit(walletDataObj.ID, userID, transactionReq)
	if err != nil {
		return make(map[string]interface{}), err
	}
	return transactionData, nil
}

func (t *TransactionController) Withdraw(token string,
	transactionReq models.TransactionRequest) (map[string]interface{}, error) {
	walletDataObj, userID, err := t.validate(token)
	if err != nil {
		return make(map[string]interface{}), err
	}

	transactionData, err := t.transactionService.Withdraw(walletDataObj.ID, userID, transactionReq)
	if err != nil {
		return make(map[string]interface{}), err
	}
	return transactionData, nil
}

func (t *TransactionController) FetchAll(token string) (map[string]interface{}, error) {
	walletDataObj, _, err := t.validate(token)
	if err != nil {
		return make(map[string]interface{}), err
	}

	transactionsData, err := t.transactionService.FetchAll(walletDataObj.ID)
	if err != nil {
		return make(map[string]interface{}), err
	}
	return transactionsData, nil

}
