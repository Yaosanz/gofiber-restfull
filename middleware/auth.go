package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang-resfull/utils"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(fiber.HeaderAuthorization)
		if authHeader == "" {
			return unauthorizedResponse(c, "Need Atuhorization")
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			return unauthorizedResponse(c, "Authorization header must start with Bearer")
		}

		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return unauthorizedResponse(c, "Invalid token: "+err.Error())
		}

		c.Locals("userID", claims["sub"])
		c.Locals("email", claims["email"])

		return c.Next()
	}
}

func unauthorizedResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": message,
	})
}
