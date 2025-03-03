package entities

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int            `gorm:"primaryKey;datatype:serial"`
	PostID    int            `gorm:"column:post_id"`
	Body      string         `gorm:"column:body"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Comment) TableName() string {
	return "public.comments"
}
