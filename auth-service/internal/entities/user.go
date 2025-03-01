package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey;datatype:serial"`
	RoleID    int    `gorm:"column:role_id"`
	Username  string `gorm:"unique"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Role struct {
	ID   int `gorm:"datatype:serial;primaryKey"`
	Name string
}

func (User) TableName() string {
	return "authentication.users"
}
