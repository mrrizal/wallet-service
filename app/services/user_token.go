package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/repositories"
)

type UserTokenService interface {
	Init(customerXID string) (string, error)
	ValidateToken(token string) (string, error)
}

type userTokenService struct {
	userTokenRepository repositories.UserTokenRepository
}

func NewUserTokenService(userTokenRepository repositories.UserTokenRepository) userTokenService {
	return userTokenService{userTokenRepository: userTokenRepository}
}

func generateRandomToken(tokenLength int) (string, error) {
	tokenLength = 42 / 2
	tokenBytes := make([]byte, tokenLength)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	token := hex.EncodeToString(tokenBytes)

	return token, nil
}

func (u *userTokenService) Init(customerXID string) (string, error) {
	token, err := generateRandomToken(42)
	if err != nil {
		return "", err
	}

	if u.userTokenRepository.IsExists("user_id", customerXID) {
		errMsg := models.UserTokenErrorMessageResponse("customer_xid", "User already have token.")
		return "", errors.New(errMsg)
	}

	tokenIsExists := false
	for i := 0; i < 5; i++ {
		if u.userTokenRepository.IsExists("token", token) {
			token, err = generateRandomToken(42)
			if err != nil {
				return "", err
			}
			tokenIsExists = true
			continue
		}
		tokenIsExists = false
		break
	}

	if tokenIsExists {
		errMsg := models.UserTokenErrorMessageResponse("token", "Failed to create token.")
		return "", errors.New(errMsg)
	}

	err = u.userTokenRepository.Create(models.UserTokenRequest{
		UserID: customerXID,
		Token:  token,
	})
	return token, err
}

func (u *userTokenService) ValidateToken(token string) (string, error) {
	return u.userTokenRepository.IsTokenValid(token)
}
