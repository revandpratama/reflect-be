package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/errorhandler"
	helper "github.com/revandpratama/reflect/helper/token"
	"github.com/revandpratama/reflect/types"
)

func AuthMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return errorhandler.BuildError(c, &types.UnauthorizedError{Message: "unauthorized"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return errorhandler.BuildError(c, &types.UnauthorizedError{Message: "unauthorized"})
		}

		encryptedToken := parts[1]
		user, err := helper.VerifyToken(encryptedToken)
		if err != nil {
			return errorhandler.BuildError(c, &types.UnauthorizedError{Message: "unauthorized"})
		}

		c.Locals("user_id", user.ID)
		c.Locals("user_name", user.Name)
		c.Locals("user_email", user.Email)

		return c.Next()
	}
}
