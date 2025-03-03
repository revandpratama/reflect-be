package dto

import (
	"mime/multipart"
	"time"
)

type PostRequest struct {
	UserID int    `json:"user_id" validate:"required"`
	Title  string `json:"title" validate:"required,min=3,max=32"`
	Body   string `json:"body" validate:"required,min=3,max=128"`
	Image  *multipart.FileHeader
}

type PostResponse struct {
	ID        int       `json:"id" validate:"required"`
	UserID    int       `json:"user_id" validate:"required"`
	Title     string    `json:"title" validate:"required,min=3,max=32"`
	Body      string    `json:"body" validate:"required,min=3,max=128"`
	ImageUrl  *string   `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}
