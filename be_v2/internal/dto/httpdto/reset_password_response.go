package httpdto

import (
	"time"

	"ewallet-server-v2/internal/model"
)

type ResetPasswordResponse struct {
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	ExpiredAt time.Time `json:"expired_at"`
}

func ConvertToResetPasswordResponse(email string, resetPasswordAttempt *model.ResetPasswordAttempt) ResetPasswordResponse {
	return ResetPasswordResponse{
		Email:     email,
		Code:      resetPasswordAttempt.Code,
		ExpiredAt: resetPasswordAttempt.ExpiredAt,
	}
}
