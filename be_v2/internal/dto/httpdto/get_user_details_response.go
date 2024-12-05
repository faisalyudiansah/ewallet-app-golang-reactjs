package httpdto

import (
	"ewallet-server-v2/internal/model"

	"github.com/shopspring/decimal"
)

type GetUserDetailsResponse struct {
	UserId       int64           `json:"user_id"`
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	WalletId     int64           `json:"wallet_id"`
	WalletNumber string          `json:"wallet_number"`
	ProfileImage string          `json:"profile_image"`
	Amount       decimal.Decimal `json:"amount"`
}

func ConvertToGetUserDetailResponse(user *model.User, wallet *model.Wallet) GetUserDetailsResponse {
	return GetUserDetailsResponse{
		UserId:       user.UserId,
		Name:         user.Name,
		Email:        user.Email,
		WalletId:     wallet.WalletId,
		WalletNumber: wallet.WalletNumber,
		ProfileImage: user.ProfileImage,
		Amount:       wallet.Amount,
	}
}
