package controllers

import (
	"mrrizal/wallet-service/app/service"
	"mrrizal/wallet-service/app/validators"
)

type UserTokenController struct {
	userTokenValidator validators.UserTokenValidator
	userTokenService   service.UserTokenService
}

func NewUserTokenController(
	userTokenValidator validators.UserTokenValidator,
	userTokenService service.UserTokenService) UserTokenController {
	return UserTokenController{
		userTokenValidator: userTokenValidator,
		userTokenService:   userTokenService,
	}
}

func (u *UserTokenController) Init(customerXID string) (string, error) {
	if err := u.userTokenValidator.ValidateUserTokenRequest(customerXID); err != nil {
		return "", err
	}
	return u.userTokenService.Init(customerXID)
}
