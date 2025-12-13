package middleware

import (
	"strings"

	"github.com/aliaxy/byte-cabinet/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware creates a JWT authentication middleware
func AuthMiddleware(jwtManager *utils.JWTManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "UNAUTHORIZED",
					"message": "Authorization header is required",
				},
			})
		}

		// Check Bearer prefix
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "UNAUTHORIZED",
					"message": "Invalid authorization header format",
				},
			})
		}

		tokenString := parts[1]

		// Validate access token
		claims, err := jwtManager.ValidateAccessToken(tokenString)
		if err != nil {
			message := "Invalid token"
			if err == utils.ErrExpiredToken {
				message = "Token has expired"
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "UNAUTHORIZED",
					"message": message,
				},
			})
		}

		// Store user info in context
		c.Locals("userID", claims.UserID)
		c.Locals("username", claims.Username)

		return c.Next()
	}
}

// OptionalAuthMiddleware creates middleware that extracts user info if token is present
// but doesn't require authentication
func OptionalAuthMiddleware(jwtManager *utils.JWTManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Next()
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return c.Next()
		}

		claims, err := jwtManager.ValidateAccessToken(parts[1])
		if err != nil {
			return c.Next()
		}

		c.Locals("userID", claims.UserID)
		c.Locals("username", claims.Username)

		return c.Next()
	}
}

// GetUserID retrieves the user ID from the context
func GetUserID(c *fiber.Ctx) (int64, bool) {
	userID, ok := c.Locals("userID").(int64)
	return userID, ok
}

// GetUsername retrieves the username from the context
func GetUsername(c *fiber.Ctx) (string, bool) {
	username, ok := c.Locals("username").(string)
	return username, ok
}
