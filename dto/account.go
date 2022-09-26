package dto

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLoginDTO struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"accessToken"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
