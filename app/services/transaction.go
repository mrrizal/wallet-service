package services

import (
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/repositories"
	"time"

	"github.com/google/uuid"
)

type TransactionService interface {
	Deposit(walletID, userID string, transactionReq models.TransactionRequest) (map[string]interface{}, error)
	Withdraw(walletID, userID string, transactionReq models.TransactionRequest) (map[string]interface{}, error)
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
}

func NewTransactionService(transactionRepository repositories.TransactionRepository) transactionService {
	return transactionService{transactionRepository: transactionRepository}
}

func (t *transactionService) Deposit(walletID, userID string,
	transactionReq models.TransactionRequest) (map[string]interface{}, error) {
	transactionData := models.Transaction{
		ID:           uuid.NewString(),
		WalletID:     walletID,
		TransactedAt: time.Now(),
		Type:         models.TransactionTypeDeposit,
		Amount:       transactionReq.Amount,
		ReferenceID:  transactionReq.ReferenceID,
	}

	if err := t.transactionRepository.Deposit(&transactionData); err != nil {
		return make(map[string]interface{}), err
	}
	return models.ParseDeposit(userID, transactionData), nil
}

func (t *transactionService) Withdraw(walletID, userID string,
	transactionReq models.TransactionRequest) (map[string]interface{}, error) {
	transactionData := models.Transaction{
		ID:           uuid.NewString(),
		WalletID:     walletID,
		TransactedAt: time.Now(),
		Type:         models.TransactionTypeWithdraw,
		Amount:       transactionReq.Amount,
		ReferenceID:  transactionReq.ReferenceID,
	}

	if err := t.transactionRepository.Withdraw(&transactionData); err != nil {
		return make(map[string]interface{}), err
	}
	return models.ParseWithdraw(userID, transactionData), nil
}
