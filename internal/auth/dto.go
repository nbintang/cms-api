package auth

type RegisterRequestDTO struct {
	Name      string `json:"name" validate:"required,min=6,max=191"`
	Email     string `json:"email" validate:"required,email"`
	AvatarURL string `json:"avatar_url" validate:"omitempty,url"`
	Password  string `json:"password" validate:"required,min=6"`
}

type VerifyRequestDTO struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"min=6,max=6"`
}
