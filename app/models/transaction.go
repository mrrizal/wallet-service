package models

import "time"

type TransactionStatus string

const (
	TransactionStatusSuccess TransactionStatus = "success"
	TransactionStatusFailed  TransactionStatus = "failed"
)

type TransactionType string

const (
	TransactionTypeWithdraw TransactionType = "withdrawal"
	TransactionTypeDeposit  TransactionType = "deposit"
)

type Transaction struct {
	ID           string            `json:"id"`
	WalletID     string            `json:"wallet_id"`
	Status       TransactionStatus `json:"status"`
	TransactedAt time.Time         `json:"transacted_at"`
	Type         TransactionType   `json:"type"`
	Amount       float64           `json:"amount"`
	ReferenceID  string            `json:"reference_id"`
}

type TransactionRequest struct {
	Type        TransactionType
	Amount      float64
	ReferenceID string
}
