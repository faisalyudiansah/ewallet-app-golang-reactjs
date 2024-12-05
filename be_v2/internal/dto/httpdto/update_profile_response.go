package httpdto

import "ewallet-server-v2/internal/model"

type UpdateProfileResponse struct {
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	ProfileImage string `json:"profile_image"`
}

func ConvertToUpdateProfileResponse(user *model.User) UpdateProfileResponse {
	return UpdateProfileResponse{
		Email:        user.Email,
		FullName:     user.FullName,
		ProfileImage: user.ProfileImage,
	}
}
