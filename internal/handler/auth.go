package handler

import (
	"github.com/aliaxy/byte-cabinet/internal/model"
	"github.com/aliaxy/byte-cabinet/internal/service"
	"github.com/aliaxy/byte-cabinet/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	authService *service.AuthService
	validate    *validator.Validate
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validator.New(),
	}
}

// Login handles user login requests
// POST /api/v1/auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req model.UserLogin

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := h.validate.Struct(&req); err != nil {
		return response.ValidationError(c, "Username and password are required")
	}

	// Attempt login
	result, err := h.authService.Login(c.Context(), &req)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			return response.Unauthorized(c, "Invalid username or password")
		}
		return response.InternalError(c, "")
	}

	return response.OK(c, result)
}

// Logout handles user logout requests
// POST /api/v1/auth/logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// For JWT-based auth, logout is typically handled client-side
	// by removing the token. This endpoint exists for consistency.
	return response.OKWithMessage(c, nil, "Logged out successfully")
}

// Me returns the current authenticated user's information
// GET /api/v1/auth/me
func (h *AuthHandler) Me(c *fiber.Ctx) error {
	// Get user ID from context (set by auth middleware)
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		return response.Unauthorized(c, "")
	}

	user, err := h.authService.GetCurrentUser(c.Context(), userID)
	if err != nil {
		if err == service.ErrUserNotFound {
			return response.NotFound(c, "User not found")
		}
		return response.InternalError(c, "")
	}

	return response.OK(c, user)
}

// ChangePassword handles password change requests
// PUT /api/v1/auth/password
func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	// Get user ID from context
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		return response.Unauthorized(c, "")
	}

	var req model.PasswordChange

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := h.validate.Struct(&req); err != nil {
		return response.ValidationError(c, "Both old and new passwords are required")
	}

	// Attempt password change
	err := h.authService.ChangePassword(c.Context(), userID, &req)
	if err != nil {
		switch err {
		case service.ErrUserNotFound:
			return response.NotFound(c, "User not found")
		case service.ErrInvalidOldPassword:
			return response.BadRequest(c, "Current password is incorrect")
		default:
			return response.InternalError(c, "")
		}
	}

	return response.OKWithMessage(c, nil, "Password changed successfully")
}

// RefreshToken handles token refresh requests
// POST /api/v1/auth/refresh
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	type RefreshRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	var req RefreshRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := h.validate.Struct(&req); err != nil {
		return response.ValidationError(c, "Refresh token is required")
	}

	// Attempt token refresh
	tokens, err := h.authService.RefreshTokens(c.Context(), req.RefreshToken)
	if err != nil {
		return response.Unauthorized(c, "Invalid or expired refresh token")
	}

	return response.OK(c, tokens)
}

// UpdateProfile handles profile update requests
// PUT /api/v1/auth/profile
func (h *AuthHandler) UpdateProfile(c *fiber.Ctx) error {
	// Get user ID from context
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		return response.Unauthorized(c, "")
	}

	var req model.UserUpdate

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := h.validate.Struct(&req); err != nil {
		return response.ValidationError(c, "Invalid profile data")
	}

	// Update profile
	user, err := h.authService.UpdateProfile(c.Context(), userID, &req)
	if err != nil {
		if err == service.ErrUserNotFound {
			return response.NotFound(c, "User not found")
		}
		return response.InternalError(c, "")
	}

	return response.OKWithMessage(c, user, "Profile updated successfully")
}

// RegisterRoutes registers all auth routes
func (h *AuthHandler) RegisterRoutes(app fiber.Router, authMiddleware fiber.Handler) {
	auth := app.Group("/auth")

	// Public routes
	auth.Post("/login", h.Login)
	auth.Post("/refresh", h.RefreshToken)

	// Protected routes
	auth.Post("/logout", authMiddleware, h.Logout)
	auth.Get("/me", authMiddleware, h.Me)
	auth.Put("/password", authMiddleware, h.ChangePassword)
	auth.Put("/profile", authMiddleware, h.UpdateProfile)
}
