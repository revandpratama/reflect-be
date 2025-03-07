package services

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/revandpratama/reflect/helper"
	"github.com/revandpratama/reflect/internal/dto"
	"github.com/revandpratama/reflect/internal/entities"
	"github.com/revandpratama/reflect/internal/repositories"
	"github.com/revandpratama/reflect/types"
	"gorm.io/gorm"
)

type postService struct {
	repo        repositories.PostRepository
	commentRepo repositories.CommentRepository
	minioClient *minio.Client
}

type PostService interface {
	GetAllPosts(ctx context.Context, page, limit int) ([]dto.PostResponse, *types.Pagination, error)
	GetPostByID(ctx context.Context, id int) (*dto.PostResponse, error)
	GetPostByUserID(ctx context.Context, userID int) ([]dto.PostResponse, error)
	CreatePost(ctx context.Context, req *dto.PostRequest) error
	UpdatePost(ctx context.Context, id int, req *dto.PostRequest) error
	DeletePost(ctx context.Context, id int) error
}

func NewPostService(repo repositories.PostRepository, commentRepo repositories.CommentRepository, minioCLient *minio.Client) PostService {
	return &postService{
		repo:        repo,
		commentRepo: commentRepo,
		minioClient: minioCLient,
	}
}

func (p *postService) CreatePost(ctx context.Context, req *dto.PostRequest) error {

	// userID, err := strconv.Atoi(req.UserID)
	// if err != nil {
	// 	return &types.BadRequestError{Message: "invalid user id"}
	// }

	post := entities.Post{
		UserID: req.UserID,
		Title:  req.Title,
		Body:   req.Body,
	}

	if req.Image != nil {
		imgUrl, err := helper.UploadObject(ctx, p.minioClient, helper.MINIO_POST_BUCKET, req.Image)
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

	comments, err := p.commentRepo.GetCommentByPostID(ctx, post.ID)
	if err != nil {
		return nil, &types.InternalServerError{Message: err.Error()}
	}

	commentResponses := make([]dto.CommentResponse, len(comments))
	for i := range comments {
		commentResponses[i] = comments[i].ToResponse()
	}

	postResponse.Comments = commentResponses

	return &postResponse, nil
}

func (p *postService) GetAllPosts(ctx context.Context, page, limit int) ([]dto.PostResponse, *types.Pagination, error) {

	post, err := p.repo.GetAllPosts(ctx, page, limit)
	if err != nil {
		return nil, nil, &types.InternalServerError{Message: err.Error()}
	}

	if len(post) == 0 {
		return nil, nil, &types.NotFoundError{Message: "post not found"}
	}

	totalPage, err := p.repo.GetTotalPage(ctx, limit)
	if page > totalPage {
		return nil, nil, &types.NotFoundError{Message: "page not found"}
	}
	if err != nil {
		return nil, nil, &types.InternalServerError{Message: err.Error()}
	}

	pagination := &types.Pagination{
		Page:      page,
		Limit:     limit,
		TotalPage: totalPage,
	}

	postResponses := make([]dto.PostResponse, len(post))
	for i, p := range post {
		postResponses[i] = p.ToResponse()
	}

	return postResponses, pagination, nil

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

func (p *postService) UpdatePost(ctx context.Context, id int, req *dto.PostRequest) error {
	post := &entities.Post{
		UserID: req.UserID,
		Title:  req.Title,
		Body:   req.Body,
	}

	if req.Image != nil {
		newUrl, err := p.updateImageUrl(ctx, p.minioClient, id, helper.MINIO_POST_BUCKET, req.Image)
		if err != nil {
			return &types.InternalServerError{Message: err.Error()}
		}

		post.ImageUrl = &newUrl
	}

	if err := p.repo.UpdatePost(ctx, id, post); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &types.NotFoundError{Message: err.Error()}
		}
		return &types.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (p *postService) DeletePost(ctx context.Context, id int) error {

	if err := p.deleteImageUrl(ctx, p.minioClient, id, helper.MINIO_POST_BUCKET); err != nil {
		return &types.InternalServerError{Message: err.Error()}
	}

	if err := p.repo.DeletePost(ctx, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &types.NotFoundError{Message: err.Error()}
		}
		return &types.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *postService) updateImageUrl(ctx context.Context, minioClient *minio.Client, postID int, bucketName string, image *multipart.FileHeader) (string, error) {
	post, err := s.repo.GetPostByID(ctx, postID)
	if err != nil {
		return "", err
	}

	newUrl, err := helper.UploadObject(ctx, minioClient, bucketName, image)
	if err != nil {
		return "", err
	}

	if post.ImageUrl != nil {
		err = helper.DeleteObject(ctx, minioClient, bucketName, *post.ImageUrl)
		if err != nil {
			return "", err
		}
	}

	return newUrl, nil
}

func (s *postService) deleteImageUrl(ctx context.Context, minioClient *minio.Client, postID int, bucketName string) error {
	post, err := s.repo.GetPostByID(ctx, postID)
	if err != nil {
		return err
	}

	if post.ImageUrl != nil {
		err = helper.DeleteObject(ctx, minioClient, bucketName, *post.ImageUrl)
		if err != nil {
			return err
		}
	}

	return nil
}
