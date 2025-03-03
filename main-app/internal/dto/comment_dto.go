package dto

import "time"

type CommentRequst struct {
	PostID    int       `json:"post_id" validate:"required"`
	Body      string    `json:"body" validate:"required,min=1,max=128"`
}

type CommentResponse struct {
	ID        int       `json:"id" validate:"required"`
	PostID    int       `json:"post_id" validate:"required"`
	Body      string    `json:"body" validate:"required,min=1,max=128"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
