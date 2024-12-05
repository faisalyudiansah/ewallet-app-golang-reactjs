package authdto

import "ewallet-server-v2/internal/model"

type RegisterDto struct {
	User   model.User
	Wallet model.Wallet
}
