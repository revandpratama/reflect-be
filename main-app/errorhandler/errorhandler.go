package errorhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/helper/response"
	"github.com/revandpratama/reflect/types"
)

func BuildError(c *fiber.Ctx, err error) error {
	var statusCode int

	switch err.(type) {
	case *types.NotFoundError:
		statusCode = fiber.StatusNotFound
	case *types.BadRequestError:
		statusCode = fiber.StatusBadRequest
	case *types.UnauthorizedError:
		statusCode = fiber.StatusUnauthorized
	case *types.InternalServerError:
		statusCode = fiber.StatusInternalServerError
	default:
		statusCode = fiber.StatusInternalServerError
	}

	response := response.NewResponse(&types.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	return c.Status(statusCode).JSON(response)
}
