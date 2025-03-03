package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/internal/handlers"
	"google.golang.org/grpc"
)

func InitAuthHandler(conn *grpc.ClientConn) handlers.AuthHandler {
	return handlers.NewAuthHandler(conn)
}

func InitAuthRoutes(r fiber.Router, handler handlers.AuthHandler) {
	r.Post("/auth/login", handler.Login)
	r.Post("/auth/register", handler.Register)
}
