package httpdto

import "ewallet-server-v2/internal/dto/authdto"

type RegisterResponse struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	WalletId     int64  `json:"wallet_id"`
	WalletNumber string `json:"wallet_number"`
}

func ConvertToRegisterResponse(reg *authdto.RegisterDto) RegisterResponse {
	return RegisterResponse{
		Username:     reg.User.Name,
		Email:        reg.User.Email,
		FullName:     reg.User.FullName,
		WalletId:     reg.Wallet.WalletId,
		WalletNumber: reg.Wallet.WalletNumber,
	}
}
