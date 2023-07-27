package models

import (
	"log"
	"mrrizal/wallet-service/utils"
)

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
		"enabled_at": wallet.EnabledAt.Format("2006-01-02T15:04:05-07:00"),
		"balance":    wallet.Balance,
	}

	result := make(map[string]interface{})
	result["wallet"] = walletData
	return result
}
