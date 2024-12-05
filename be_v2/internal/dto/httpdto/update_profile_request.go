package httpdto

type UpdateProfileRequest struct {
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	ProfileImage string `json:"profile_image"`
}
