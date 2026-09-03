package dto

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type LoginResponseDto struct {
	Token      string `json:"token"`
	ExpireDate string `json:"expire_date"`
}

type RegisterRequestDto struct {
	FullName string `json:"full_name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,email"`
}
