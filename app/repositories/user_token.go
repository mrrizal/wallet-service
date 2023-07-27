package repositories

import (
	"context"
	"errors"
	"fmt"
	"mrrizal/wallet-service/app/database"
	"mrrizal/wallet-service/app/models"
)

type UserTokenRepository interface {
	IsExists(field, value string) bool
	Create(request models.UserTokenRequest) error
	IsTokenValid(token string) (string, error)
}

type userTokenRepository struct {
	db database.DB
}

func NewUserTokenRepository(db database.DB) userTokenRepository {
	return userTokenRepository{db: db}
}

func (u *userTokenRepository) IsExists(field, value string) bool {
	var userTokenID int

	sqlStmt := fmt.Sprintf("SELECT id FROM user_token WHERE %s = '%s'", field, value)
	err := u.db.QueryRow(context.Background(), sqlStmt).Scan(&userTokenID)
	if err != nil {
		return false
	}
	return userTokenID != 0
}

func (u *userTokenRepository) Create(request models.UserTokenRequest) error {
	sqlStmt := `INSERT INTO user_token (user_id, token) VALUES ($1, $2) RETURNING ID`

	var userTokenID int
	row := u.db.QueryRow(context.Background(), sqlStmt, request.UserID, request.Token)
	err := row.Scan(&userTokenID)
	return err
}

// IsTokenValid will return user id if valid, else return empty string
func (u *userTokenRepository) IsTokenValid(token string) (string, error) {
	var userID string

	sqlStmt := fmt.Sprintf("SELECT user_id FROM user_token WHERE %s = '%s'", "token", token)
	err := u.db.QueryRow(context.Background(), sqlStmt).Scan(&userID)
	if err != nil {
		fmt.Println("hehe")
		return "", errors.New("Invalid token.")
	}

	if userID == "" {
		return "", errors.New("Invalid token.")
	}

	return userID, nil
}
