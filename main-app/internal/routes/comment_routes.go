package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/internal/handlers"
	"github.com/revandpratama/reflect/internal/middleware"
	"github.com/revandpratama/reflect/internal/repositories"
	"github.com/revandpratama/reflect/internal/services"
	"gorm.io/gorm"
)

func InitCommentHandler(db *gorm.DB) handlers.CommentHandler {
	repo := repositories.NewCommentRepository(db)
	service := services.NewCommentService(repo)
	return handlers.NewCommentHandler(service)
}

func InitCommentRoutes(r fiber.Router, handler handlers.CommentHandler) {
	comment := r.Group("/comments")
	comment.Use(middleware.AuthMiddleware())

	comment.Get("/posts/:postid", handler.GetCommentByPostID)
	comment.Get("/:id", handler.GetCommentByID)
	comment.Post("/", handler.CreateComment)
	comment.Put("/:id", handler.UpdateComment)
	comment.Delete("/:id", handler.DeleteComment)
}
