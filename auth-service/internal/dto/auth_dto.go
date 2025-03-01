package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	RoleID   int    `json:"role_id" validate:"required"`
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=128"`
	Email    string `json:"email" validate:"required,email"`
}
