package services

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/lautarooyuela/recetapp-backend/models"
)

func Decode64(account models.Account, encodedString map[string]string) models.Account {

	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString["token"])
	if err != nil {
		fmt.Println("Error:", err)
	}
	decodedString := string(decodedBytes)
	fmt.Println(decodedString)
	parts := strings.Split(decodedString, ":")
	account.Email = parts[0]
	account.Password = parts[1]

	return account
}
