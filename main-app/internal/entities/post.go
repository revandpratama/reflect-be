package entities

import (
	"time"

	"github.com/revandpratama/reflect/helper"
	"github.com/revandpratama/reflect/internal/dto"
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

func (p *Post) ToResponse() dto.PostResponse {

	relativePathImageUrl := helper.ToRelativePath(*p.ImageUrl)

	return dto.PostResponse{
		ID:        p.ID,
		UserID:    p.UserID,
		Title:     p.Title,
		Body:      p.Body,
		ImageUrl:  &relativePathImageUrl,
		CreatedAt: p.CreatedAt,
	}
}
