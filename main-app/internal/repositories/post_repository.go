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
	GetAllPosts(ctx context.Context) ([]entities.Post, error)
	GetPostByUserID(ctx context.Context, userID int) ([]entities.Post, error)
	UpdatePost(ctx context.Context, post *entities.Post) error
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

func (p *postRepository) GetAllPosts(ctx context.Context) ([]entities.Post, error) {
	var posts []entities.Post

	err := p.db.WithContext(ctx).Find(&posts).Error

	return posts, err
}

func (p *postRepository) GetPostByUserID(ctx context.Context, userID int) ([]entities.Post, error) {
	var posts []entities.Post

	err := p.db.WithContext(ctx).Where("user_id = ?", userID).Find(&posts).Error

	return posts, err
}

// TODO: Update post still need to be fixed
func (p *postRepository) UpdatePost(ctx context.Context, post *entities.Post) error {
	// ? use map to update values
	err := p.db.WithContext(ctx).Where("id = ?", post.ID).Updates(&post).Error

	return err
}

func (p *postRepository) DeletePost(ctx context.Context, id int) error {
	err := p.db.WithContext(ctx).Where("id = ?", id).Delete(&entities.Post{}).Error

	return err
}
