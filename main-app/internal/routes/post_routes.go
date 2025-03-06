package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/revandpratama/reflect/internal/handlers"
	"github.com/revandpratama/reflect/internal/middleware"
	"github.com/revandpratama/reflect/internal/repositories"
	"github.com/revandpratama/reflect/internal/services"
	"gorm.io/gorm"
)

func InitPostHandler(db *gorm.DB, minioClient *minio.Client) handlers.PostHandler {
	repo := repositories.NewPostRepository(db)
	service := services.NewPostService(repo, minioClient)
	return handlers.NewPostHandler(service)
}

func InitPostRoutes(r fiber.Router, handler handlers.PostHandler) {
	post := r.Group("/posts")
	post.Use(middleware.AuthMiddleware())

	post.Get("/", handler.GetAllPosts)
	post.Get("/:id", handler.GetPostByID)
	post.Post("/", handler.CreatePost)
	post.Put("/:id", handler.UpdatePost)
	post.Delete("/:id", handler.DeletePost)
}
