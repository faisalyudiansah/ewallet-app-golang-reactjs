package dtos

type RequestRegisterUser struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=5"`
	FullName  string `json:"fullname" binding:"required"`
	BirthDate string `json:"birthdate" binding:"required"`
}

type RequestLoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type RequestForgetPassword struct {
	Email string `json:"email" binding:"required,email"`
}

type RequestResetPassword struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=5"`
}

type RequestTopUpWallet struct {
	Amount       float64 `json:"amount" binding:"required,gte=50000,lte=10000000"`
	SourceOfFund int64   `json:"source_of_fund_id" binding:"required"`
}

type RequestChooseBox struct {
	BoxIndex int `json:"box_index" binding:"required"`
}

type RequestTransferFund struct {
	ToWalletNumber string  `json:"to_wallet_number" binding:"required"`
	Amount         float64 `json:"amount" binding:"required,gte=1000"`
	Description    string  `json:"description" binding:"required,max=35"`
}
