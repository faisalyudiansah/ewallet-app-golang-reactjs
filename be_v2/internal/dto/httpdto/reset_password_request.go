package httpdto

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
