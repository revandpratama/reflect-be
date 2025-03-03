package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/errorhandler"
	"github.com/revandpratama/reflect/helper/response"
	"github.com/revandpratama/reflect/internal/dto"
	pb "github.com/revandpratama/reflect/internal/generatedProtobuf/auth"
	"github.com/revandpratama/reflect/types"
	"google.golang.org/grpc"
)

type authHandler struct {
	ctx    context.Context
	client pb.AuthServiceClient
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

func NewAuthHandler(conn *grpc.ClientConn) AuthHandler {
	return &authHandler{
		ctx:    context.Background(),
		client: pb.NewAuthServiceClient(conn),
	}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	res, err := h.client.Login(h.ctx, &pb.LoginRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return errorhandler.BuildError(c, &types.InternalServerError{Message: err.Error()})
	}

	response := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "login success",
		Data:       dto.LoginResponse{Token: res.AccessToken},
	})
	c.JSON(response)

	return nil
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	res, err := h.client.Register(h.ctx, &pb.RegisterRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return errorhandler.BuildError(c, &types.InternalServerError{Message: err.Error()})
	}
	response := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    res.Message,
	})
	return c.JSON(response)
}
