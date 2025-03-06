package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/errorhandler"
	"github.com/revandpratama/reflect/helper"
	"github.com/revandpratama/reflect/helper/response"
	"github.com/revandpratama/reflect/internal/dto"
	"github.com/revandpratama/reflect/internal/services"
	"github.com/revandpratama/reflect/types"
)

type commentHandler struct {
	service services.CommentService
	ctx     context.Context
}

type CommentHandler interface {
	CreateComment(c *fiber.Ctx) error
	GetCommentByID(c *fiber.Ctx) error
	GetCommentByPostID(c *fiber.Ctx) error
	UpdateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
}

func NewCommentHandler(service services.CommentService) CommentHandler {
	return &commentHandler{
		service: service,
		ctx:     context.Background(),
	}
}

func (h *commentHandler) CreateComment(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(h.ctx, 5*time.Second)
	defer cancel()

	var req dto.CommentRequest
	if err := c.BodyParser(&req); err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	errs := helper.ValidateStruct(&req)
	if len(errs) > 0 {
		res := response.NewResponse(&types.ResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "validation error",
			Errors:     errs,
		})

		return c.JSON(res)
	}

	if err := h.service.CreateComment(ctx, &req); err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success create comment",
	})

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *commentHandler) GetCommentByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(h.ctx, 5*time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	comment, err := h.service.GetCommentByID(ctx, id)
	if err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve comment",
		Data:       comment,
	})

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *commentHandler) GetCommentByPostID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(h.ctx, 5*time.Second)
	defer cancel()

	postID, err := c.ParamsInt("postid")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	comments, err := h.service.GetCommentByPostID(ctx, postID)
	if err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve comments",
		Data:       comments,
	})

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *commentHandler) UpdateComment(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(h.ctx, 5*time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	var req dto.CommentRequest
	if err := c.BodyParser(&req); err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	errs := helper.ValidateStruct(&req)
	if len(errs) > 0 {
		res := response.NewResponse(&types.ResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "validation error",
			Errors:     errs,
		})

		return c.JSON(res)
	}

	if err := h.service.UpdateComment(ctx, id, &req); err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success update comment",
	})

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *commentHandler) DeleteComment(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(h.ctx, 5*time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	if err := h.service.DeleteComment(ctx, id); err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success delete comment",
	})

	return c.Status(fiber.StatusOK).JSON(res)
}
