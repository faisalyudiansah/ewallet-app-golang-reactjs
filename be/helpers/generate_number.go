package helpers

import (
	"fmt"
	"time"
)

type GenerateNumberInterface interface {
	GenerateTokenResetPassword(int64) string
}

type GenerateNumber struct{}

func NewGenerateNumber() *GenerateNumber {
	return &GenerateNumber{}
}

func (gn *GenerateNumber) GenerateTokenResetPassword(userID int64) string {
	currentTime := time.Now()
	dateFormat := currentTime.Format("20060102150405")
	walletNumber := fmt.Sprintf("%s%d", dateFormat, userID)
	return walletNumber
}
