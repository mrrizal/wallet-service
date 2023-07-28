package repositories

import (
	"context"
	"fmt"
	"mrrizal/wallet-service/app/database"
	"mrrizal/wallet-service/app/models"
	"time"
)

type WalletRepository interface {
	Enable(wallet models.Wallet) error
	Update(wallet *models.Wallet) error
	GetWallet(userID string) models.Wallet
}

type walletRepository struct {
	db database.DB
}

func NewWalletRepository(db database.DB) walletRepository {
	return walletRepository{db: db}
}

func (w *walletRepository) Enable(wallet models.Wallet) error {
	sqlStmt := `INSERT INTO wallet (id, user_id, status, enabled_at, balance) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	row := w.db.QueryRow(context.Background(), sqlStmt, wallet.ID, wallet.UserID, wallet.Status,
		wallet.EnabledAt, wallet.Balance)

	err := row.Scan(&wallet.ID)
	return err
}

func (w *walletRepository) Update(wallet *models.Wallet) error {
	var row database.Row
	wallet.EnabledAt = time.Now()
	if wallet.Status == models.WalletStatusEnabled {
		sqlStmt := `UPDATE wallet set status = $1, enabled_at = $2 WHERE id = $3 RETURNING status`
		row = w.db.QueryRow(context.Background(), sqlStmt, wallet.Status, wallet.EnabledAt, wallet.ID)
	} else {
		sqlStmt := `UPDATE wallet set status = $1 WHERE id = $2 RETURNING status`
		row = w.db.QueryRow(context.Background(), sqlStmt, wallet.Status, wallet.ID)
	}

	err := row.Scan(&wallet.Status)
	return err
}

func (w *walletRepository) GetWallet(userID string) models.Wallet {
	var wallet models.Wallet

	sqlStmt := fmt.Sprintf("SELECT id, user_id, status, enabled_at, balance FROM wallet WHERE user_id = '%s'", userID)

	err := w.db.QueryRow(context.Background(), sqlStmt).Scan(&wallet.ID, &wallet.UserID, &wallet.Status,
		&wallet.EnabledAt, &wallet.Balance)

	if err != nil {
		return models.Wallet{}
	}

	return wallet
}
