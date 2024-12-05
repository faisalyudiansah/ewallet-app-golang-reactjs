package httpdto

type ConfirmResetPasswordRequest struct {
	Code     string `json:"code" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
