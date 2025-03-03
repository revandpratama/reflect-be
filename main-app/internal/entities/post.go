package entities

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int     `gorm:"primaryKey;datatype:serial"`
	UserID    int     `gorm:"column:user_id"`
	Title     string  `gorm:"column:title"`
	Body      string  `gorm:"column:body"`
	ImageUrl  *string `gorm:"column:image_url"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (Post) TableName() string {
	return "public.posts"
}
