package repositories

import (
	"context"

	"github.com/revandpratama/reflect/internal/entities"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *entities.Comment) error
	GetCommentByID(ctx context.Context, id int) (*entities.Comment, error)
	GetAllComments(ctx context.Context) ([]entities.Comment, error)
	GetCommentByPostID(ctx context.Context, postID int) ([]entities.Comment, error)
	UpdateComment(ctx context.Context, id int, comment *entities.Comment) error
	DeleteComment(ctx context.Context, id int) error
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) CreateComment(ctx context.Context, comment *entities.Comment) error {
	err := c.db.WithContext(ctx).Create(&comment).Error

	return err
}

func (c *commentRepository) GetCommentByID(ctx context.Context, id int) (*entities.Comment, error) {
	var comment entities.Comment

	err := c.db.WithContext(ctx).Where("id = ?", id).Take(&comment).Error

	return &comment, err
}

func (c *commentRepository) GetAllComments(ctx context.Context) ([]entities.Comment, error) {
	var comments []entities.Comment

	err := c.db.WithContext(ctx).Find(&comments).Error

	return comments, err
}

func (c *commentRepository) GetCommentByPostID(ctx context.Context, postID int) ([]entities.Comment, error) {
	var comments []entities.Comment

	err := c.db.WithContext(ctx).Where("post_id = ?", postID).Find(&comments).Error

	return comments, err
}

func (c *commentRepository) UpdateComment(ctx context.Context, id int, comment *entities.Comment) error {
	err := c.db.WithContext(ctx).Where("id = ?", id).Updates(&comment).Error

	return err
}

func (c *commentRepository) DeleteComment(ctx context.Context, id int) error {
	err := c.db.WithContext(ctx).Delete(&entities.Comment{}, "id = ?", id).Error

	return err
}
