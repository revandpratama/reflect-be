package dto

type UserResponse struct {
	ID       int    `json:"id" validate:"required"`
	RoleID   int    `json:"role_id" validate:"required"`
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type UserRequest struct {
	RoleID   int    `json:"role_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
