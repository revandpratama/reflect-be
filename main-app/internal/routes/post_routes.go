package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/internal/handlers"
	"github.com/revandpratama/reflect/internal/repositories"
	"github.com/revandpratama/reflect/internal/services"
	"gorm.io/gorm"
)

func InitPostHandler(db *gorm.DB) handlers.PostHandler {
	repo := repositories.NewPostRepository(db)
	service := services.NewPostService(repo)
	return handlers.NewPostHandler(service)
}

func InitPostRoutes(r fiber.Router, handler handlers.PostHandler) {
	r.Get("/posts", handler.GetAllPosts)
	r.Get("/posts/:id", handler.GetPostByID)
	r.Post("/posts", handler.CreatePost)
	r.Put("/posts/:id", handler.UpdatePost)
	r.Delete("/posts/:id", handler.DeletePost)
}


