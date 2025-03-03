package entities

import (
	"time"

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