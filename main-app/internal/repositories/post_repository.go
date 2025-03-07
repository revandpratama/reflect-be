package repositories

import (
	"context"

	"github.com/revandpratama/reflect/internal/entities"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

type PostRepository interface {
	CreatePost(ctx context.Context, post *entities.Post) error
	GetPostByID(ctx context.Context, id int) (*entities.Post, error)
	GetAllPosts(ctx context.Context, id, limit int) ([]entities.Post, error)
	GetTotalPage(ctx context.Context, limit int) (int, error)
	GetPostByUserID(ctx context.Context, userID int) ([]entities.Post, error)
	UpdatePost(ctx context.Context, id int, post *entities.Post) error
	DeletePost(ctx context.Context, id int) error
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) CreatePost(ctx context.Context, post *entities.Post) error {
	err := p.db.WithContext(ctx).Create(&post).Error

	return err
}

func (p *postRepository) GetPostByID(ctx context.Context, id int) (*entities.Post, error) {
	var post entities.Post

	err := p.db.WithContext(ctx).Where("id = ?", id).Take(&post).Error

	return &post, err
}

func (p *postRepository) GetAllPosts(ctx context.Context, page, limit int) ([]entities.Post, error) {
	var posts []entities.Post

	err := p.db.WithContext(ctx).Limit(limit).Offset((page - 1) * limit).Find(&posts).Error

	return posts, err
}
func (p *postRepository) GetTotalPage(ctx context.Context, limit int) (int, error) {
	var count int64

	err := p.db.WithContext(ctx).Model(&entities.Post{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	totalPages := (int(count) + limit - 1) / limit

	return totalPages, nil
}

func (p *postRepository) GetPostByUserID(ctx context.Context, userID int) ([]entities.Post, error) {
	var posts []entities.Post

	err := p.db.WithContext(ctx).Where("user_id = ?", userID).Find(&posts).Error

	return posts, err
}

func (p *postRepository) UpdatePost(ctx context.Context, id int, post *entities.Post) error {
	err := p.db.WithContext(ctx).Where("id = ?", id).Updates(&post).Error

	return err
}

func (p *postRepository) DeletePost(ctx context.Context, id int) error {
	err := p.db.WithContext(ctx).Where("id = ?", id).Delete(&entities.Post{}).Error

	return err
}
