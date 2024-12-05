package dtos

type RequestValidationMiddleware struct {
	Email          string  `json:"email" binding:"required"`
	Password       string  `json:"password" binding:"required"`
	NewPassword    string  `json:"new_password" binding:"required"`
	FullName       string  `json:"fullname" binding:"required"`
	BirthDate      *string `json:"birthdate" binding:"required"`
	Amount         *string `json:"amount" binding:"required"`
	SourceOfFund   *string `json:"source_of_fund_id" binding:"required"`
	ToWalletNumber *string `json:"to_wallet_number" binding:"required"`
	Description    *string `json:"description" binding:"required"`
}
