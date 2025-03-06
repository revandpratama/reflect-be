package entities

import (
	"time"

	"github.com/revandpratama/reflect/internal/dto"
	"gorm.io/gorm"
)

type Comment struct {
	ID        int            `gorm:"primaryKey;datatype:serial"`
	UserID    int            `gorm:"column:user_id"`
	PostID    int            `gorm:"column:post_id"`
	Body      string         `gorm:"column:body"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Comment) TableName() string {
	return "public.comments"
}

func (c *Comment) ToResponse() dto.CommentResponse {
	return dto.CommentResponse{
		ID:        c.ID,
		PostID:    c.PostID,
		UserID:    c.UserID,
		Body:      c.Body,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
