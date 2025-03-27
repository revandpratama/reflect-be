package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/errorhandler"
	"github.com/revandpratama/reflect/helper/response"
	"github.com/revandpratama/reflect/internal/services"
	"github.com/revandpratama/reflect/types"
)

type UserHandler interface {
	GetUserByID(c *fiber.Ctx) error
}

type userHandler struct {
	service services.UserService
	ctx     context.Context
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{
		service: service,
		ctx:     context.Background(),
	}
}

func (h *userHandler) GetUserByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	user, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve user",
		Data:       user,
	})

	return c.Status(fiber.StatusOK).JSON(res)
}
