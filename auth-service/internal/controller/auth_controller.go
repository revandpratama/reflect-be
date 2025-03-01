package controller

import (
	"context"

	"github.com/revandpratama/reflect/auth-service/internal/dto"
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

func (c *authController) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	request := dto.RegisterRequest{
		RoleID:   int(req.RoleId),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := c.service.Register(ctx, request); err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create new user: %v", err)
	}

	return &pb.RegisterResponse{
		Message: "user created successfully",
	}, nil
}
