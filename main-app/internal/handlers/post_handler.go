package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/errorhandler"
	"github.com/revandpratama/reflect/helper/response"
	"github.com/revandpratama/reflect/internal/dto"
	"github.com/revandpratama/reflect/internal/services"
	"github.com/revandpratama/reflect/types"
)

type postHandler struct {
	service services.PostService
	ctx     context.Context
}

type PostHandler interface {
	GetAllPosts(c *fiber.Ctx) error
	GetPostByID(c *fiber.Ctx) error
	GetPostByUserID(c *fiber.Ctx) error
	CreatePost(c *fiber.Ctx) error
	UpdatePost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
}

func NewPostHandler(service services.PostService) PostHandler {
	return &postHandler{
		service: service,
		ctx:     context.Background(),
	}
}

func (h *postHandler) GetAllPosts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	posts, err := h.service.GetAllPosts(ctx)
	if err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve all posts",
		Data:       posts,
	})

	return c.JSON(res)
}

func (h *postHandler) GetPostByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	post, err := h.service.GetPostByID(ctx, id)
	if err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve post",
		Data:       post,
	})

	return c.JSON(res)
}

func (h *postHandler) GetPostByUserID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, err := c.ParamsInt("userid")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	posts, err := h.service.GetPostByUserID(ctx, userID)
	if err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve post",
		Data:       posts,
	})

	return c.JSON(res)
}

func (h *postHandler) CreatePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.PostRequest
	if err := c.BodyParser(&req); err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	image, err := c.FormFile("image")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	req.Image = image

	if err := h.service.CreatePost(ctx, &req); err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success create post",
	})

	return c.JSON(res)
}

func (h *postHandler) UpdatePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	var req dto.PostRequest
	if err := c.BodyParser(&req); err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	image, err := c.FormFile("image")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	req.Image = image

	if err := h.service.UpdatePost(ctx, id, &req); err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success update post",
	})

	return c.JSON(res)
}

func (h *postHandler) DeletePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	if err := h.service.DeletePost(ctx, id); err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success delete post",
	})

	return c.Status(fiber.StatusOK).JSON(res)
}
