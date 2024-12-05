package httpdto

import "ewallet-server-v2/internal/dto/authdto"

type LoginResponse struct {
	Token string `json:"token"`
}

func ConvertToLoginResponse(log *authdto.LoginDto) LoginResponse {
	return LoginResponse{
		Token: log.Token,
	}
}
