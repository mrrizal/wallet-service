package models

import "time"

type WalletStatus string

const (
	WalletStatusEnabled  WalletStatus = "enabled"
	WalletStatusDisabled WalletStatus = "disabled"
)

type Wallet struct {
	ID        string       `json:"id"`
	UserID    string       `json:"user_id"`
	Status    WalletStatus `json:"status"`
	EnabledAt time.Time    `json:"enabled_at"`
	Balance   float64      `json:"balance"`
}
