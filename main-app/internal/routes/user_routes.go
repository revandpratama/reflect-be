package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/internal/handlers"
	"github.com/revandpratama/reflect/internal/middleware"
	"github.com/revandpratama/reflect/internal/repositories"
	"github.com/revandpratama/reflect/internal/services"
	"gorm.io/gorm"
)

func InitUserHandler(db *gorm.DB) handlers.UserHandler {
	repo := repositories.NewUserRepository(db)
	services := services.NewUserService(repo)
	return handlers.NewUserHandler(services)
}

func InitUserRoutes(r fiber.Router, userHandler handlers.UserHandler) {
	user := r.Group("/users")
	user.Use(middleware.AuthMiddleware())

	user.Get("/:id", userHandler.GetUserByID)
}
