package validators

import (
	"errors"
	"mrrizal/wallet-service/app/models"
	"mrrizal/wallet-service/app/services"
	"regexp"
)

type UserTokenValidator interface {
	ValidateUserTokenRequest(customerXID string) error
	ValidateToken(token string) (string, error)
}

type userTokenValidator struct {
	userTokenService services.UserTokenService
}

func NewUserTokenValidator(userTokenService services.UserTokenService) userTokenValidator {
	return userTokenValidator{userTokenService: userTokenService}
}

func isValidUUID(input string) bool {
	uuidPattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	match, err := regexp.MatchString(uuidPattern, input)
	if err != nil {
		return false
	}
	return match
}

func (u *userTokenValidator) ValidateUserTokenRequest(customerXID string) error {
	switch {
	case customerXID == "":
		msg := models.UserTokenErrorMessageResponse("customer_xid", "Missing data for required field.")
		return errors.New(msg)
	case !isValidUUID(customerXID):
		msg := models.UserTokenErrorMessageResponse("customer_xid", "Invalid input.")
		return errors.New(msg)
	}
	return nil
}

func (u *userTokenValidator) ValidateToken(token string) (string, error) {
	return u.userTokenService.ValidateToken(token)
}
