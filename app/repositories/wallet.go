package repositories

import (
	"context"
	"mrrizal/wallet-service/app/database"
	"mrrizal/wallet-service/app/models"
)

type WalletRepository interface {
	Enable(wallet models.Wallet) error
}

type walletRepository struct {
	db database.DB
}

func NewWalletRepository(db database.DB) walletRepository {
	return walletRepository{db: db}
}

func (u *walletRepository) Enable(wallet models.Wallet) error {
	sqlStmt := `INSERT INTO wallet (id, user_id, status, enabled_at, balance) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	row := u.db.QueryRow(context.Background(), sqlStmt, wallet.ID, wallet.UserID, wallet.Status,
		wallet.EnabledAt, wallet.Balance)

	err := row.Scan(&wallet.ID)
	return err
}
