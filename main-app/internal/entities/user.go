package entities

import (
	"time"

	"github.com/revandpratama/reflect/internal/dto"
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey;datatype:serial"`
	RoleID    int `gorm:"column:role_id"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (User) TableName() string {
	return "authentication.users"
}

func (u *User) ToResponse() dto.UserResponse {
	return dto.UserResponse{
		ID:     u.ID,
		RoleID: u.RoleID,
		Name:   u.Name,
		Email:  u.Email,
	}
}
