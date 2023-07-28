package repositories

import (
	"context"
	"fmt"
	"mrrizal/wallet-service/app/database"
	"mrrizal/wallet-service/app/models"
)

type TransactionRepository interface {
	Deposit(transaction *models.Transaction) error
	Withdraw(transaction *models.Transaction) error
	FetchAll(walletID string) ([]models.Transaction, error)
}

type transactionRepository struct {
	db database.DB
}

func NewTransactionRepository(db database.DB) transactionRepository {
	return transactionRepository{db: db}
}

func (t *transactionRepository) Deposit(transaction *models.Transaction) error {
	ctx := context.Background()
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return err
	}

	// get current balance
	var balance float64
	sqlStmt := fmt.Sprintf("SELECT balance FROM wallet WHERE id = '%s'", transaction.WalletID)
	rows, err := tx.Query(ctx, sqlStmt)
	defer rows.Close()

	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&balance)
	}

	// add balance with deposit amount
	_, err = tx.Exec(ctx, `UPDATE wallet SET balance = $1 WHERE id = $2`,
		balance+transaction.Amount, transaction.WalletID)
	if err != nil {
		return err
	}

	transaction.Status = models.TransactionStatusSuccess
	// insert transaction data
	_, err = tx.Exec(ctx, `INSERT INTO transactions (id, wallet_id, status, transacted_at, type, amount, reference_id)
	VALUES ($1, $2, $3, $4, $5 , $6 , $7)`, transaction.ID, transaction.WalletID, transaction.Status,
		transaction.TransactedAt, transaction.Type, transaction.Amount, transaction.ReferenceID,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return err
	}

	return nil
}

func (t *transactionRepository) Withdraw(transaction *models.Transaction) error {
	ctx := context.Background()
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return err
	}

	// get current balance
	var balance float64
	sqlStmt := fmt.Sprintf("SELECT balance FROM wallet WHERE id = '%s'", transaction.WalletID)
	rows, err := tx.Query(ctx, sqlStmt)
	defer rows.Close()

	if err != nil {
		return err
	}

	for rows.Next() {
		rows.Scan(&balance)
	}

	// deduct balance with deposit amount
	if balance-transaction.Amount >= 0 {
		_, err = tx.Exec(ctx, `UPDATE wallet SET balance = $1 WHERE id = $2`,
			balance-transaction.Amount, transaction.WalletID)
		if err != nil {
			return err
		}
		transaction.Status = models.TransactionStatusSuccess
	} else {
		transaction.Status = models.TransactionStatusFailed
	}

	// insert transaction data
	_, err = tx.Exec(ctx, `INSERT INTO transactions (id, wallet_id, status, transacted_at, type, amount, reference_id)
	VALUES ($1, $2, $3, $4, $5 , $6 , $7)`, transaction.ID, transaction.WalletID, transaction.Status,
		transaction.TransactedAt, transaction.Type, transaction.Amount, transaction.ReferenceID,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return err
	}

	return nil
}

func (w *transactionRepository) FetchAll(walletID string) ([]models.Transaction, error) {
	sqlStmt := `SELECT id, wallet_id, status, transacted_at, type, amount,
	reference_id FROM transactions WHERE wallet_id = $1`

	rows, err := w.db.Query(context.Background(), sqlStmt, walletID)
	defer rows.Close()

	if err != nil {
		return []models.Transaction{}, err
	}

	transactions := []models.Transaction{}
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.WalletID,
			&transaction.Status,
			&transaction.TransactedAt,
			&transaction.Type,
			&transaction.Amount,
			&transaction.ReferenceID,
		)
		if err != nil {
			return []models.Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
