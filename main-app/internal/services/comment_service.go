package services

import (
	"context"

	"github.com/revandpratama/reflect/internal/dto"
	"github.com/revandpratama/reflect/internal/entities"
	"github.com/revandpratama/reflect/internal/repositories"
	"github.com/revandpratama/reflect/types"
	"gorm.io/gorm"
)

type CommentService interface {
	CreateComment(ctx context.Context, req *dto.CommentRequest) error
	GetCommentByID(ctx context.Context, id int) (*dto.CommentResponse, error)
	GetCommentByPostID(ctx context.Context, postID int) ([]dto.CommentResponse, error)
	UpdateComment(ctx context.Context, id int, req *dto.CommentRequest) error
	DeleteComment(ctx context.Context, id int) error
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{
		repo: repo,
	}
}

func (c *commentService) CreateComment(ctx context.Context, req *dto.CommentRequest) error {

	comment := entities.Comment{
		PostID: req.PostID,
		UserID: req.UserID,
		Body:   req.Body,
	}

	if err := c.repo.CreateComment(ctx, &comment); err != nil {
		return &types.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (c *commentService) GetCommentByID(ctx context.Context, id int) (*dto.CommentResponse, error) {
	comment, err := c.repo.GetCommentByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &types.NotFoundError{Message: err.Error()}
		}
		return nil, &types.InternalServerError{Message: err.Error()}
	}

	commentResponse := comment.ToResponse()

	return &commentResponse, nil
}

func (c *commentService) GetCommentByPostID(ctx context.Context, postID int) ([]dto.CommentResponse, error) {
	comments, err := c.repo.GetCommentByPostID(ctx, postID)
	if err != nil {
        return nil, &types.InternalServerError{Message: err.Error()}
    }
    if len(comments) == 0 {
        return nil, &types.NotFoundError{Message: "No comments found for this post"}
    }

	commentResponses := make([]dto.CommentResponse, len(comments))
	for i := range comments {
		commentResponses[i] = comments[i].ToResponse()
	}

	return commentResponses, nil
}

func (c *commentService) UpdateComment(ctx context.Context, id int, req *dto.CommentRequest) error {
	comment := entities.Comment{
		PostID: req.PostID,
		UserID: req.UserID,
		Body:   req.Body,
	}

	if err := c.repo.UpdateComment(ctx, id, &comment); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &types.NotFoundError{Message: err.Error()}
		}
		return &types.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (c *commentService) DeleteComment(ctx context.Context, id int) error {
	err := c.repo.DeleteComment(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &types.NotFoundError{Message: err.Error()}
		}
		return &types.InternalServerError{Message: err.Error()}
	}

	return nil
}
