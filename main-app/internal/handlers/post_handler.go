package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/revandpratama/reflect/errorhandler"
	"github.com/revandpratama/reflect/helper"
	"github.com/revandpratama/reflect/helper/response"
	"github.com/revandpratama/reflect/internal/dto"
	"github.com/revandpratama/reflect/internal/services"
	"github.com/revandpratama/reflect/types"
	"github.com/valyala/fasthttp"
)

type postHandler struct {
	service     services.PostService
	ctx         context.Context
	redisClient *redis.Client
}

type PostHandler interface {
	GetAllPosts(c *fiber.Ctx) error
	GetPostByID(c *fiber.Ctx) error
	GetPostByUserID(c *fiber.Ctx) error
	CreatePost(c *fiber.Ctx) error
	UpdatePost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
}

func NewPostHandler(service services.PostService, redisClient *redis.Client) PostHandler {
	return &postHandler{
		service:     service,
		ctx:         context.Background(),
		redisClient: redisClient,
	}
}

func (h *postHandler) GetAllPosts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 5)

	redisKey := fmt.Sprintf("posts:page:%d:limit-%d", page, limit)

	byteVal, err := helper.GetFromRedis(ctx, h.redisClient, redisKey)
	if err == nil {
		var res any
		if err := json.Unmarshal(byteVal, &res); err != nil {
			return errorhandler.BuildError(c, &types.InternalServerError{Message: err.Error()})
		}

		// res := response.NewResponse(&responseParam)
		helper.NewLog().Info("success retrieve all posts from redis")
		return c.Status(fiber.StatusOK).JSON(res)
	}

	posts, pagination, err := h.service.GetAllPosts(ctx, page, limit)
	if err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success retrieve all posts",
		Data:       posts,
		Pagination: pagination,
	})

	if err := helper.SetToRedis(ctx, h.redisClient, redisKey, res); err != nil {
		helper.NewLog().Error(fmt.Sprintf("error set redis cache: %v", err))
	}

	return c.Status(fiber.StatusOK).JSON(res)
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

	return c.Status(fiber.StatusOK).JSON(res)
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

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *postHandler) CreatePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.PostRequest

	image, err := c.FormFile("image")
	if err == nil {
		req.Image = image // Set the image only if it's provided
	} else if err != fasthttp.ErrMissingFile {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: fmt.Sprintf("Invalid image: %s", err.Error())})
	}

	if err := c.BodyParser(&req); err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	req.UserID = c.Locals("user_id").(int)

	errs := helper.ValidateStruct(&req)
	if len(errs) > 0 {
		res := response.NewResponse(&types.ResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "validation error",
			Errors:     errs,
		})

		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if err := h.service.CreatePost(ctx, &req); err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success create post",
	})

	if err := helper.DeleteRedisCache(ctx, h.redisClient, helper.REDIS_POSTS_PATTERN); err != nil {
		helper.NewLog().Error(fmt.Sprintf("error delete redis cache: %v", err))
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *postHandler) UpdatePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	var req dto.PostRequest

	image, err := c.FormFile("image")
	if err == nil {
		req.Image = image // Set the image only if it's provided
	} else if err != fasthttp.ErrMissingFile {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: fmt.Sprintf("Invalid image: %s", err.Error())})
	}

	if err := c.BodyParser(&req); err != nil {
		return errorhandler.BuildError(c, &types.BadRequestError{Message: err.Error()})
	}

	req.UserID = c.Locals("id").(int)

	errs := helper.ValidateStruct(&req)
	if len(errs) > 0 {
		res := response.NewResponse(&types.ResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "validation error",
			Errors:     errs,
		})

		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if err := h.service.UpdatePost(ctx, id, &req); err != nil {
		return errorhandler.BuildError(c, err)
	}

	res := response.NewResponse(&types.ResponseParams{
		StatusCode: fiber.StatusOK,
		Message:    "success update post",
	})

	if err := helper.DeleteRedisCache(ctx, h.redisClient, helper.REDIS_POSTS_PATTERN); err != nil {
		helper.NewLog().Error(fmt.Sprintf("error delete redis cache: %v", err))
	}

	return c.Status(fiber.StatusOK).JSON(res)
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

	if err := helper.DeleteRedisCache(ctx, h.redisClient, helper.REDIS_POSTS_PATTERN); err != nil {
		helper.NewLog().Error(fmt.Sprintf("error delete redis cache: %v", err))
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
