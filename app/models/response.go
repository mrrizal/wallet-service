package models

import (
	"log"
	"mrrizal/wallet-service/utils"
	"time"
)

const timeFormat = "2006-01-02T15:04:05-07:00"

func Response(data map[string]interface{}, status string) map[string]interface{} {
	dataMap := map[string]interface{}{
		"data":   data,
		"status": status,
	}
	return dataMap
}

func UserTokenErrorMessageResponse(field, msg string) string {
	errMessage := map[string]interface{}{
		"error": map[string][]string{},
	}
	errMessage["error"] = map[string][]string{
		field: {},
	}

	errMessage["error"].(map[string][]string)[field] = []string{msg}
	errMsg, err := utils.MapToJson(errMessage)
	if err != nil {
		log.Fatal(err.Error())
	}
	return errMsg
}

func ErrResponse(err error) map[string]interface{} {
	tempErr := err.Error()
	errMsg, err := utils.JsonToMap(tempErr)
	if err != nil {
		temp := map[string]interface{}{
			"error": tempErr,
		}
		errMsg = temp
	}
	return Response(errMsg, "fail")
}

func ParseWallet(wallet Wallet) map[string]interface{} {
	walletData := map[string]interface{}{
		"id":         wallet.ID,
		"owned_by":   wallet.UserID,
		"status":     wallet.Status,
		"enabled_at": wallet.EnabledAt.Format(timeFormat),
		"balance":    wallet.Balance,
	}

	result := make(map[string]interface{})
	result["wallet"] = walletData
	return result
}

func ParseDeposit(userID string, transaction Transaction) map[string]interface{} {
	return ParseTransaction("deposit", userID, transaction)
}

func ParseWithdraw(userID string, transaction Transaction) map[string]interface{} {
	return ParseTransaction("withdraw", userID, transaction)
}

func ParseTransaction(key, userID string, transaction Transaction) map[string]interface{} {
	transactionData := map[string]interface{}{
		"id":           transaction.ID,
		"deposited_by": userID,
		"status":       transaction.Status,
		"deposited_at": transaction.TransactedAt.Format(timeFormat),
		"amount":       transaction.Amount,
		"reference_id": transaction.ReferenceID,
	}
	result := make(map[string]interface{})
	result[key] = transactionData
	return result
}

func MapTopWallet(walletData map[string]interface{}) Wallet {
	if walletData == nil {
		return Wallet{}
	}

	walletData = walletData["wallet"].(map[string]interface{})

	enableAt, err := time.Parse(timeFormat, walletData["enabled_at"].(string))
	if err != nil {
		return Wallet{}
	}

	wallet := Wallet{
		ID:        walletData["id"].(string),
		UserID:    walletData["owned_by"].(string),
		Status:    walletData["status"].(WalletStatus),
		EnabledAt: enableAt,
		Balance:   walletData["balance"].(float64),
	}
	return wallet
}
