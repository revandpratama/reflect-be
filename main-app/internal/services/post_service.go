package services

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/revandpratama/reflect/helper"
	"github.com/revandpratama/reflect/internal/dto"
	"github.com/revandpratama/reflect/internal/entities"
	"github.com/revandpratama/reflect/internal/repositories"
	"github.com/revandpratama/reflect/types"
	"gorm.io/gorm"
)

type postService struct {
	repo repositories.PostRepository
	minioClient *minio.Client
}

type PostService interface {
	GetAllPosts(ctx context.Context) ([]dto.PostResponse, error)
	GetPostByID(ctx context.Context, id int) (*dto.PostResponse, error)
	GetPostByUserID(ctx context.Context, userID int) ([]dto.PostResponse, error)
	CreatePost(ctx context.Context, req *dto.PostRequest) error
	UpdatePost(ctx context.Context, req *dto.PostRequest) error
	DeletePost(ctx context.Context, id int) error
}

func NewPostService(repo repositories.PostRepository, minioCLient *minio.Client) PostService {
	return &postService{
		repo: repo,
		minioClient: minioCLient,
	}
}

func (p *postService) CreatePost(ctx context.Context, req *dto.PostRequest) error {

	// TODO: implement image upload
	post := entities.Post{
		UserID:   req.UserID,
		Title:    req.Title,
		Body:     req.Body,
	}

	if req.Image != nil {
		imgUrl, err := helper.UploadObject(ctx, p.minioClient, "posts", req.Image)
		if err != nil {
			return &types.InternalServerError{Message: err.Error()}
		}
		post.ImageUrl = &imgUrl
	}

	err := p.repo.CreatePost(ctx, &post)
	if err != nil {
		return &types.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (p *postService) GetPostByID(ctx context.Context, id int) (*dto.PostResponse, error) {

	post, err := p.repo.GetPostByID(ctx, id)
	if err != nil {
		return nil, &types.NotFoundError{Message: err.Error()}
	}
	postResponse := post.ToResponse()
	return &postResponse, nil
}

func (p *postService) GetAllPosts(ctx context.Context) ([]dto.PostResponse, error) {

	post, err := p.repo.GetAllPosts(ctx)
	if err != nil {
		return nil, &types.InternalServerError{Message: err.Error()}
	}

	if len(post) == 0 {
		return nil, &types.NotFoundError{Message: "post not found"}
	}

	postResponses := make([]dto.PostResponse, len(post))
	for i, p := range post {
		postResponses[i] = p.ToResponse()
	}

	return postResponses, nil

}

func (p *postService) GetPostByUserID(ctx context.Context, userID int) ([]dto.PostResponse, error) {
	post, err := p.repo.GetPostByUserID(ctx, userID)

	if err != nil || len(post) == 0 {
		return nil, &types.NotFoundError{Message: err.Error()}
	}

	postResponses := make([]dto.PostResponse, len(post))
	for i, p := range post {
		postResponses[i] = p.ToResponse()
	}

	return postResponses, nil
}

func (p *postService) UpdatePost(ctx context.Context, req *dto.PostRequest) error {
	post := &entities.Post{
		UserID: req.UserID,
		Title:  req.Title,
		Body:   req.Body,
		// TODO : add image upload
	}
	if err := p.repo.UpdatePost(ctx, post); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &types.NotFoundError{Message: err.Error()}
		}
		return &types.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (p *postService) DeletePost(ctx context.Context, id int) error {
	if err := p.repo.DeletePost(ctx, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &types.NotFoundError{Message: err.Error()}
		}
		return &types.InternalServerError{Message: err.Error()}
	}

	return nil
}
