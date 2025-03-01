package controller

import (
	"context"

	pb "github.com/revandpratama/reflect/auth-service/internal/generatedProtobuf/auth"
	"github.com/revandpratama/reflect/auth-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authController struct {
	pb.UnimplementedAuthServiceServer
	service services.AuthService
}

func NewAuthController(service services.AuthService) pb.AuthServiceServer {
	return &authController{
		service: service,
	}
}

func (c *authController) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := c.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials: %v", err)
	}

	return &pb.LoginResponse{AccessToken: token}, nil
}
